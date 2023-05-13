package zk_test

import (
	"testing"

	"github.com/jackalLabs/canine-chain/x/storage/zk"
	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	r := require.New(t)

	data := "secret_message"

	wp, err := zk.HashData([]byte(data))

	r.NoError(err)

	verified := zk.VerifyHash(wp)

	r.Equal(true, verified)
}
