package types

import "github.com/cosmos/cosmos-sdk/types/bech32"

const UFidPrefix = "jkluf"

func (u *UFID) ToString() (string, error) {
	return bech32.ConvertAndEncode(UFidPrefix, u.GetHash())
}
