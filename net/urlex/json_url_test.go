package urlex_test

import (
	"net/url"
	"testing"

	"github.com/rstudio/goex/net/urlex"
	"github.com/stretchr/testify/require"
)

func TestJSONURL(t *testing.T) {
	r := require.New(t)

	r.Nil(urlex.NewJSONURL(nil))

	ju := urlex.NewJSONURL(&url.URL{Host: "squirtle.example.org"})

	r.NotNil(ju)
	r.Equal(&url.URL{Host: "squirtle.example.org"}, ju.URLCopy())

	jub, err := ju.MarshalJSON()
	r.Nil(err)
	r.NotNil(jub)

	ju2 := urlex.NewJSONURL(&url.URL{Host: "charizard.example.org"})
	r.NotEqual(ju2, ju)
	r.Nil(ju2.UnmarshalJSON(jub))
	r.Equal(ju2, ju)

	ju2 = urlex.NewJSONURL(&url.URL{Host: "charizard.example.org"})

	r.Nil(ju2.UnmarshalJSON([]byte("\"\"")))
	r.Equal(&urlex.JSONURL{URL: url.URL{}}, ju2)

	ju2 = urlex.NewJSONURL(&url.URL{Host: "charizard.example.org"})

	r.Nil(ju2.UnmarshalJSON([]byte("null")))
	r.Equal(&urlex.JSONURL{URL: url.URL{}}, ju2)

	r.NotNil(ju.UnmarshalJSON([]byte("")))
	r.NotNil(ju.UnmarshalJSON([]byte("\"\x00\xba\xdd\xad\"")))
}
