package zapex_test

import (
	"testing"

	"github.com/rstudio/goex/zapex"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	for _, tc := range []struct {
		envName string
		debug   bool
	}{
		{
			envName: "dev",
			debug:   false,
		},
		{
			envName: "test",
			debug:   false,
		},
		{
			envName: "anythingelse",
			debug:   true,
		},
	} {
		zl := zapex.NewLogger(tc.envName, tc.debug)
		assert.NotNil(t, zl, "for case %v+%v", tc.envName, tc.debug)
	}
}
