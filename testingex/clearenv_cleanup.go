package testingex

import (
	"os"
	"strings"
)

func ClearenvCleanup() func() {
	beginEnv := os.Environ()
	os.Clearenv()

	return func() {
		os.Clearenv()

		for _, pair := range beginEnv {
			parts := strings.SplitN(pair, "=", 2)

			if len(parts) < 2 {
				continue
			}

			os.Setenv(parts[0], parts[1])
		}
	}
}
