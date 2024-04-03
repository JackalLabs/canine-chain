package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k msgServer) SignContract(goCtx context.Context, msg *types.MsgSignContract) (*types.MsgSignContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contract, found := k.GetContracts(ctx, msg.Cid)
	if !found {
		return nil, fmt.Errorf("contract not found")
	}

	_, found = k.GetActiveDeals(ctx, msg.Cid)
	if found {
		return nil, fmt.Errorf("contract already exists")
	}

	_, found = k.GetStrays(ctx, msg.Cid)
	if found {
		return nil, fmt.Errorf("contract already exists")
	}

	if contract.Signee != msg.Creator {
		return nil, fmt.Errorf("you do not have permission to approve this contract")
	}

	size, ok := sdk.NewIntFromString(contract.Filesize)
	if !ok {
		return nil, fmt.Errorf("cannot parse size")
	}

	pieces := size.Quo(sdk.NewInt(k.GetParams(ctx).ChunkSize))

	var pieceToStart int64

	if !pieces.IsZero() {
		pieceToStart = ctx.BlockHeight() % pieces.Int64()
	}

	var end int64
	if msg.PayOnce {

		s := size.Quo(sdk.NewInt(1_000_000)).Int64() // round to mbs
		if s <= 0 {
			s = 1
		}
		cost := k.GetStorageCostKbs(ctx, s*1000, 720*12*200) // pay for 200 years in mbs

		deposit, err := sdk.AccAddressFromBech32(k.GetParams(ctx).DepositAccount)
		if err != nil {
			return nil, err
		}

		senderAddress, err := sdk.AccAddressFromBech32(msg.Creator)
		if err != nil {
			return nil, err
		}
		costCoins := sdk.NewCoins(sdk.NewCoin("ujkl", cost))
		err = k.bankkeeper.SendCoinsFromAccountToModule(ctx, senderAddress, types.ModuleName, costCoins)
		if err != nil {
			return nil, err
		}
		err = k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, deposit, costCoins)
		if err != nil {
			return nil, err
		}

		end = (200*31_536_000)/6 + ctx.BlockHeight()
	}

	deal := types.ActiveDeals{
		Cid:          contract.Cid,
		Signee:       contract.Signee,
		Provider:     contract.Creator,
		Startblock:   fmt.Sprintf("%d", ctx.BlockHeight()),
		Endblock:     fmt.Sprintf("%d", end),
		Filesize:     contract.Filesize,
		LastProof:    ctx.BlockHeight(),
		Blocktoprove: fmt.Sprintf("%d", pieceToStart),
		Creator:      msg.Creator,
		Proofsmissed: "0",
		Merkle:       contract.Merkle,
		Fid:          contract.Fid,
	}

	if end == 0 {
		fsize, ok := sdk.NewIntFromString(contract.Filesize)
		if !ok {
			return nil, fmt.Errorf("cannot parse file size")
		}
		payInfo, found := k.GetStoragePaymentInfo(ctx, msg.Creator)
		if !found {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "payment info not found, please purchase storage space")
		}

		// check if user has any free space
		if (payInfo.SpaceUsed + (fsize.Int64() * 3)) > payInfo.SpaceAvailable {
			return nil, fmt.Errorf("not enough storage space")
		}
		// check if storage subscription still active
		if payInfo.End.Before(ctx.BlockTime()) {
			return nil, fmt.Errorf("storage subscription has expired")
		}

		payInfo.SpaceUsed += fsize.Int64() * 3

		k.SetStoragePaymentInfo(ctx, payInfo)
	}

	k.SetActiveDeals(ctx, deal)
	k.RemoveContracts(ctx, contract.Cid)

	ftc, found := k.GetFidCid(ctx, contract.Fid)

	cids := []string{contract.Cid}

	if found {
		var ncids []string
		err := json.Unmarshal([]byte(ftc.Cids), &ncids)
		if err != nil {
			return nil, err
		}

		cids = append(cids, ncids...)
	}

	for i := 0; i < 2; i++ {
		h := sha256.New()
		_, err := io.WriteString(h, fmt.Sprintf("%s%d", contract.Cid, i))
		if err != nil {
			return nil, err
		}
		hashName := h.Sum(nil)

		scid, err := MakeCid(hashName)
		if err != nil {
			return nil, err
		}

		newContract := types.Strays{
			Cid:      scid,
			Signee:   contract.Signee,
			Fid:      contract.Fid,
			Filesize: contract.Filesize,
			Merkle:   contract.Merkle,
			End:      end,
		}

		cids = append(cids, scid)

		k.SetStrays(ctx, newContract)

	}

	cidarr, err := json.Marshal(cids)
	if err != nil {
		return nil, err
	}

	nftc := types.FidCid{
		Fid:  contract.Fid,
		Cids: string(cidarr),
	}

	k.SetFidCid(ctx, nftc)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	b := "false"
	if msg.PayOnce {
		b = "true"
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSignContract,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyContract, msg.Cid),
			sdk.NewAttribute(types.AttributeKeyPayOnce, b),
		),
	)

	return &types.MsgSignContractResponse{}, nil
}
