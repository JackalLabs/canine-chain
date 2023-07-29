package gmp_testing

import (
	"crypto/sha256"
	"fmt"
	"os"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/suite"
)

func (chain *TestChain) StoreContractCode(suite *suite.Suite, path string) {
	jackalApp := chain.GetJackalApp()

	govKeeper := jackalApp.GetGovKeeper()
	wasmCode, err := os.ReadFile(path)
	suite.Require().NoError(err)
	fmt.Println(govKeeper)
	fmt.Println(wasmCode)

	addr := jackalApp.AccountKeeper.GetModuleAddress(govtypes.ModuleName)
	src := wasmtypes.StoreCodeProposalFixture(func(p *wasmtypes.StoreCodeProposal) {
		p.RunAs = addr.String()
		p.WASMByteCode = wasmCode
		checksum := sha256.Sum256(wasmCode)
		p.CodeHash = checksum[:]
	})
	fmt.Println(src)

	// when stored
	storedProposal, err := govKeeper.SubmitProposal(chain.GetContext(), src)
	suite.Require().NoError(err)
	fmt.Println(storedProposal)

	// and proposal execute
	handler := govKeeper.Router().GetRoute(storedProposal.ProposalRoute())
	err = handler(chain.GetContext(), storedProposal.GetContent())
	suite.Require().NoError(err)
}
