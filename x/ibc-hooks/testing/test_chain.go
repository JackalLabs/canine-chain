package testing

import (
	"encoding/json"

	ibctesting "github.com/cosmos/ibc-go/v4/testing"
	"github.com/jackalLabs/canine-chain/v3/app"
)

type TestChain struct {
	*ibctesting.TestChain
}

func SetupTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	// Use the public Setup function that doesn't require testing.TB
	jackalApp, _ := app.Setup(false, 0)
	return jackalApp, app.NewDefaultGenesisState()
}

// GetJackalApp returns the current chain's app as a JackalApp
func (chain *TestChain) GetJackalApp() *app.JackalApp {
	v, _ := chain.App.(*app.JackalApp)
	return v
}
