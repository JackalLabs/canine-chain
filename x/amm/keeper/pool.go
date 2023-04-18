package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/jackalLabs/canine-chain/x/amm/types"
)

// Mint pool token and send to an account
// Returns error when minting or sending fails.
func (k Keeper) MintAndSendPoolToken(
	ctx sdk.Context,
	pool types.Pool,
	toAddr sdk.AccAddress,
	amount sdk.Int) error {

	poolTokens := sdk.NewCoins(sdk.NewCoin(pool.PoolToken.Denom, amount))

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
func NewPool(id uint64, coins sdk.Coins, ammId uint32, swapFeeMulti string, penaltyMulti string, lockDuration int64) types.Pool {

	var pool = types.Pool{
		Id: id,
		Coins: coins,
		AMM_Id:          ammId,
		SwapFeeMulti:    swapFeeMulti,
		MinLockDuration: lockDuration,
		PenaltyMulti:    penaltyMulti,
		PoolToken: sdk.NewCoin(GetPoolTokenDenom(id), sdk.NewIntFromUint64(100)),
	}

	return pool
}

// SetPool set a specific lPool in the store from its index
func (k Keeper) SetPool(ctx sdk.Context, pool types.Pool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))
	b := k.cdc.MustMarshal(&pool)
	store.Set(types.PoolKey(pool.Id), b)
}

// GetPool returns a Pool from its index
func (k Keeper) GetPool(ctx sdk.Context, id uint64) (val types.Pool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))

	b := store.Get(types.PoolKey(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePool removes a pool from the store
func (k Keeper) RemovePool(ctx sdk.Context,	id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))
	store.Delete(types.PoolKey(id))
}

// GetAllPool returns all pool
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

func GetPoolTokenDenom(id uint64) string {
	return fmt.Sprintf("amm/pool/%d", id)
}

// Set pool token metadata to the bank module
func (k Keeper) registerPoolToken(ctx sdk.Context, denom string) {
	_, found := k.bankKeeper.GetDenomMetaData(ctx, denom)
	if found {
		return;
	}

	metaData := banktypes.Metadata{
		Description: "Jackal liquidity pool token",
		DenomUnits:  []*banktypes.DenomUnit{
				&banktypes.DenomUnit{
					Denom: "u" + denom,
					Exponent: 0,
				},
				&banktypes.DenomUnit{
					Denom: denom,
					Exponent: 6,
				},
			},
		Base: denom,
	}

	k.bankKeeper.SetDenomMetaData(ctx, metaData)
}
