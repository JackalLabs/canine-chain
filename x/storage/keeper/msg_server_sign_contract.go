package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/storage/types"
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

	if contract.Signer != msg.Creator {
		return nil, fmt.Errorf("you do not have permission to approve this contract")
	}

	size := sdk.NewInt(contract.FileSize)

	pieces := size.Quo(sdk.NewInt(k.GetParams(ctx).ChunkSize))

	var pieceToStart int64

	if !pieces.IsZero() {
		pieceToStart = ctx.BlockHeight() % pieces.Int64()
	}

	var end int64
	if msg.PayOnce {
		s := size.Quo(sdk.NewInt(1_000_000_000)).Int64()
		if s <= 0 {
			s = 1
		}
		cost := k.GetStorageCost(ctx, s, 720*12*200) // pay for 200 years
		deposit, err := sdk.AccAddressFromBech32(k.GetParams(ctx).DepositAccount)
		if err != nil {
			return nil, err
		}
		err = k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, deposit, sdk.NewCoins(sdk.NewCoin("ujkl", cost)))
		if err != nil {
			return nil, err
		}

		end = (200*31_536_000)/6 + ctx.BlockHeight()
	}

	deal := types.ActiveDealsV2{
		Cid:           contract.Cid,
		Signer:        contract.Signer,
		Provider:      contract.Creator,
		StartBlock:    ctx.BlockHeight(),
		EndBlock:      end,
		FileSize:      contract.FileSize,
		ProofVerified: false,
		BlockToProve:  pieceToStart,
		Creator:       msg.Creator,
		ProofsMissed:  0,
		Merkle:        contract.Merkle,
		Fid:           contract.Fid,
		DealVersion:   0,
	}

	if end == 0 {
		payInfo, found := k.GetStoragePaymentInfo(ctx, msg.Creator)
		if !found {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "payment info not found, please purchase storage space")
		}

		// check if user has any free space
		if (payInfo.SpaceUsed + (contract.FileSize * 3)) > payInfo.SpaceAvailable {
			return nil, fmt.Errorf("not enough storage space")
		}
		// check if storage subscription still active
		if payInfo.End.Before(ctx.BlockTime()) {
			return nil, fmt.Errorf("storage subscription has expired")
		}

		payInfo.SpaceUsed += contract.FileSize * 3

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

		newContract := types.StrayV2{
			Cid:      scid,
			Signer:   contract.Signer,
			Fid:      contract.Fid,
			FileSize: contract.FileSize,
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

	return &types.MsgSignContractResponse{}, nil
}
