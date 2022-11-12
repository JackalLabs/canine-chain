package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/jackalLabs/canine-chain/x/storage/client/cli"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"
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
	err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
	if err != nil {
		fmt.Println(err)
	}
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

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
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
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) { // Every x blocks we check for proven deals

	allDeals := am.keeper.GetAllActiveDeals(ctx)

	height := ctx.BlockHeight()

	const fchunks int64 = 1024

	var dayBlocks int64 = 10 * 5 // 10 blocks is about 1 minute

	ctx.Logger().Debug("blockdiff : %d\n", height%dayBlocks)
	if height%dayBlocks == 0 {
		ctx.Logger().Debug("%s\n", "checking blocks")

		var networkSize int64
		for i := 0; i < len(allDeals); i++ {
			deal := allDeals[i]
			ss, ok := sdk.NewIntFromString(deal.Filesize)
			if !ok {
				continue
			}
			networkSize += ss.Int64()
		}

		address := am.accountKeeper.GetModuleAddress(types.ModuleName)
		balance := am.bankKeeper.GetBalance(ctx, address, "ujkl")

		for i := 0; i < len(allDeals); i++ {
			deal := allDeals[i]

			toprove, ok := sdk.NewIntFromString(deal.Blocktoprove)
			if !ok {
				ctx.Logger().Debug("Int Parse Failed!\n")
				continue
			}

			iprove := toprove.Int64()

			totalSize, ok := sdk.NewIntFromString(deal.Filesize)
			if !ok {
				ctx.Logger().Debug("Int Parse Failed!\n")
				continue
			}

			byteHash := ctx.HeaderHash().Bytes()[0] + ctx.HeaderHash().Bytes()[1] + ctx.HeaderHash().Bytes()[2]

			d := totalSize.Int64() / fchunks

			if d > 0 {
				iprove = (iprove + ctx.BlockHeight()*int64(byteHash)) % d
			}

			deal.Blocktoprove = fmt.Sprintf("%d", iprove)

			verified, errb := strconv.ParseBool(deal.Proofverified)

			if errb != nil {
				ctx.Logger().Debug("ERR %v\n", errb)
				continue
			}

			if !verified {
				ctx.Logger().Debug("%s\n", "Not verified!")
				intt, ok := sdk.NewIntFromString(deal.Proofsmissed)
				if !ok {
					ctx.Logger().Debug("Int Parse Failed!\n")
					continue
				}

				sb, ok := sdk.NewIntFromString(deal.Startblock)
				if !ok {
					ctx.Logger().Debug("Int Parse Failed!\n")
					continue
				}

				if sb.Int64() >= height-dayBlocks {
					continue
				}

				misses := intt.Int64() + 1
				const missesToBurn int64 = 3

				if misses > missesToBurn {
					provider, ok := am.keeper.GetProviders(ctx, deal.Provider)
					if !ok {
						continue
					}

					curburn, ok := sdk.NewIntFromString(provider.BurnedContracts)
					if !ok {
						continue
					}
					provider.BurnedContracts = fmt.Sprintf("%d", curburn.Int64()+1)
					am.keeper.SetProviders(ctx, provider)

					// Creating new stray file from the burned active deal
					strayDeal := types.Strays{
						Cid:      deal.Cid,
						Fid:      deal.Fid,
						Signee:   deal.Signee,
						Filesize: deal.Filesize,
						Merkle:   deal.Merkle,
					}
					am.keeper.SetStrays(ctx, strayDeal)
					am.keeper.RemoveActiveDeals(ctx, deal.Cid)
					continue
				}

				deal.Proofsmissed = fmt.Sprintf("%d", misses)
				am.keeper.SetActiveDeals(ctx, deal)
				continue
			}

			sizeint, ok := sdk.NewIntFromString(deal.Filesize)
			if !ok {
				ctx.Logger().Error("Cannot parse filesize as int")
				continue
			}

			ctx.Logger().Debug(fmt.Sprintf("File size: %s\n", deal.Filesize))
			ctx.Logger().Debug(fmt.Sprintf("Total size: %d\n", networkSize))

			sid := sdk.NewDec(sizeint.Int64())
			ts := sdk.NewDec(networkSize)

			siv, err := sid.Float64()
			if err != nil {
				ctx.Logger().Error(err.Error())
				continue
			}
			div, err := ts.Float64()
			if err != nil {
				ctx.Logger().Error(err.Error())
				continue
			}

			if div == 0 {
				ctx.Logger().Error(sdkerrors.Wrap(types.ErrDivideByZero, "dividing by zero").Error())
				continue
			}

			ratio := siv / div

			ctx.Logger().Debug("Ratio: %f\n", ratio)

			ff, err := sdk.NewDec(balance.Amount.Int64()).Float64()
			if err != nil {
				ctx.Logger().Error(err.Error())
				continue
			}
			coinfloat := ratio * ff
			ctx.Logger().Debug("Coins: %f * %f = %f\n", ratio, ff, coinfloat)

			dd, err := sdk.NewDecFromStr(fmt.Sprintf("%f", coinfloat))
			if err != nil {
				ctx.Logger().Error(err.Error())
				continue
			}

			ctx.Logger().Debug("%f\n", dd)
			coin := sdk.NewInt64Coin("ujkl", dd.TruncateInt64())
			coins := sdk.NewCoins(coin)

			provider, err := sdk.AccAddressFromBech32(deal.Provider)
			if err != nil {
				ctx.Logger().Error(err.Error())
				continue
			}
			ctx.Logger().Debug("Sending coins to %s\n", provider.String())
			errorr := am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, provider, coins)
			if errorr != nil {
				ctx.Logger().Debug("ERR: %v\n", errorr)
				ctx.Logger().Error(errorr.Error())
				continue
			}

			ctx.Logger().Debug("%s\n", deal.Cid)

			deal.Proofverified = "false"
			am.keeper.SetActiveDeals(ctx, deal)

		}
		balance = am.bankKeeper.GetBalance(ctx, am.accountKeeper.GetModuleAddress(types.ModuleName), "ujkl")

		err := am.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(balance))
		if err != nil {
			ctx.Logger().Error("ERR: %v\n", err)
			return
		}
	}
}

// EndBlock executes all ABCI EndBlock logic respective to the capability module. It
// returns no validator updates.
func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
