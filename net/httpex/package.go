package httpex

import (
	"net"
	"net/http"
)

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
