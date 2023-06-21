package keeper

import (
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

var _ types.QueryServer = Keeper{}
