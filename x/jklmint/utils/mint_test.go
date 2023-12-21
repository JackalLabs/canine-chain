package utils_test

import (
	"github.com/jackalLabs/canine-chain/v3/x/jklmint/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOwedTokens(t *testing.T) {
	r := require.New(t)

	m := utils.GetTokensOwed(100, 6) // basic test
	r.Equal(int64(6), m)

	m = utils.GetTokensOwed(65000000, 6) // using 6 decimals
	r.Equal(int64(3900000), m)
}

func TestLastBlockMint(t *testing.T) {
	r := require.New(t)
	var bpy int64 = (365 * 24 * 60 * 60) / 6

	var dec int64 = 6

	m := utils.GetMintForBlock(4_200_000, bpy, dec)
	r.Equal(int64(4_199_999), m)

	m = utils.GetMintForBlock(m, bpy, dec)
	r.Equal(int64(4_199_998), m)
}
