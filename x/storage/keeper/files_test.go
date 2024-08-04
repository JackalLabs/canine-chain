package keeper_test

import (
	"encoding/base64"

	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

// testing files.go file
func (suite *KeeperTestSuite) TestFiles() {
	suite.SetupSuite()

	merkle, err := base64.StdEncoding.DecodeString("d9RSWckxX0kFeMt7Ip0GSbhM+eJApeUgicZyL9qBoNiGiMtvibI8XjqXsyTdJC8cVftC8z1BJRrxAtRKe8GVEg==")
	suite.Require().NoError(err)

	providers := suite.storageKeeper.GetActiveProviders(suite.ctx, "")

	owner := "jkl10k05lmc88q5ft3lm00q30qkd9x6654h3lejnct"

	var start int64 = 10

	file := types.UnifiedFile{
		Merkle:        merkle,
		Owner:         owner,
		Start:         start,
		Expires:       0,
		FileSize:      2071,
		ProofInterval: 1800,
		ProofType:     0,
		Proofs:        []string(nil),
		MaxProofs:     3,
		Note:          "test file",
	}
	suite.storageKeeper.SetFile(suite.ctx, file)

	ips := make([]string, 0)

	for i, provider := range providers { // adding all provers
		if i >= int(file.MaxProofs) {
			break
		}
		file.AddProver(suite.ctx, suite.storageKeeper, provider.Address)

		prv, found := suite.storageKeeper.GetProviders(suite.ctx, provider.Address)
		if !found {
			continue
		}

		ips = append(ips, prv.Ip)
	}

	suite.Require().Equal(0, len(ips))

	f, found := suite.storageKeeper.GetFile(suite.ctx, merkle, owner, start)

	suite.Require().True(found)
	suite.Require().Equal(file, f)

	files := suite.storageKeeper.GetAllFilesWithMerkle(suite.ctx, merkle)
	suite.Require().Equal(1, len(files))
}
