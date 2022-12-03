package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k msgServer) SignContract(goCtx context.Context, msg *types.MsgSignContract) (*types.MsgSignContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contract, found := k.GetContracts(ctx, msg.Cid)
	if !found {
		return nil, fmt.Errorf("contract not found")
	}

	if contract.Signee != msg.Creator {
		return nil, fmt.Errorf("you do not have permission to approve this contract")
	}

	size, ok := sdk.NewIntFromString(contract.Filesize)
	if !ok {
		return nil, fmt.Errorf("cannot parse size")
	}

	pieces := size.Quo(sdk.NewInt(1024))

	deal := types.ActiveDeals{
		Cid:           contract.Cid,
		Signee:        contract.Signee,
		Provider:      contract.Creator,
		Startblock:    fmt.Sprintf("%d", ctx.BlockHeight()),
		Endblock:      fmt.Sprintf("%d", ctx.BlockHeight()),
		Filesize:      contract.Filesize,
		Proofverified: "false",
		Blocktoprove:  fmt.Sprintf("%d", ctx.BlockHeight()%pieces.Int64()),
		Creator:       msg.Creator,
		Proofsmissed:  "0",
		Merkle:        contract.Merkle,
		Fid:           contract.Fid,
	}

	usage, found := k.GetClientUsage(ctx, msg.Creator)
	if !found {
		usage = types.ClientUsage{
			Address: msg.Creator,
			Usage:   "0",
		}
	}

	size, ok := sdk.NewIntFromString(contract.Filesize)
	if !ok {
		return nil, fmt.Errorf("cannot parse filesize")
	}

	used, ok := sdk.NewIntFromString(usage.Usage)
	if !ok {
		return nil, fmt.Errorf("cannot parse usage")
	}

	usage.Usage = fmt.Sprintf("%d", used.Int64()+size.Int64())

	k.SetClientUsage(ctx, usage)
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
		_, err := io.WriteString(h, fmt.Sprintf("%s%s%d", contract.Creator, contract.Fid, i))
		if err != nil {
			return nil, err
		}
		hashName := h.Sum(nil)

		scid := fmt.Sprintf("%x", hashName)

		newContract := types.Strays{
			Cid:      scid,
			Signee:   contract.Signee,
			Fid:      contract.Fid,
			Filesize: contract.Filesize,
			Merkle:   contract.Merkle,
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
