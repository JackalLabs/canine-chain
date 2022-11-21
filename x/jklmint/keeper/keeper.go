package keeper

import (
	"fmt"
	"math/big"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/jackalLabs/canine-chain/x/jklmint/types"
)

type (
	Keeper struct {
		cdc              codec.BinaryCodec
		storeKey         storetypes.StoreKey
		paramSpace       paramtypes.Subspace
		stakingKeeper    types.StakingKeeper
		bankKeeper       types.BankKeeper
		feeCollectorName string
		// miningName       string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	paramSpace paramtypes.Subspace,
	sk types.StakingKeeper,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	feeCollectorName string,
	// miningName string,
) Keeper {
	// ensure mint module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the mint module account has not been set")
	}

	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:              cdc,
		storeKey:         key,
		paramSpace:       paramSpace,
		stakingKeeper:    sk,
		bankKeeper:       bk,
		feeCollectorName: feeCollectorName,
		// miningName:       miningName,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

func FloatToBigInt(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result)

	return result
}

func (k Keeper) GetInflation(ctx sdk.Context) (sdk.Dec, error) {
	coins := k.bankKeeper.GetSupply(ctx, k.GetParams(ctx).MintDenom)
	zeroDec, err := sdk.NewDecFromStr("0")
	if err != nil {
		return zeroDec, types.ErrCannotParseFloat
	}

	amt := coins.Amount.ToDec()
	famt, err := amt.Float64()
	if err != nil {
		return zeroDec, types.ErrCannotParseFloat
	}

	var tokens float64 = 4 // TODO: Update to 10 when storage goes live
	highDec, err := sdk.NewDecFromStr("1.0")
	if err != nil {
		return zeroDec, types.ErrCannotParseFloat
	}

	if amt.IsZero() {
		return sdk.NewDec(0), nil
	}

	var blocksPerYearEstiamte int64 = (365 * 24 * 60 * 60) / 6

	ratio := tokens / famt

	inflate := sdk.NewDec(printedPerYear)

	ratioSDK := sdk.NewDecFromBigInt(ratioDec)

	return ratioSDK, nil
}

// AddCollectedFees implements an alias call to the underlying supply keeper's
// AddCollectedFees to be used in BeginBlocker.
func (k Keeper) AddCollectedFees(ctx sdk.Context, fees sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, fees)
}
