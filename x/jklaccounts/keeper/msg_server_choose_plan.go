package keeper

import (
	"context"

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

	account.Available = msg.TbCount

	k.SetAccounts(ctx, account)

	return &types.MsgChoosePlanResponse{}, nil
}
