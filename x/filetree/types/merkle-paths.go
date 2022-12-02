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

	// If the end of the path has a slash, .e.g. "home/movies/", the .Split function would create a []String with length 3, i.e. the last element is " "
	// Cutting off the trailing '/' made everything compatible
	trimPath := strings.TrimSuffix(path, "/")
	chunks := strings.Split(trimPath, "/")

	total := ""

	for _, chunk := range chunks {
		h := sha256.New()
		h.Write([]byte(chunk))
		b := fmt.Sprintf("%x", h.Sum(nil))
		k := fmt.Sprintf("%s%s", total, b)

		h1 := sha256.New()
		h1.Write([]byte(k))
		total = fmt.Sprintf("%x", h1.Sum(nil))
	}

	return total
}
