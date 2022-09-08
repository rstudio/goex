package randex_test

import (
	"testing"

	"github.com/rstudio/goex/crypto/randex"
	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	assert.Len(t, randex.String(5), 5)
	assert.Len(t, randex.String(12), 12)
}
