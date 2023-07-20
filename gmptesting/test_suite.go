package gmp_testing

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	app "github.com/jackalLabs/canine-chain/v3/app"
	"github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/stretchr/testify/suite"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"
)

// Could this be moved to the app directory for general purpose app testing?

// typically called 'KeeperTestHelper' but this middleware is not stateful so don't think a keeper is needed for now.
type GMPTestHelper struct {
	suite.Suite

	App         *app.JackalApp
	Ctx         sdk.Context
	QueryHelper *baseapp.QueryServiceTestHelper
	TestAccs    []string // we can change to sdk.AccAddress
}

// Setup sets up basic environment for suite (App, Ctx, and test accounts)
func (s *GMPTestHelper) Setup() {
	var t *testing.T
	s.App = app.SetupWithEmptyStore(t) // Might need to replace with 'SetupWithGenesisValSet'
	s.Ctx = s.App.BaseApp.NewContext(false, tmtypes.Header{Height: 1, ChainID: "jackal-1", Time: time.Now().UTC()})
	s.QueryHelper = &baseapp.QueryServiceTestHelper{
		GRPCQueryRouter: s.App.GRPCQueryRouter(),
		Ctx:             s.Ctx,
	}

	// s.SetEpochStartTime() // not sure our use case is concerned with time at this moment

	// to do: handle the error
	s.TestAccs, _ = testutil.CreateTestAddresses("jkl", 3) // could use 'createRandomAccounts' from /app/test_helpers.go
}
