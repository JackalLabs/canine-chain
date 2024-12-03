package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/jklmint/types"
)

func (k Keeper) MintedTokens(c context.Context, req *types.QueryMintedTokens) (*types.QueryMintedTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	minted, _ := k.GetMintedBlock(ctx, req.Block)
	return &types.QueryMintedTokensResponse{Tokens: minted.Minted}, nil
}
