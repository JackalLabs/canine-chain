package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      sdk.StoreKey
		paramStore    paramtypes.Subspace
		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
		oracleKeeper  types.OracleKeeper
		rnsKeeper     types.RnsKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	ps paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	oracleKeeper types.OracleKeeper,
	rnsKeeper types.RnsKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		paramStore:    ps,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		oracleKeeper:  oracleKeeper,
		rnsKeeper:     rnsKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
