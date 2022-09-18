package ptr_test

import (
	"fmt"
	"testing"

	"github.com/rstudio/goex/ptr"
	"github.com/stretchr/testify/require"
)

func TestTo(t *testing.T) {
	for _, in := range []any{
		"greninja",
		true,
		-11,
		uint(11),
		22.2,
		[]int{4, 2},
		map[string]bool{"chomp": true},
		struct{ banjo string }{banjo: "kazooie"},
	} {
		t.Run(fmt.Sprintf("%[1]v (%[1]T)", in), func(t *testing.T) {
			require.NotNil(t, ptr.To(in))
		})
	}
}

func TestFrom(t *testing.T) {
	for _, tc := range []struct {
		in  any
		out any
	}{
		{"greninja", "greninja"},
		{true, true},
		{-11, -11},
		{uint(11), uint(11)},
		{22.2, 22.2},
		{[]int{4, 2}, []int{4, 2}},
		{map[string]bool{"chomp": true}, map[string]bool{"chomp": true}},
		{struct{ banjo string }{banjo: "kazooie"}, struct{ banjo string }{banjo: "kazooie"}},
	} {
		t.Run(fmt.Sprintf("%[1]T", tc.in), func(t *testing.T) {
			require.Equal(t, tc.out, ptr.From(ptr.To(tc.in)))
		})
	}

	t.Run("*string to zero", func(t *testing.T) {
		var v *string
		require.Equal(t, "", ptr.From(v))
	})

	t.Run("*bool to zero", func(t *testing.T) {
		var v *bool
		require.Equal(t, false, ptr.From(v))
	})

	t.Run("*int to zero", func(t *testing.T) {
		var v *int
		require.Equal(t, 0, ptr.From(v))
	})

	t.Run("*uint to zero", func(t *testing.T) {
		var v *uint
		require.Equal(t, uint(0), ptr.From(v))
	})

	t.Run("*float64 to zero", func(t *testing.T) {
		var v *float64
		require.Equal(t, float64(0), ptr.From(v))
	})

	t.Run("*[]int to zero", func(t *testing.T) {
		var v *[]int
		require.Equal(t, ([]int)(nil), ptr.From(v))
	})

	t.Run("*map[string]bool to zero", func(t *testing.T) {
		var v *map[string]bool
		require.Equal(t, (map[string]bool)(nil), ptr.From(v))
	})

	t.Run("*struct to zero", func(t *testing.T) {
		var v *struct{ banjo string }
		require.Equal(t, struct{ banjo string }{}, ptr.From(v))
	})
}

func TestRoundTrip(t *testing.T) {
	for _, in := range []any{
		"greninja",
		true,
		-11,
		uint(11),
		22.2,
		[]int{4, 2},
		map[string]bool{"chomp": true},
		struct{ banjo string }{banjo: "kazooie"},
		(*string)(nil),
		(*bool)(nil),
		(*int)(nil),
		(*uint)(nil),
		(*float64)(nil),
		(*[]int)(nil),
		(*map[string]bool)(nil),
		(*struct{ banjo string })(nil),
	} {
		t.Run(fmt.Sprintf("%[1]T", in), func(t *testing.T) {
			require.Equal(t, in, ptr.From(ptr.To(in)))
		})
	}
}
