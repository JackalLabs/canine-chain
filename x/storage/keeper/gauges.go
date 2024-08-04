package keeper

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (k Keeper) NewGauge(ctx sdk.Context, coins sdk.Coins, end time.Time) types.PaymentGauge {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d--%d--%s", ctx.BlockHeight(), end.UnixMicro(), coins.String())))
	id := h.Sum(nil)

	pg := types.PaymentGauge{
		Id:    id,
		Start: ctx.BlockTime(),
		End:   end,
		Coins: coins,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaymentGaugeKeyPrefix))
	b := k.cdc.MustMarshal(&pg)
	store.Set(types.PaymentGaugeKey(
		id,
	), b)

	return pg
}

// IterateGauges iterates and runs `fn` for every gauge
func (k Keeper) IterateGauges(ctx sdk.Context, fn func(pg types.PaymentGauge)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaymentGaugeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PaymentGauge
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		fn(val)
	}
}

// RemoveGauge removes a PaymentGauge  from the store
func (k Keeper) RemoveGauge(
	ctx sdk.Context,
	id []byte,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaymentGaugeKeyPrefix))
	store.Delete(types.PaymentGaugeKey(
		id,
	))
}

// SetPaymentGauge set a specific payment gauge in the store from its index
func (k Keeper) SetPaymentGauge(ctx sdk.Context, gauge types.PaymentGauge) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaymentGaugeKeyPrefix))
	b := k.cdc.MustMarshal(&gauge)
	store.Set(types.PaymentGaugeKey(
		gauge.Id,
	), b)
}

// GetAllPaymentGauges returns all payment gauges
func (k Keeper) GetAllPaymentGauges(ctx sdk.Context) (list []types.PaymentGauge) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaymentGaugeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PaymentGauge
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
