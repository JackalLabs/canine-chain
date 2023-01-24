package sha512

import "golang.org/x/crypto/blake2b"

// SHA512 is the SHA512 hashing method
type SHA512 struct{}

// New creates a new SHA512 hashing method
func New() *SHA512 {
	return &SHA512{}
}

// Hash generates a SHA512 hash from a byte array
func (h *SHA512) Hash(data []byte) []byte {
	hash := blake2b.Sum512(data)
	return hash[:]
}
