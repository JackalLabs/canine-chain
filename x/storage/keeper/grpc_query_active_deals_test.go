package keeper_test

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	jklapp "github.com/jackalLabs/canine-chain/v3/app"
	"github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/tendermint/tendermint/libs/log"
)

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

func (suite *KeeperTestSuite) SetupTest() {
	app := setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, app.StorageKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.ctx = ctx

	suite.queryClient = queryClient
}

func (suite *KeeperTestSuite) TestAllFiles() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	testAccount := testAddresses[0]
	depoAccount := testAddresses[1]

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(100000000000))) // Send some coins to their account
	testAcc, _ := sdk.AccAddressFromBech32(testAccount)
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, testAcc, coins)
	suite.Require().NoError(err)

	suite.storageKeeper.SetParams(suite.ctx, types.Params{
		DepositAccount:         depoAccount,
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     8,
		CollateralPrice:        2,
		CheckWindow:            10,
	})

	suite.storageKeeper.SetFile(suite.ctx, types.UnifiedFile{
		Merkle:        []byte("merkle"),
		Owner:         testAccount,
		Start:         0,
		Expires:       0,
		FileSize:      1024,
		ProofInterval: 400,
		ProofType:     0,
		Proofs:        make([]string, 0),
		MaxProofs:     3,
		Note:          "test",
	})

	res, err := suite.queryClient.AllFiles(context.Background(), &types.QueryAllFiles{
		Pagination: &query.PageRequest{
			Offset:  0,
			Reverse: false,
		},
	})
	suite.Require().NoError(err)

	suite.Require().Equal(1, len(res.Files))

	suite.reset()
}
