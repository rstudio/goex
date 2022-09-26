package httpex

import (
	"net/http"
)

func ServerMiddleware(name, version string, tags []string) MiddlewareFunc {
	headerValue := Server(name, version, tags)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Server", headerValue)
			next.ServeHTTP(w, req)
		})
	}
}
