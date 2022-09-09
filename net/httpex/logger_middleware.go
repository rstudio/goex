package httpex

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rstudio/goex/crypto/randex"
	"github.com/rstudio/goex/zapex"
	"go.uber.org/zap"
)

const (
	reqIDLen = 12
)

func LoggerMiddleware(logger *zap.SugaredLogger) MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			reqLog := zapex.LoggerWithTraceFields(req.Context(), logger)
			reqID := randex.String(reqIDLen)
			started := time.Now()

			reqLog.Infow("request",
				"id", reqID,
				"method", req.Method,
				"path", req.URL.String(),
				"from", From(req),
			)

			h.ServeHTTP(w, req.WithContext(zapex.ContextWithLogger(req.Context(), reqLog)))

			since := time.Since(started).String()
			reqLog.Infow("response",
				"id", reqID,
				"method", req.Method,
				"path", req.URL.String(),
				"from", From(req),
				"time", since,
			)

			if err := zap.L().Sync(); err != nil {
				fmt.Fprintf(os.Stdout, "zapex:ERROR: failed to sync global zap logger: %v\n", err)
			}
		})
	}
}
