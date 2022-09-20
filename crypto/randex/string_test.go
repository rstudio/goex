package randex_test

import (
	"testing"
	"unicode"

	"github.com/rstudio/goex/crypto/randex"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	t.Run("lengths", func(t *testing.T) {
		r := require.New(t)

		r.Len(randex.String(5), 5)
		r.Len(randex.String(12), 12)
	})

	t.Run("nonzero", func(t *testing.T) {
		for i := 0; i < 10_000; i++ {
			require.NotEqual(t, "00000", randex.String(5))
		}
	})

	t.Run("printable", func(t *testing.T) {
		for i := 0; i < 10_000; i++ {
			s := randex.String(len(randex.StringRunes()))
			for _, chr := range s {
				require.True(t, unicode.IsPrint(chr))
			}
		}
	})
}

func TestRuneOf(t *testing.T) {
	letters := []rune("correcthorsebatterystaple")

	for i := 0; i < 100_000; i++ {
		require.NotEqual(t, rune(0), randex.RuneOf(letters))
	}
}
