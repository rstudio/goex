package httpex

import (
	"encoding/json"
	"net/http"

	"github.com/rstudio/goex/zapex"
)

// JSON writes a response and ensures the content type is set to
// application/json. In the case that data satisfies the error
// interface, the response body will have a single "error" key with
// value err.Error().
func JSON(w http.ResponseWriter, req *http.Request, status int, data any) {
	if err, ok := data.(error); ok {
		data = map[string]string{"error": err.Error()}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		if logger, ok := zapex.LoggerFromContext(req.Context()); ok {
			logger.Errorw("failed to encode response", "err", err)
		}
	}
}
