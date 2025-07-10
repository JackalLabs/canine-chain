package v4_test

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cosmos/btcutil/bech32"

	types2 "github.com/jackalLabs/canine-chain/v4/x/storage/types"

	v4 "github.com/jackalLabs/canine-chain/v4/app/upgrades/v4"
	"github.com/jackalLabs/canine-chain/v4/x/filetree/types"
)

func (suite *UpgradeTestKeeper) TestUpgrade() {
	suite.SetupSuite()
	setupMsgServer(suite)

	fidMerkleMap := make(map[string][]byte)

	for i := range 10 {
		n := rand.Int63()
		s := sha256.New()
		fmt.Fprintf(s, "%d", n)
		d := s.Sum(nil)
		conv, err := bech32.ConvertBits(d, 8, 5, true)
		suite.Require().NoError(err)
		fid, err := bech32.Encode("jklf", conv)
		suite.Require().NoError(err)

		fidContents := v4.FidContents{Fid: []string{fid}}
		data, err := json.Marshal(fidContents)
		suite.Require().NoError(err)

		s = sha256.New()
		s.Write([]byte("cosmosjoe"))
		conv, err = bech32.ConvertBits(s.Sum(nil), 8, 5, true)
		suite.Require().NoError(err)
		owner, err := bech32.Encode("jkl", conv)
		suite.Require().NoError(err)

		address := fmt.Sprintf("this is file %d", i)
		f := types.Files{
			Address:        address,
			Contents:       string(data),
			Owner:          owner,
			ViewingAccess:  "{}",
			EditAccess:     "{}",
			TrackingNumber: fmt.Sprintf("%d", i),
		}
		s = sha256.New()
		s.Write([]byte(address))
		fidMerkleMap[fid] = s.Sum(nil)

		suite.filetreeKeeper.SetFiles(suite.ctx, f)

		fj, err := json.MarshalIndent(f, "", "  ")
		suite.Require().NoError(err)
		fmt.Println(string(fj))
	}

	v4.UpdateFileTree(suite.ctx, suite.filetreeKeeper, fidMerkleMap)

	allFiles := suite.filetreeKeeper.GetAllFiles(suite.ctx)

	for _, file := range allFiles {
		var mct v4.MerkleContents
		b := []byte(file.Contents)
		err := json.Unmarshal(b, &mct)
		suite.Require().NoError(err)

		s := sha256.New()
		s.Write([]byte(file.Address))

		suite.Require().Equal(s.Sum(nil), mct.Merkles[0])

		fj, err := json.MarshalIndent(file, "", "  ")
		suite.Require().NoError(err)
		fmt.Println(string(fj))
	}
}

func (suite *UpgradeTestKeeper) TestStorageUpgrade() {
	suite.SetupSuite()
	setupMsgServer(suite)

	ad := types2.LegacyActiveDeals{
		Cid:          "cid",
		Signee:       "signee",
		Provider:     "provider",
		Startblock:   "0",
		Endblock:     "0",
		Filesize:     "1024",
		LastProof:    suite.ctx.BlockHeight(),
		Proofsmissed: "0",
		Blocktoprove: "0",
		Creator:      "creator",
		Merkle:       "941cb8791cb5441674b06de1a931cd101da54457c41e87e9a8ce56e1d39c96bc",
		Fid:          "fid",
	}
	suite.storageKeeper.SetLegacyActiveDeals(suite.ctx, ad)

	v4.UpdateFiles(suite.ctx, suite.storageKeeper)

	f := suite.storageKeeper.GetAllFileByMerkle(suite.ctx)

	suite.Require().Equal(1, len(f))
}
