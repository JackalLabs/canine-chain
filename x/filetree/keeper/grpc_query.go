package keeper

import (
	"github.com/jackalLabs/canine-chain/v5/x/filetree/types"
)

var _ types.QueryServer = Keeper{}
