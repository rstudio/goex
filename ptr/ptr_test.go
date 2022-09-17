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
		// NOTE: figure out how to do this (if it's even possible ?)
		// {(*string)(nil), ""},
		// {(*bool)(nil), false},
		// {(*int)(nil), 0},
		// {(*uint)(nil), uint(0)},
		// {(*float64)(nil), float64(0)},
		// {(*[]int)(nil), []int{}},
		// {(*map[string]bool)(nil), map[string]bool{}},
		// {(*struct{ banjo string })(nil), struct{ banjo string }{}},
	} {
		t.Run(fmt.Sprintf("%[1]T", tc.in), func(t *testing.T) {
			require.Equal(t, tc.out, ptr.From(ptr.To(tc.in)))
		})
	}
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
