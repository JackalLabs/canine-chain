package keeper

import (
	"context"
	"encoding/hex"

	"github.com/jackalLabs/canine-chain/v3/x/storage/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerr "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func MakeAddress() (sdk.AccAddress, error) {
	oracleAddressString := "storage_oracle"
	oracleAddressHex := hex.EncodeToString([]byte(oracleAddressString))
	return sdk.AccAddressFromHex(oracleAddressHex)
}

func (k msgServer) RequestChunk(goCtx context.Context, msg *types.MsgRequestChunk) (*types.MsgRequestChunkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bid := msg.Bid
	bidAsCoins := sdk.NewCoins(bid)

	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerr.Wrapf(err, "cannot get address from %s", msg.Creator)
	}

	oracleAddress, err := MakeAddress()
	if err != nil {
		return nil, sdkerr.Wrap(err, "could not make oracle address")
	}

	err = k.bankkeeper.SendCoinsFromAccountToModule(ctx, creatorAddress, types.ModuleName, bidAsCoins)
	if err != nil {
		return nil, sdkerr.Wrap(err, "user does not have enough coins")
	}
	err = k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, oracleAddress, bidAsCoins)
	if err != nil {
		return nil, sdkerr.Wrap(err, "somehow the module itself ate all our coins, we need to fail here")
	}

	req := types.OracleRequest{
		Requester: msg.Creator,
		Merkle:    msg.Merkle,
		Chunk:     msg.Chunk,
		Bid:       bid,
	}

	k.SetOracleRequest(ctx, req)

	return &types.MsgRequestChunkResponse{}, nil
}

func (k msgServer) FulfillRequest(goCtx context.Context, msg *types.MsgFulfillRequest) (*types.MsgFulfillRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	verified := utils.VerifyProof(msg.Merkle, msg.HashList, msg.Chunk, msg.Data)
	if !verified {
		return nil, sdkerr.Wrap(sdkerr.ErrUnauthorized, "cannot verify data against merkle")
	}

	req, found := k.GetOracleRequest(ctx, msg.Requester, msg.Merkle, msg.Chunk)
	if !found {
		return nil, sdkerr.Wrap(sdkerr.ErrNotFound, "cannot find oracle request")
	}

	bid := req.Bid
	bidAsCoins := sdk.NewCoins(bid)

	oracleAddress, err := MakeAddress()
	if err != nil {
		return nil, sdkerr.Wrap(err, "could not make oracle address")
	}
	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerr.Wrapf(err, "could not make user address from %s", msg.Creator)
	}

	err = k.bankkeeper.SendCoinsFromAccountToModule(ctx, oracleAddress, types.ModuleName, bidAsCoins)
	if err != nil {
		return nil, sdkerr.Wrap(err, "oracle somehow does not have enough coins")
	}
	err = k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAddress, bidAsCoins)
	if err != nil {
		return nil, sdkerr.Wrap(err, "somehow the module itself ate all our coins, we need to fail here")
	}

	entry := types.OracleEntry{
		Owner:  msg.Requester,
		Merkle: msg.Merkle,
		Chunk:  msg.Chunk,
		Data:   msg.Data,
	}

	k.SetOracleEntry(ctx, entry)

	k.RemoveOracleRequest(ctx, msg.Requester, msg.Merkle, msg.Chunk)

	return &types.MsgFulfillRequestResponse{}, nil
}
