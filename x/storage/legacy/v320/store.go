package v320

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

type ActiveDealSetter interface {
	SetActiveDeals(ctx sdk.Context, activeDeals types.ActiveDeals)
}

func MigrateStore(ctx sdk.Context, k ActiveDealSetter, sk sdk.StoreKey, codec codec.BinaryCodec) error {
	IterateLegacyActiveDeals(ctx, sk, codec, func(deal types.LegacyActiveDeals) bool {
		k.SetActiveDeals(ctx, types.ActiveDeals{
			Cid:          deal.Cid,
			Signee:       deal.Signee,
			Provider:     deal.Provider,
			Startblock:   deal.Startblock,
			Endblock:     deal.Endblock,
			Filesize:     deal.Filesize,
			LastProof:    ctx.BlockHeight(),
			Proofsmissed: deal.Proofsmissed,
			Blocktoprove: deal.Blocktoprove,
			Creator:      deal.Creator,
			Merkle:       deal.Merkle,
			Fid:          deal.Fid,
		})

		return false
	})

	return nil
}
