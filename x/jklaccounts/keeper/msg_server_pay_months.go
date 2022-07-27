package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/jklaccounts/types"
)

const (
	blocksPerMonth int64 = 100
	pricePerMonth  int64 = 8000000
)

func (k msgServer) PayMonths(goCtx context.Context, msg *types.MsgPayMonths) (*types.MsgPayMonthsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	account, found := k.GetAccounts(ctx, msg.Address)

	months, ok := sdk.NewIntFromString(msg.Months)
	if !ok {
		return nil, fmt.Errorf("month value is not a number")
	}

	monthcount := months.Int64()

	blocksToAdd := blocksPerMonth * monthcount

	cost := pricePerMonth * monthcount

	add, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, err
	}

	coin := sdk.NewCoin(msg.PaymentDenom, sdk.NewInt(cost))
	coins := sdk.NewCoins(coin)

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, add, types.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	currentBlock := ctx.BlockHeight()

	if !found {
		account = types.Accounts{
			Address:     msg.Address,
			Available:   "1",
			Used:        "0",
			ExpireBlock: fmt.Sprintf("%d", currentBlock),
		}
	}

	exp, ok := sdk.NewIntFromString(account.GetExpireBlock())
	if !ok {
		return nil, fmt.Errorf("error parsing expiration block")
	}

	newexp := exp.Int64() + blocksToAdd

	account.ExpireBlock = fmt.Sprintf("%d", newexp)

	k.SetAccounts(ctx, account)

	return &types.MsgPayMonthsResponse{}, nil
}
