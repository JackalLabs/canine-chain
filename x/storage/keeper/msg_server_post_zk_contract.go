package keeper

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k msgServer) PostZKContract(goCtx context.Context, msg *types.MsgPostZKContract) (*types.MsgPostZKContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, ok := k.GetProviders(ctx, msg.Creator)
	if !ok {
		return nil, fmt.Errorf("can't find provider")
	}

	ts, ok := sdk.NewIntFromString(provider.Totalspace)

	if !ok {
		return nil, fmt.Errorf("error parsing total space")
	}

	if k.GetProviderUsing(ctx, msg.Creator)+msg.FileSize > ts.Int64() {
		return nil, fmt.Errorf("not enough space on provider")
	}

	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%s%s%s", msg.Signer, msg.Creator, msg.Fid))
	if err != nil {
		return nil, err
	}
	hashName := h.Sum(nil)

	cid, err := MakeCid(hashName)
	if err != nil {
		return nil, err
	}

	_, cidtaken := k.GetContracts(ctx, cid)
	if cidtaken {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "cannot post the same contract twice")
	}

	newContract := types.ContractV2{
		Cid:      cid,
		Signer:   msg.Signer,
		Fid:      msg.Fid,
		FileSize: msg.FileSize,
		Creator:  msg.Creator,
		Merkle:   msg.Merkle,
		Age:      ctx.BlockHeight(),
	}

	k.SetContracts(ctx, newContract)

	return &types.MsgPostZKContractResponse{}, nil
}
