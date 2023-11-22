package v4_test

import (
	"encoding/json"
	"fmt"

	types2 "github.com/jackalLabs/canine-chain/v3/x/storage/types"

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

func (suite *UpgradeTestKeeper) TestStorageUpgrade() {
	suite.SetupSuite()
	setupMsgServer(suite)

	ad := types2.LegacyActiveDeals{
		Cid:           "cid",
		Signee:        "signee",
		Provider:      "provider",
		Startblock:    "0",
		Endblock:      "0",
		Filesize:      "1024",
		Proofverified: "false",
		Proofsmissed:  "0",
		Blocktoprove:  "0",
		Creator:       "creator",
		Merkle:        "941cb8791cb5441674b06de1a931cd101da54457c41e87e9a8ce56e1d39c96bc",
		Fid:           "fid",
	}
	suite.storageKeeper.SetLegacyActiveDeals(suite.ctx, ad)

	v4.UpdateFiles(suite.ctx, *suite.storageKeeper)

	f := suite.storageKeeper.GetAllFileByMerkle(suite.ctx)

	suite.Require().Equal(1, len(f))
}
