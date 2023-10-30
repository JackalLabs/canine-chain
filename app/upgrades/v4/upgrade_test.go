package v4_test

import (
	"encoding/json"
	"fmt"

	v4 "github.com/jackalLabs/canine-chain/v3/app/upgrades/v4"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func (suite *UpgradeTestKeeper) TestUpgrade() {
	suite.SetupSuite()
	setupMsgServer(suite)

	fidMerkleMap := make(map[string][]byte)

	for i := 0; i < 10; i++ {
		fid := fmt.Sprintf("jklf1oooooooo%d", i)

		fidContents := v4.FidContents{Fid: []string{fid}}
		data, err := json.Marshal(fidContents)
		suite.Require().NoError(err)

		address := fmt.Sprintf("this is file %d", i)
		f := types.Files{
			Address:        address,
			Contents:       string(data),
			Owner:          "cosmosjoe",
			ViewingAccess:  "{}",
			EditAccess:     "{}",
			TrackingNumber: fmt.Sprintf("%d", i),
		}
		fidMerkleMap[fid] = []byte(address)

		suite.filetreeKeeper.SetFiles(suite.ctx, f)
	}

	v4.UpdateFileTree(suite.ctx, *suite.filetreeKeeper, fidMerkleMap)

	allFiles := suite.filetreeKeeper.GetAllFiles(suite.ctx)

	for _, file := range allFiles {
		var mct v4.MerkleContents
		b := []byte(file.Contents)
		err := json.Unmarshal(b, &mct)
		suite.Require().NoError(err)

		suite.Require().Equal([]byte(file.Address), mct.Merkles[0])
	}
}
