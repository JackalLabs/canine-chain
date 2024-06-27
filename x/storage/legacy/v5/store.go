package v5

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/exported"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func MigrateStore(ctx sdk.Context, legacySubspace exported.Subspace, paramsSubspace *paramstypes.Subspace) error {
	ctx.Logger().Info("Migrating params to v4...")

	var currParams v4Params
	legacySubspace.GetParamSet(ctx, &currParams)

	params := types.Params{
		DepositAccount:         currParams.DepositAccount,
		ProofWindow:            currParams.ProofWindow,
		ChunkSize:              currParams.ChunkSize,
		MissesToBurn:           currParams.MissesToBurn,
		PriceFeed:              currParams.PriceFeed,
		MaxContractAgeInBlocks: currParams.MaxContractAgeInBlocks,
		PricePerTbPerMonth:     15,
		AttestFormSize:         5,
		AttestMinToPass:        3,
		CollateralPrice:        10_000_000_000,
		CheckWindow:            100,
		ReferralCommission:     25,
		PolRatio:               40,
	}

	if err := params.Validate(); err != nil {
		return err
	}

	paramsSubspace.SetParamSet(ctx, &params)

	return nil
}
