package keeper

import (
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

var _ types.QueryServer = Keeper{}
