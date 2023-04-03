package keeper

import (
	"sort"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/jackalLabs/canine-chain/x/amm/types"
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
func (k Keeper) MintAndSendPToken(
	ctx sdk.Context,
	pool types.Pool,
	toAddr sdk.AccAddress,
	amount sdk.Int) error {

	poolToken := sdk.NewCoin(pool.PoolToken.Denom, amount)
	poolTokens := sdk.NewCoins(poolToken)

	sdkError := k.bankKeeper.MintCoins(ctx, types.ModuleName, poolTokens)

	if sdkError != nil {
		return sdkError
	}

	sdkError = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddr, poolTokens)

	if sdkError != nil {
		return sdkError
	}

	return nil
}

// Returns a Pool with passed values.
// It does not validate the message.
func (k Keeper) NewPool(ctx sdk.Context, msg *types.MsgCreatePool) types.Pool {

	coins := sdk.NewCoins(msg.Coins...)
	poolName := generatePoolName(coins)

	var pool = types.Pool{
		Index:           poolName,
		Name:            poolName,
		Coins: coins,
		AMM_Id:          msg.AmmId,
		SwapFeeMulti:    msg.SwapFeeMulti,
		MinLockDuration: msg.MinLockDuration,
		PenaltyMulti:    msg.PenaltyMulti,
		PoolToken: sdk.NewCoin("", sdk.ZeroInt()),
	}

	return pool
}

// SetPool set a specific lPool in the store from its index
func (k Keeper) SetPool(ctx sdk.Context, lPool types.Pool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))
	b := k.cdc.MustMarshal(&lPool)
	store.Set(types.PoolKey(
		lPool.Index,
	), b)
}

// GetPool returns a Pool from its index
func (k Keeper) GetPool(
	ctx sdk.Context,
	index string,

) (val types.Pool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))

	b := store.Get(types.PoolKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePool removes a lPool from the store
func (k Keeper) RemovePool(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))
	store.Delete(types.PoolKey(
		index,
	))
}

// GetAllPool returns all lPool
func (k Keeper) GetAllPool(ctx sdk.Context) (list []types.Pool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Pool
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

// Generate a denom unit for PToken with itself as smallest unit.
func generatePTokenDenomUnit(denom string, aliase string) []*banktypes.DenomUnit {
	// More info about denom units:
	// https://pkg.go.dev/github.com/cosmos/cosmos-sdk@v0.46.0/x/bank/types#DenomUnit
	tokenDenomUnit := banktypes.DenomUnit{
		Denom:    denom,
		Exponent: 0,
	}
	return []*banktypes.DenomUnit{&tokenDenomUnit}
}

// Register pool token to bank module
func (k Keeper) registerPoolToken(ctx sdk.Context, denom string) {
	_, found := k.bankKeeper.GetDenomMetaData(ctx, denom)

	aliase := "JKLP"

	if !found {
		// Register PoolTokenDenom meta data.
		// Step 1: generate denom units for PoolToken.
		denomUnits := generatePTokenDenomUnit(denom, aliase)

		// Step 2: add it to bank's denom meta data store.
		metaData := banktypes.Metadata{
			Description: "Jackal liquidity pool token",
			DenomUnits:  denomUnits,
			Base:        denomUnits[0].Denom,
		}

		k.bankKeeper.SetDenomMetaData(ctx, metaData)
	}
}
