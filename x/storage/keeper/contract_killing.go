package keeper

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k Keeper) KillOldContracts(ctx sdk.Context) {
	maxContractAgeInBlocks := k.GetParams(ctx).MaxContractAgeInBlocks
	contracts := k.GetAllContracts(ctx)

	for _, contract := range contracts {
		if contract.Age+maxContractAgeInBlocks < ctx.BlockHeight() {
			k.RemoveContracts(ctx, contract.Cid)
		}
	}
}

func (k Keeper) DropDeal(ctx sdk.Context, deal types.ActiveDeals, burn bool) error {
	intBlock, ok := sdk.NewIntFromString(deal.Endblock)
	if !ok {
		return sdkerror.Wrapf(sdkerror.ErrInvalidType, "int parse failed for endblock")
	}
	// Creating new stray file from the burned active deal
	strayDeal := types.Strays{
		Cid:      deal.Cid,
		Fid:      deal.Fid,
		Signee:   deal.Signee,
		Filesize: deal.Filesize,
		Merkle:   deal.Merkle,
		End:      intBlock.Int64(),
	}
	k.SetStrays(ctx, strayDeal)
	k.RemoveActiveDeals(ctx, deal.Cid)

	if burn {
		provider, found := k.GetProviders(ctx, deal.Provider)
		if !found {
			return nil
		}

		burnString := provider.BurnedContracts
		burns, err := strconv.ParseInt(burnString, 10, 64)
		if err != nil {
			return err
		}
		burns++

		provider.BurnedContracts = fmt.Sprintf("%d", burns)
		k.SetProviders(ctx, provider)
	}

	return nil
}
