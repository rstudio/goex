package httpex_test

import (
	"errors"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rstudio/goex/net/httpex"
	"github.com/rstudio/goex/zapex"
	"github.com/stretchr/testify/require"
)

func TestJSON(t *testing.T) {
	r := require.New(t)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "http://zebes.example.org/zebesians", nil)

	httpex.JSON(
		w, req, http.StatusTeapot,
		map[string]any{"count": 4, "level": "purple"},
	)

	r.Equal(http.StatusTeapot, w.Code)
	r.JSONEq(`{"count":4,"level":"purple"}`, w.Body.String())

	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "http://zebes.example.org/geemers", nil)

	httpex.JSON(
		w, req, http.StatusNotFound,
		errors.New("no geemers error"),
	)

	r.Equal(http.StatusNotFound, w.Code)
	r.JSONEq(`{"error":"no geemers error"}`, w.Body.String())

	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "http://zebes.example.org/zoomers", nil)
	req = req.WithContext(zapex.ContextWithLogger(req.Context(), zapex.NewLogger("test", false).Sugar()))

	httpex.JSON(
		w, req, http.StatusNotFound,
		math.Inf(-1),
	)

	r.Equal(http.StatusNotFound, w.Code)
	r.Equal("", w.Body.String())
}
