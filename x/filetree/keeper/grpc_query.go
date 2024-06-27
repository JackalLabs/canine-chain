package keeper

import (
	"github.com/jackalLabs/canine-chain/v4/x/filetree/types"
)

var _ types.QueryServer = Keeper{}
