package zk_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/jackalLabs/canine-chain/x/storage/zk"
	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	r := require.New(t)

	data := "secret_message"

	ccs, err := zk.GetCircuit()
	r.NoError(err)

	wp, hash, err := zk.HashData([]byte(data), ccs)
	r.NoError(err)

	verified := zk.VerifyHash(wp, hash)

	r.Equal(true, verified)
}

func TestMultiHash(t *testing.T) {
	r := require.New(t)
	ccs, err := zk.GetCircuit()
	r.NoError(err)

	data := []string{"secret_message", "secret_message_2", "secret_message_3"}
	hashes := make([][]byte, len(data))
	for i, datum := range data {
		wp, hash, err := zk.HashData([]byte(datum), ccs)
		r.NoError(err)

		verified := zk.VerifyHash(wp, hash)

		r.Equal(true, verified)

		hashes[i] = hash
	}
	r.Equal(false, bytes.Equal(hashes[0], hashes[1]) || bytes.Equal(hashes[0], hashes[2]))
}

func TestEncodeDecode(t *testing.T) {
	r := require.New(t)
	data := "secret_message"
	ccs, err := zk.GetCircuit()
	r.NoError(err)
	wp, hash, err := zk.HashData([]byte(data), ccs)
	r.NoError(err)

	wpenc, err := wp.Encode()
	r.NoError(err)

	fmt.Println(wpenc)
	fmt.Println(len(wpenc.VerifyingKey) + len(wpenc.Proof))

	wpnew, err := zk.Decode(wpenc)
	r.NoError(err)

	verified := zk.VerifyHash(wpnew, hash)

	r.Equal(true, verified)
}
