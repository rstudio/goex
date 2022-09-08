package httpex_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rstudio/goex/net/httpex"
	"github.com/stretchr/testify/assert"
)

func TestFrom(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/explosions", nil)
	req.Header.Set("X-Forwarded-For", "onion-scooter.local")

	assert.Equal(t, httpex.From(req), "onion-scooter.local")

	req = httptest.NewRequest(http.MethodGet, "/guacola", nil)
	req.RemoteAddr = "ketchup-packet.local:80"

	assert.Equal(t, httpex.From(req), "ketchup-packet.local")

	req.RemoteAddr = "\x00\xba\xdd\xad"

	assert.Equal(t, httpex.From(req), "\x00\xba\xdd\xad")
}
