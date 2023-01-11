package containerex_test

import (
	"testing"

	"github.com/rstudio/goex/containerex"
	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	t.Run("string slice", func(t *testing.T) {
		strSl := []string{"happy", "happy", "joy", "joy"}

		require.True(t, containerex.Contains(strSl, "happy"))
		require.True(t, containerex.Contains(strSl, "joy"))
		require.False(t, containerex.Contains(strSl, "toast"))
	})

	t.Run("int slice", func(t *testing.T) {
		intSl := []int{1, 31, 2}

		require.True(t, containerex.Contains(intSl, 2))
		require.True(t, containerex.Contains(intSl, 31))
		require.False(t, containerex.Contains(intSl, 3))
	})

	t.Run("bool slice", func(t *testing.T) {
		boolSl := []bool{true, true, true, true, true}

		require.True(t, containerex.Contains(boolSl, true))
		require.False(t, containerex.Contains(boolSl, false))
	})

	t.Run("complex comparable slice", func(t *testing.T) {
		owl := &aminal{
			Personality: 4,
			Sound:       "hoot",
		}

		amnSl := []*aminal{
			{
				Personality: 9001,
				Sound:       "ruff",
			},
			owl,
		}

		require.True(t, containerex.Contains(amnSl, owl))
		require.False(t, containerex.Contains(amnSl, &aminal{Personality: 9001, Sound: "ruff"}))
		require.False(t, containerex.Contains(amnSl, &aminal{Personality: 8000, Sound: "mew"}))
	})
}

func TestContainsEq(t *testing.T) {
	t.Run("eqer slice", func(t *testing.T) {
		owl := &aminal{
			Personality: 4,
			Sound:       "hoot",
		}

		amnSl := []*aminal{
			{
				Personality: 9001,
				Sound:       "ruff",
			},
			owl,
		}

		require.True(t, containerex.ContainsEq(amnSl, owl))
		require.True(t, containerex.ContainsEq(amnSl, &aminal{Personality: 9001, Sound: "ruff"}))
		require.False(t, containerex.ContainsEq(amnSl, &aminal{Personality: 8000, Sound: "mew"}))
	})

	t.Run("spoofy eqer slice", func(t *testing.T) {
		tree := &plamt{
			Personality: 99,
		}

		plSl := []*plamt{
			{
				Personality: 9001,
				Sound:       "ssssssssssssssssssss",
			},
			tree,
			{
				Personality: 8,
				Sound:       "bup",
			},
		}

		require.True(t, containerex.ContainsEq(plSl, tree))
		require.True(t, containerex.ContainsEq(plSl, &plamt{Personality: 101}))
		require.False(t, containerex.ContainsEq(plSl, &plamt{Personality: 8000, Sound: "ffff"}))
	})
}

type aminal struct {
	Personality uint64
	Sound       string
}

func (am *aminal) Eq(other *aminal) bool {
	return other != nil && am.Personality == other.Personality && am.Sound == other.Sound
}

type plamt struct {
	Personality uint64
	Sound       string
}

func (pl *plamt) Eq(other *plamt) bool {
	return other != nil && other.Sound == ""
}
