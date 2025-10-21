package v4

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/jackalLabs/canine-chain/v5/x/storage/exported"
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
)

func MigrateStore(ctx sdk.Context, legacySubspace exported.Subspace, paramsSubspace *paramstypes.Subspace) error {
	var currParams LegacyParams
	legacySubspace.GetParamSet(ctx, &currParams)

	params := types.Params{
		DepositAccount:         currParams.DepositAccount,
		ProofWindow:            currParams.ProofWindow,
		ChunkSize:              currParams.ChunkSize,
		MissesToBurn:           currParams.MissesToBurn,
		PriceFeed:              currParams.PriceFeed,
		MaxContractAgeInBlocks: currParams.MaxContractAgeInBlocks,
		PricePerTbPerMonth:     currParams.PricePerTbPerMonth,
		AttestFormSize:         5,
		AttestMinToPass:        3,
		CollateralPrice:        10_000_000_000,
	}

	if err := params.Validate(); err != nil {
		return err
	}

	paramsSubspace.SetParamSet(ctx, &params)

	return nil
}
