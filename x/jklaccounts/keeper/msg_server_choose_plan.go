package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/jklaccounts/types"
)

func (k msgServer) ChoosePlan(goCtx context.Context, msg *types.MsgChoosePlan) (*types.MsgChoosePlanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	account, found := k.GetAccounts(ctx, msg.Creator)

	if !found {
		account = types.Accounts{
			Address:     msg.Creator,
			Available:   msg.TbCount,
			Used:        "0",
			ExpireBlock: "0",
		}
	}

	expir, ok := sdk.NewIntFromString(account.ExpireBlock)
	if !ok {
		return nil, fmt.Errorf("cannot make int from string: %s", account.ExpireBlock)
	}
	blocksBetween := expir.Int64() - ctx.BlockHeight()

	if blocksBetween >= blocksPerMonth {
		months := blocksBetween / blocksPerMonth

		price := pricePerMonth * months

		add, err := sdk.AccAddressFromBech32(msg.Creator)

		if err != nil {
			return nil, err
		}

		coin := sdk.NewCoin(msg.PaymentDenom, sdk.NewInt(price))
		coins := sdk.NewCoins(coin)

		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, add, types.ModuleName, coins)
		if err != nil {
			return nil, err
		}
	}

	account.Available = msg.TbCount

	k.SetAccounts(ctx, account)

	return &types.MsgChoosePlanResponse{}, nil
}
