package atexit_test

import (
	"sync"
	"testing"

	"github.com/rstudio/goex/atexit"
	"github.com/stretchr/testify/require"
)

func TestAtExit(t *testing.T) {
	r := require.New(t)

	state := []string{}
	stateLock := &sync.Mutex{}

	atexit.Register("boop", func() {
		stateLock.Lock()
		defer stateLock.Unlock()

		state = append(state, "booped")
	})

	atexit.Register("scared", func() {
		panic("oh no")
	})

	atexit.Run()
	atexit.Run()

	r.Contains(state, "booped")
	r.Len(state, 1)

	state = []string{}

	atexit.Register("boop2", func() {
		stateLock.Lock()
		defer stateLock.Unlock()

		state = append(state, "booped")
	})

	atexit.Register("scared2", func() {
		panic("oh no")
	})

	r.True(atexit.Unregister("boop2"))
	r.False(atexit.Unregister("boop"))

	atexit.Run()
	atexit.Run()

	r.NotContains(state, "booped")
	r.Len(state, 0)
}
