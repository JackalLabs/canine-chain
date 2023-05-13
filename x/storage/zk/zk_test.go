package zk_test

import (
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

	wp, err := zk.HashData([]byte(data), ccs)
	r.NoError(err)

	verified := zk.VerifyHash(wp)

	fmt.Println(wp)

	r.Equal(true, verified)
}
