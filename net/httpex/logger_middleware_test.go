package httpex_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rstudio/goex/net/httpex"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestLoggerMiddleware(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if req.URL.Path == "/ohno" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if req.URL.Path == "/boom" {
			panic("boom")
		}

		w.WriteHeader(http.StatusNotFound)
	})

	for _, tc := range []struct {
		name   string
		req    *http.Request
		verify func(*testing.T, []observer.LoggedEntry, *httptest.ResponseRecorder)
	}{
		{
			name: "success",
			req:  httptest.NewRequest(http.MethodGet, "/", nil),
			verify: func(t *testing.T, entries []observer.LoggedEntry, w *httptest.ResponseRecorder) {
				r := require.New(t)

				r.Len(entries, 3)

				r.Equal("no tracing span found in context", entries[0].Entry.Message)
				r.Equal("request", entries[1].Entry.Message)

				reqFields := map[string]zapcore.Field{}
				for _, fld := range entries[1].Context {
					reqFields[fld.Key] = fld
				}

				r.NotEqual("", reqFields["id"].String)
				reqID := reqFields["id"].String

				r.Equal(http.MethodGet, reqFields["method"].String)
				r.Equal("/", reqFields["url"].String)
				r.NotEqual("", reqFields["from"].String)
				reqFrom := reqFields["from"].String

				r.Equal("response", entries[2].Entry.Message)

				respFields := map[string]zapcore.Field{}
				for _, fld := range entries[2].Context {
					respFields[fld.Key] = fld
				}

				r.Equal(reqID, respFields["id"].String)
				r.Equal(http.MethodGet, respFields["method"].String)
				r.Equal("/", respFields["url"].String)
				r.Equal(reqFrom, respFields["from"].String)
				r.NotEqual(0, respFields["time"].Integer)
				r.Equal(int64(http.StatusOK), respFields["status"].Integer)

				r.Equal(http.StatusOK, w.Code)
			},
		},
		{
			name: "error",
			req:  httptest.NewRequest(http.MethodPost, "/ohno", nil),
			verify: func(t *testing.T, entries []observer.LoggedEntry, w *httptest.ResponseRecorder) {
				r := require.New(t)

				r.Len(entries, 3)

				r.Equal("no tracing span found in context", entries[0].Entry.Message)
				r.Equal("request", entries[1].Entry.Message)

				reqFields := map[string]zapcore.Field{}
				for _, fld := range entries[1].Context {
					reqFields[fld.Key] = fld
				}

				r.NotEqual("", reqFields["id"].String)
				reqID := reqFields["id"].String

				r.Equal(http.MethodPost, reqFields["method"].String)
				r.Equal("/ohno", reqFields["url"].String)
				r.NotEqual("", reqFields["from"].String)
				reqFrom := reqFields["from"].String

				r.Equal("response", entries[2].Entry.Message)

				respFields := map[string]zapcore.Field{}
				for _, fld := range entries[2].Context {
					respFields[fld.Key] = fld
				}

				r.Equal(reqID, respFields["id"].String)
				r.Equal(http.MethodPost, respFields["method"].String)
				r.Equal("/ohno", respFields["url"].String)
				r.Equal(reqFrom, respFields["from"].String)
				r.NotEqual(0, respFields["time"].Integer)
				r.Equal(int64(http.StatusInternalServerError), respFields["status"].Integer)

				r.Equal(http.StatusInternalServerError, w.Code)
			},
		},
		{
			name: "unknown",
			req:  httptest.NewRequest(http.MethodPatch, "/wat", nil),
			verify: func(t *testing.T, entries []observer.LoggedEntry, w *httptest.ResponseRecorder) {
				r := require.New(t)

				r.Len(entries, 3)

				r.Equal("no tracing span found in context", entries[0].Entry.Message)
				r.Equal("request", entries[1].Entry.Message)

				reqFields := map[string]zapcore.Field{}
				for _, fld := range entries[1].Context {
					reqFields[fld.Key] = fld
				}

				r.NotEqual("", reqFields["id"].String)
				reqID := reqFields["id"].String

				r.Equal(http.MethodPatch, reqFields["method"].String)
				r.Equal("/wat", reqFields["url"].String)
				r.NotEqual("", reqFields["from"].String)
				reqFrom := reqFields["from"].String

				r.Equal("response", entries[2].Entry.Message)

				respFields := map[string]zapcore.Field{}
				for _, fld := range entries[2].Context {
					respFields[fld.Key] = fld
				}

				r.Equal(reqID, respFields["id"].String)
				r.Equal(http.MethodPatch, respFields["method"].String)
				r.Equal("/wat", respFields["url"].String)
				r.Equal(reqFrom, respFields["from"].String)
				r.NotEqual(0, respFields["time"].Integer)
				r.Equal(int64(http.StatusNotFound), respFields["status"].Integer)

				r.Equal(http.StatusNotFound, w.Code)
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			zc, logs := observer.New(zap.DebugLevel)
			logger := zap.New(zc).Sugar()
			httpex.LoggerMiddleware(logger)(h).ServeHTTP(w, tc.req)
			tc.verify(t, logs.AllUntimed(), w)
		})
	}

	t.Run("panics", func(t *testing.T) {
		w := httptest.NewRecorder()
		zc, _ := observer.New(zap.DebugLevel)
		logger := zap.New(zc).Sugar()

		require.Panics(t, func() {
			httpex.LoggerMiddleware(logger)(h).
				ServeHTTP(w, httptest.NewRequest(http.MethodOptions, "/boom", nil))
		})
	})
}
