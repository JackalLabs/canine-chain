package sha512

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// _byteArray is a helper to turn a string in to a byte array
func _byteArray(input string) []byte {
	x, _ := hex.DecodeString(input)
	return x
}

func TestHash(t *testing.T) {
	tests := []struct {
		data   []byte
		output []byte
	}{
		{
			data:   _byteArray("e9e0083e456539e9f6336164cd98700e668178f98af147ef750eb90afcf2f637"),
			output: _byteArray("9744850fc1d693a3cba541f9367a8eb4c736bcd24dc97db3e4d2c9e99c771fdc8cff9ae752eaa99ca969def7d5c38a844ce55edb9b12b9a9408c62732a59ce6b"),
		},
	}

	hash := New()
	for i, test := range tests {
		output := hash.Hash(test.data)
		assert.Equal(t, test.output, output, fmt.Sprintf("failed at test %d", i))
	}
}
