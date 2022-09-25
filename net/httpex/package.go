package httpex

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

type MiddlewareFunc func(http.Handler) http.Handler

func From(req *http.Request) string {
	if f := req.Header.Get("X-Forwarded-For"); f != "" {
		return f
	}

	f := req.RemoteAddr

	ip, _, err := net.SplitHostPort(f)
	if err != nil {
		return f
	}

	return ip
}

func Server(name, version string, tags []string) string {
	headerParts := []string{name + "/" + version}

	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}

		headerParts = append(headerParts, fmt.Sprintf("(%[1]s)", tag))
	}

	return strings.Join(headerParts, " ")
}
