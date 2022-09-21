package types

import (
	"crypto/sha256"
	fmt "fmt"
	"strings"
)

func AddToMerkle(path string, append string) string {
	total := path

	k := fmt.Sprintf("%s%s", total, append)

	h := sha256.New()
	h.Write([]byte(k))
	total = fmt.Sprintf("%x", h.Sum(nil))
	return total
}

func MerklePath(path string) string { // ex: hello/world/path -> ["hello", "world", "path"] -> 3867baa2724c672442e4ba21b6fa532a6380d06a2f8779f11d626bd840d1cdee
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
