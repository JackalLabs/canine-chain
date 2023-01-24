package simulation

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

func StringWithCharset(r *rand.Rand, length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}
