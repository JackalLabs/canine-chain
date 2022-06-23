package jklmining

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/jackal-dao/canine/x/jklmining/client/cli"
	"github.com/jackal-dao/canine/x/jklmining/keeper"
	"github.com/jackal-dao/canine/x/jklmining/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface for the capability module.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the capability module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

// RegisterInterfaces registers the module's interface types
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns the capability module's default genesis state.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis performs genesis state validation for the capability module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterRESTRoutes registers the capability module's REST service handlers.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the capability module's root tx command.
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns the capability module's root query command.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey)
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface for the capability module.
type AppModule struct {
	AppModuleBasic

	keeper            keeper.Keeper
	accountKeeper     types.AccountKeeper
	bankKeeper        types.BankKeeper
	jklAccountsKeeper types.JackalAccountsKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	jklAccountsKeeper types.JackalAccountsKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic:    NewAppModuleBasic(cdc),
		keeper:            keeper,
		accountKeeper:     accountKeeper,
		bankKeeper:        bankKeeper,
		jklAccountsKeeper: jklAccountsKeeper,
	}
}

// Name returns the capability module's name.
func (am AppModule) Name() string {
	return am.AppModuleBasic.Name()
}

// Route returns the capability module's message routing key.
func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, NewHandler(am.keeper))
}

// QuerierRoute returns the capability module's query routing key.
func (AppModule) QuerierRoute() string { return types.QuerierRoute }

// LegacyQuerierHandler returns the capability module's Querier.
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers a GRPC query service to respond to the
// module-specific GRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the capability module's invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the capability module's genesis initialization It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, genState)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the capability module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion implements ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 2 }

// BeginBlock executes all ABCI BeginBlock logic respective to the capability module.
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	//STARTBLOCK

	bheight := ctx.BlockHeight()

	index := am.keeper.GetMinedStarting(ctx)
	ctx.Logger().Info(fmt.Sprintf("Loading index from storage makes it: %d", index))

	var f uint64 = 0
	j, found := am.keeper.GetMined(ctx, index)
	if found {

		k, _ := sdk.NewIntFromString(j.Pcount)
		for k.Int64()+20 < bheight {
			f += 1
			j, found = am.keeper.GetMined(ctx, index+f)
			if !found {
				ctx.Logger().Error(fmt.Sprintf("Mined block not found at index: %d", index+f))
				break
			}
			k, _ = sdk.NewIntFromString(j.Pcount)
		}
		ctx.Logger().Info(fmt.Sprintf("Moving the index up from %d to %d", index, index+f))
		am.keeper.PushMinedStarting(ctx, f)

		toburn := am.bankKeeper.GetBalance(ctx, am.accountKeeper.GetModuleAddress(am.Name()), "ujkl")

		dex := index + f
		total := am.keeper.GetMinedCount(ctx) - dex

		if total > 0 {
			coinValue := toburn.Amount.Uint64() / total

			m, _ := sdk.NewIntFromString(fmt.Sprintf("%d", coinValue))

			var l uint64 = 0
			for l < total {
				block, _ := am.keeper.GetMined(ctx, l+index)
				claim, _ := am.keeper.GetMinerClaims(ctx, block.Hash)
				address, _ := sdk.AccAddressFromBech32(claim.Creator)
				am.bankKeeper.SendCoinsFromModuleToAccount(ctx, am.Name(), address, sdk.NewCoins(sdk.NewCoin("ujkl", m)))

				l++
			}
		}

	} else {
		ctx.Logger().Error(fmt.Sprintf("Mined block not found at index: %d", index))
	}

	toburn := am.bankKeeper.GetBalance(ctx, am.accountKeeper.GetModuleAddress(am.Name()), "ujkl")
	toburns := sdk.NewCoins(toburn)
	err := am.bankKeeper.BurnCoins(ctx, am.Name(), toburns)
	if err != nil {
		ctx.Logger().Error(fmt.Sprintf("%s", err.Error()))
	}

}

// EndBlock executes all ABCI EndBlock logic respective to the capability module. It
// returns no validator updates.
func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
