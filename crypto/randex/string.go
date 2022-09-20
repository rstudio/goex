package randex

import (
	"crypto/rand"
	"math/big"
	"time"
)

const (
	defaultRandIntSleep = 100 * time.Nanosecond
)

// StringRunes returns the rune slice of characters used in String.
func StringRunes() []rune {
	return []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}

// String returns a random string of length n derived from the
// runes in StringRunes.
func String(n int) string {
	return StringOf(n, StringRunes())
}

// StringOf returns a random string of length n derived from the
// runes in letters.
func StringOf(n int, letters []rune) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = RuneOf(letters)
	}

	return string(b)
}

// RuneOf returns a random rune derived from the runes in letters.
// In the event that crypto/rand.Int returns an error, the
// operation will be retried after 100ns until successful. This
// means that this func will potentially block when entropy is
// unavailable.
func RuneOf(letters []rune) rune {
	max := big.NewInt(int64(len(letters)))

	for {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			time.Sleep(defaultRandIntSleep)
			continue
		}

		return letters[n.Int64()]
	}
}
