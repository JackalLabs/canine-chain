package v320

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func MigrateStore(ctx sdk.Context, ps *paramstypes.Subspace) error {
	var params types.Params
	ps.GetParamSet(ctx, &params)
	params.ProofWindow *= 2 // doubling proof window to freshen up the network a bit
	ps.SetParamSet(ctx, &params)

	return nil
}
