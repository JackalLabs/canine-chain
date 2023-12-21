package types_test

import (
	"github.com/jackalLabs/canine-chain/v3/x/jklmint/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNames(t *testing.T) {
	r := require.New(t)
	devAccount, err := types.GetDevAndGrantAccount()
	r.NoError(err)
	polAccount, err := types.GetPOLAccount()
	r.NoError(err)

	r.Equal("cosmos1v3jhvhmpdej97emjv9h8guc3fkme2", devAccount.String())
	r.Equal("cosmos1wpex7ar0vdhkchm0wahx2ezld35hzupf8ev", polAccount.String())
}
