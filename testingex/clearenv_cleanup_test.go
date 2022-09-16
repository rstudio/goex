package testingex_test

import (
	"os"
	"testing"

	"github.com/rstudio/goex/testingex"
	"github.com/stretchr/testify/require"
)

func TestClearenvCleanup(t *testing.T) {
	os.Setenv("ANSWER", "42")

	t.Run("sub", func(t *testing.T) {
		t.Cleanup(testingex.ClearenvCleanup())

		_, ok := os.LookupEnv("ANSWER")
		require.False(t, ok)

		os.Setenv("LEAKED", "oh no")
	})

	require.Equal(t, "42", os.Getenv("ANSWER"))

	_, ok := os.LookupEnv("LEAKED")
	require.False(t, ok)
}
