package httpex_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"

	"github.com/rstudio/goex/net/httpex"
	"github.com/stretchr/testify/require"
)

func TestServerMiddleware(t *testing.T) {
	for _, tc := range []struct {
		name    string
		version string
		tags    []string
		verify  func(*testing.T, string)
	}{
		{
			name:    "typical",
			version: "v1.31.2",
			tags:    []string{runtime.GOOS, runtime.GOARCH, runtime.Version()},
			verify: func(t *testing.T, server string) {
				require.Equal(
					t,
					fmt.Sprintf(
						"typical/v1.31.2 (%[1]v) (%[2]v) (%[3]v)",
						runtime.GOOS,
						runtime.GOARCH,
						runtime.Version(),
					),
					server,
				)
			},
		},
		{
			name:    "minimal",
			version: "v0",
			verify: func(t *testing.T, server string) {
				require.Equal(t, "minimal/v0", server)
			},
		},
		{
			name:    "chatty",
			version: "v13.12.1312-alpha2",
			tags: []string{
				"correct",
				"horse",
				runtime.GOOS,
				runtime.GOARCH,
				runtime.Version(),
				"battery",
				"staple",
			},
			verify: func(t *testing.T, server string) {
				r := require.New(t)
				r.True(
					strings.HasPrefix(
						server,
						"chatty/v13.12.1312-alpha2 (correct) (horse)",
					),
				)
				r.True(
					strings.HasSuffix(
						server,
						"(battery) (staple)",
					),
				)
			},
		},
	} {
		t.Run(
			tc.name,
			func(t *testing.T) {
				w := httptest.NewRecorder()
				httpex.ServerMiddleware(
					tc.name, tc.version, tc.tags,
				)(
					http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
						w.WriteHeader(http.StatusTeapot)
					}),
				).ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/", nil))
				tc.verify(t, w.Header().Get("Server"))
			},
		)
	}
}
