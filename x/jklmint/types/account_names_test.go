package types_test

import (
	"testing"

	"github.com/jackalLabs/canine-chain/v5/x/jklmint/types"
	"github.com/stretchr/testify/require"
)

func TestNames(t *testing.T) {
	r := require.New(t)
	devAccount, err := types.GetDevAndGrantAccount()
	r.NoError(err)

	r.Equal("cosmos1v3jhvhmpdej97emjv9h8guc3fkme2", devAccount.String())
}
