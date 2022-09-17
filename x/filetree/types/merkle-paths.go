package types

import (
	"crypto/sha256"
	fmt "fmt"
	"strings"
)

func MerklePath(path string) string {
	chunks := strings.Split(path, "/")

	total := ""

	for _, chunk := range chunks {
		h := sha256.New()
		h.Write([]byte(chunk))
		b := fmt.Sprintf("%x", h.Sum(nil))
		k := fmt.Sprintf("%s%s", total, b)

		h = sha256.New()
		h.Write([]byte(k))
		total = fmt.Sprintf("%x", h.Sum(nil))
	}

	return total
}
