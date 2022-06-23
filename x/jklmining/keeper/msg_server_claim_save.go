package keeper

import (
	"context"

	"crypto/sha256"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/jklmining/types"
)

func (k msgServer) ClaimSave(goCtx context.Context, msg *types.MsgClaimSave) (*types.MsgClaimSaveResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	savefile, _ := k.GetSaveRequests(
		ctx,
		msg.Saveindex,
	)

	if savefile.Approved == "true" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Already claimed.")
	}

	sum := sha256.Sum256([]byte(msg.Key))
	s := strings.ToUpper(fmt.Sprintf("%x", sum))
	i := strings.ToUpper(savefile.Index)
	if s != i {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s is not %s", s, i))
	}

	jak := k.Keeper.jklAccountsKeeper

	jaccount, found := jak.GetAccounts(ctx, msg.Creator)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "account not set up")
	}

	av, _ := sdk.NewIntFromString(jaccount.Available)
	us, _ := sdk.NewIntFromString(jaccount.Used)
	sz, _ := sdk.NewIntFromString(savefile.Size_)

	if av.Int64()-us.Int64() < sz.Int64() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "not enough space on account")
	}
	fsz, _ := sdk.NewIntFromString(savefile.Size_)
	fus := us.Int64() + fsz.Int64()

	jaccount.Used = fmt.Sprintf("%d", fus)

	jak.SetAccounts(ctx, jaccount)

	m := types.Mined{}

	m.Datasize = savefile.Size_
	m.Hash = savefile.Index
	m.Pcount = fmt.Sprintf("%d", ctx.BlockHeight())

	k.AppendMined(ctx, m)

	var saveRequests = types.SaveRequests{
		Creator:  savefile.Creator,
		Index:    savefile.Index,
		Size_:    savefile.Size_,
		Approved: "true",
	}

	k.SetSaveRequests(
		ctx,
		saveRequests,
	)

	clm := types.MinerClaims{}
	clm.Creator = msg.Creator
	clm.Hash = savefile.Index

	k.SetMinerClaims(ctx, clm)

	return &types.MsgClaimSaveResponse{}, nil
}
