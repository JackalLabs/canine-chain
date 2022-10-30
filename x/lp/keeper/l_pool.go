package keeper

import (
	"sort"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/jackalLabs/canine-chain/x/lp/types"
)

func SortDenoms(denoms []string) []string {
	sort.Strings(denoms)
	return denoms
}

// Return unsorted list of coin denoms
func GetAllDenoms(coins sdk.Coins) []string {
	var denoms []string

	for _, coin := range coins {
		denoms = append(denoms, coin.GetDenom())
	}

	return denoms
}

// Mint liquidity token and send it to toAddr.
// Returns error when minting or sending fails.
func (k Keeper) MintAndSendLPToken(
	ctx sdk.Context,
	pool types.LPool,
	toAddr sdk.AccAddress,
	amount sdk.Int,
) error {
	lPToken := sdk.NewCoin(pool.LptokenDenom, amount)
	lPTokens := sdk.NewCoins(lPToken)

	sdkError := k.bankKeeper.MintCoins(ctx, types.ModuleName, lPTokens)

	if sdkError != nil {
		return sdkError
	}

	sdkError = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddr, lPTokens)

	if sdkError != nil {
		return sdkError
	}

	return nil
}

// Returns a LPool with passed values.
// It does not validate the message.
func (k Keeper) NewLPool(ctx sdk.Context, msg *types.MsgCreateLPool) types.LPool {
	normCoins := sdk.NormalizeCoins(msg.Coins)
	poolName := generatePoolName(normCoins)

	pool := types.LPool{
		Index:           poolName,
		Name:            poolName,
		Coins:           normCoins,
		AMM_Id:          msg.Amm_Id,
		SwapFeeMulti:    msg.SwapFeeMulti,
		MinLockDuration: msg.MinLockDuration,
		PenaltyMulti:    msg.PenaltyMulti,
		// NOTE: use chain token alias
		LptokenDenom:   generatePoolName(normCoins) + "-JKL",
		LPTokenBalance: "",
	}

	return pool
}

// SetLPool set a specific lPool in the store from its index
func (k Keeper) SetLPool(ctx sdk.Context, lPool types.LPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LPoolKeyPrefix))
	b := k.cdc.MustMarshal(&lPool)
	store.Set(types.LPoolKey(
		lPool.Index,
	), b)
}

// GetLPool returns a lPool from its index
func (k Keeper) GetLPool(
	ctx sdk.Context,
	index string,
) (val types.LPool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LPoolKeyPrefix))

	b := store.Get(types.LPoolKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLPool removes a lPool from the store
func (k Keeper) RemoveLPool(
	ctx sdk.Context,
	index string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LPoolKeyPrefix))
	store.Delete(types.LPoolKey(
		index,
	))
}

// GetAllLPool returns all lPool
func (k Keeper) GetAllLPool(ctx sdk.Context) (list []types.LPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LPoolKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LPool
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Generate pool name from coins.
func generatePoolName(coins sdk.Coins) string {
	denoms := GetAllDenoms(coins)

	return strings.Join(SortDenoms(denoms), "-")
}

// Generate a denom unit for LPToken with itself as smallest unit.
func generateLPTokenDenomUnit(denom string, aliase string) []*banktypes.DenomUnit {
	// More info about denom units:
	// https://pkg.go.dev/github.com/cosmos/cosmos-sdk@v0.46.0/x/bank/types#DenomUnit
	tokenDenomUnit := banktypes.DenomUnit{
		Denom:    denom,
		Exponent: 0,
	}
	return []*banktypes.DenomUnit{&tokenDenomUnit}
}

func (k Keeper) registerLPToken(ctx sdk.Context, denom string) {
	_, found := k.bankKeeper.GetDenomMetaData(ctx, denom)

	aliase := "JKLLP"

	if !found {
		// Register LPTokenDenom meta data.
		// Step 1: generate denom units for LPToken.
		denomUnits := generateLPTokenDenomUnit(denom, aliase)

		// Step 2: add it to bank's denom meta data store.
		metaData := banktypes.Metadata{
			Description: "Jackal liquidity pool token",
			DenomUnits:  denomUnits,
			Base:        denomUnits[0].Denom,
		}

		k.bankKeeper.SetDenomMetaData(ctx, metaData)
	}
}
