package randex

import (
	"crypto/rand"
	"math/big"
	"time"
)

const (
	defaultRandIntSleep = 1 * time.Microsecond
)

func String(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letter))))
		if err != nil {
			time.Sleep(defaultRandIntSleep)
			continue
		}

		b[i] = letter[n.Int64()]
	}

	return string(b)
}
