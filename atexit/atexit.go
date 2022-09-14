package atexit

import (
	"fmt"
	"os"
	"sync"
)

var (
	atExitFuncs = []*namedFunc{}
	atExitLock  = &sync.Mutex{}
)

type namedFunc struct {
	name string
	f    func()
}

func Register(name string, f func()) {
	atExitLock.Lock()
	defer atExitLock.Unlock()

	atExitFuncs = append(atExitFuncs, &namedFunc{name: name, f: f})
}

func Unregister(name string) bool {
	atExitLock.Lock()
	defer atExitLock.Unlock()

	newValue := []*namedFunc{}

	seen := false

	for _, nf := range atExitFuncs {
		if nf.name != name {
			newValue = append(newValue, nf)
			continue
		}

		seen = true
	}

	atExitFuncs = newValue

	return seen
}

func Run() {
	atExitLock.Lock()
	defer atExitLock.Unlock()

	for _, nf := range atExitFuncs {
		func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Fprintf(os.Stderr, "ERROR: in atexit func: %[1]v\n", err)
				}
			}()

			nf.f()
		}()
	}

	atExitFuncs = []*namedFunc{}
}
