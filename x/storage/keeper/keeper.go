package keeper

import (
	"database/sql"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

type (
	Keeper struct {
		cdc              codec.BinaryCodec
		storeKey         sdk.StoreKey
		paramStore       paramtypes.Subspace
		bankKeeper       types.BankKeeper
		accountKeeper    types.AccountKeeper
		oracleKeeper     types.OracleKeeper
		rnsKeeper        types.RnsKeeper
		feeCollectorName string

		filebase *sql.DB
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
	feeCollectorName string,
	filebase *sql.DB,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	err := CreateTablesIfNotExist(filebase)
	if err != nil {
		panic(err)
	}

	return &Keeper{
		cdc:              cdc,
		storeKey:         storeKey,
		paramStore:       ps,
		bankKeeper:       bankKeeper,
		accountKeeper:    accountKeeper,
		oracleKeeper:     oracleKeeper,
		rnsKeeper:        rnsKeeper,
		feeCollectorName: feeCollectorName,
		filebase:         filebase,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) AddCollectedFees(ctx sdk.Context, fees sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, fees)
}
