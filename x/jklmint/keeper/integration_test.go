package keeper_test

import (
	"encoding/json"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v5/x/jklmint/types"
	dbm "github.com/tendermint/tm-db"

	jklapp "github.com/jackalLabs/canine-chain/v5/app"
)

// returns context and an app with updated mint keeper
//
//nolint:unused
func createTestApp(isCheckTx bool) (*jklapp.JackalApp, sdk.Context) {
	app := setup(isCheckTx)

	ctx := app.NewContext(isCheckTx, tmproto.Header{})

	app.MintKeeper.SetParams(ctx, types.DefaultParams())

	return app, ctx
}

func setup(isCheckTx bool) *jklapp.JackalApp {
	app, genesisState := genApp(!isCheckTx, 5)
	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

func genApp(withGenesis bool, invCheckPeriod uint) (*jklapp.JackalApp, jklapp.GenesisState) {
	db := dbm.NewMemDB()
	encCdc := jklapp.MakeEncodingConfig()
	app := jklapp.NewJackalApp(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		simapp.DefaultNodeHome,
		invCheckPeriod,
		encCdc,
		jklapp.GetEnabledProposals(),
		simapp.EmptyAppOptions{},
		jklapp.GetWasmOpts(simapp.EmptyAppOptions{}),
	)

	if withGenesis {
		return app, jklapp.NewDefaultGenesisState()
	}

	return app, jklapp.GenesisState{}
}
