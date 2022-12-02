package keeper

import (
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

var _ types.QueryServer = Keeper{}
