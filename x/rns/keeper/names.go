package keeper

import (
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

func (k Keeper) SetPrimaryName(ctx sdk.Context, owner, name, tld string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrimaryNameKeyPrefix))
	store.Set(types.PrimaryNameKey(
		owner,
	), []byte(fmt.Sprintf("%s.%s", name, tld)))
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventSetPrimaryName,
			sdk.NewAttribute(types.AttributeName, fmt.Sprintf("%s.%s", name, tld)),
			sdk.NewAttribute(types.AttributeValue, owner),
		),
	)
}

func (k Keeper) GetPrimaryName(
	ctx sdk.Context,
	owner string,
) (val types.Names, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrimaryNameKeyPrefix))

	b := store.Get(types.PrimaryNameKey(
		owner,
	))
	if b == nil {
		return val, false
	}

	n := string(b)

	nameString := strings.Split(n, ".")

	return k.GetNames(ctx, nameString[0], nameString[1])
}

// SetNames set a specific names in the store from its index
func (k Keeper) SetNames(ctx sdk.Context, names types.Names) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NamesKeyPrefix))
	b := k.cdc.MustMarshal(&names)
	store.Set(types.NamesKey(
		names.Name,
		names.Tld,
	), b)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventSetName,
			sdk.NewAttribute(types.AttributeName, fmt.Sprintf("%s.%s", names.Name, names.Tld)),
			sdk.NewAttribute(types.AttributeValue, names.Value),
			sdk.NewAttribute(types.AttributeExpires, fmt.Sprintf("%d", names.Expires)),
		),
	)
}

// GetNames returns a names from its index
func (k Keeper) GetNames(
	ctx sdk.Context,
	name string,
	tld string,
) (val types.Names, found bool) {
	name = strings.ToLower(name)
	tld = strings.ToLower(tld)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NamesKeyPrefix))

	b := store.Get(types.NamesKey(
		name,
		tld,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNames removes a names from the store
func (k Keeper) RemoveNames(
	ctx sdk.Context,
	name string,
	tld string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NamesKeyPrefix))
	store.Delete(types.NamesKey(
		name,
		tld,
	))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRemoveName,
			sdk.NewAttribute(types.AttributeName, fmt.Sprintf("%s.%s", name, tld)),
		),
	)
}

// GetAllNames returns all names
func (k Keeper) GetAllNames(ctx sdk.Context) (list []types.Names) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NamesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Names
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CheckExistence quickly checks if there are any domains registered
func (k Keeper) CheckExistence(ctx sdk.Context) bool {
	// initializing the iterator
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NamesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	// looping to see if at least 1 element exists
	i := 0
	for ; iterator.Valid(); iterator.Next() {
		if i > 0 {
			break
		}
		i++
	}
	exist := false
	if i > 0 {
		exist = true
	}
	return exist
}

func (k Keeper) Resolve(ctx sdk.Context, name string) (sdk.AccAddress, error) {
	if len(name) == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "name cannot be empty")
	}

	adr, err := sdk.AccAddressFromBech32(name) // the name passed was actually already bech32
	if err == nil {
		return adr, nil
	}

	n, tld, err := GetNameAndTLD(name)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot parse the name and tld from given rns")
	}

	rnsName, found := k.GetNames(ctx, n, tld)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "cannot find name in key store")
	}

	val := rnsName.Value

	address, err := sdk.AccAddressFromBech32(val)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot parse an address from rns entry")
	}

	return address, nil
}
