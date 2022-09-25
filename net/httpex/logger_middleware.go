package httpex

import (
	"net/http"
	"time"

	"github.com/rstudio/goex/crypto/randex"
	"github.com/rstudio/goex/zapex"
	"go.uber.org/zap"
)

const (
	defaultReqIDLen = 12
)

type LoggerMiddlewareOptions struct {
	ReqIDLen int
}

type loggingResponseWriter struct {
	http.ResponseWriter

	status int
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	if lrw.status == 0 {
		lrw.WriteHeader(http.StatusOK)
	}

	return lrw.ResponseWriter.Write(b)
}

func (lrw *loggingResponseWriter) WriteHeader(status int) {
	if lrw.status != 0 {
		return
	}

	lrw.ResponseWriter.WriteHeader(status)
	lrw.status = status
}

func LoggerMiddleware(logger *zap.SugaredLogger, optFn ...func(*LoggerMiddlewareOptions)) MiddlewareFunc {
	opts := &LoggerMiddlewareOptions{
		ReqIDLen: defaultReqIDLen,
	}

	for _, f := range optFn {
		f(opts)
	}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			reqLog := zapex.LoggerWithTraceFields(req.Context(), logger).Desugar()

			reqID := randex.String(opts.ReqIDLen)
			started := time.Now()

			reqLog.Info(
				"request",
				zap.String("id", reqID),
				zap.String("method", req.Method),
				zap.String("url", req.URL.String()),
				zap.String("from", From(req)),
			)

			lrw := &loggingResponseWriter{ResponseWriter: w}

			h.ServeHTTP(lrw, req.WithContext(zapex.ContextWithLogger(req.Context(), reqLog.Sugar())))

			fields := []zap.Field{
				zap.String("id", reqID),
				zap.String("method", req.Method),
				zap.String("url", req.URL.String()),
				zap.String("from", From(req)),
				zap.Duration("time", time.Since(started)),
			}

			if lrw.status != 0 {
				fields = append(fields, zap.Int("status", lrw.status))
			}

			reqLog.Info("response", fields...)
		})
	}
}
