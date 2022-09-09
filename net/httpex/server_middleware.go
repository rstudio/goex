package httpex

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

func ServerMiddleware(name, version string, tags []string) MiddlewareFunc {
	headerParts := []string{name + "/" + version}

	for _, tag := range append(tags, runtime.Version()) {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}

		headerParts = append(headerParts, fmt.Sprintf("(%[1]s)", tag))
	}

	headerValue := strings.Join(headerParts, " ")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Server", headerValue)
			next.ServeHTTP(w, req)
		})
	}
}
