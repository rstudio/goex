package tlsex_test

import (
	"crypto/tls"
	"testing"

	"github.com/rstudio/goex/crypto/tlsex"
	"github.com/stretchr/testify/require"
)

func TestThumbprint(t *testing.T) {
	t.Run("missing port", func(t *testing.T) {
		r := require.New(t)

		_, err := tlsex.Thumbprint("posit.co", nil)
		r.Error(err)
	})

	t.Run("bogus address", func(t *testing.T) {
		r := require.New(t)

		_, err := tlsex.Thumbprint("\xb0\xb1\x00", nil)
		r.Error(err)
	})

	t.Run("known thumbprint", func(t *testing.T) {
		r := require.New(t)

		actual, err := tlsex.Thumbprint("posit.co:443", nil)
		r.NoError(err)
		// NOTE: This assertion *will fail* when the root certificate changes
		r.Equal("00abefd055f9a9c784ffdeabd1dcdd8fed741436", actual)
	})

	t.Run("known thumbprint with tls config", func(t *testing.T) {
		r := require.New(t)

		actual, err := tlsex.Thumbprint("posit.co:443", &tls.Config{ServerName: "posit.co"})
		r.NoError(err)
		// NOTE: This assertion *will fail* when the root certificate changes
		r.Equal("00abefd055f9a9c784ffdeabd1dcdd8fed741436", actual)
	})

	t.Run("other thumbprint", func(t *testing.T) {
		r := require.New(t)

		actual, err := tlsex.Thumbprint("example.com:443", nil)
		r.NoError(err)
		r.NotEqual("", actual)
	})

	t.Run("other thumbprint with tls config", func(t *testing.T) {
		r := require.New(t)

		actual, err := tlsex.Thumbprint("example.com:443", &tls.Config{InsecureSkipVerify: true})
		r.NoError(err)
		r.NotEqual("", actual)
	})
}
