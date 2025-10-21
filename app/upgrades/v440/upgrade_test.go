package v440_test

import (
	"github.com/jackalLabs/canine-chain/v5/app/upgrades"
	v440 "github.com/jackalLabs/canine-chain/v5/app/upgrades/v440"
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
)

func (suite *UpgradeTestKeeper) TestUpgrade() {
	suite.SetupSuite()
	setupMsgServer(suite)

	m := []byte("test_file")
	o := "jkl123"
	var s int64 = 1
	suite.storageKeeper.SetFile(suite.ctx, types.UnifiedFile{
		Merkle:        m,
		Owner:         o,
		Start:         s,
		Expires:       10000,
		FileSize:      100,
		ProofInterval: 3600,
		ProofType:     0,
		Proofs:        make([]string, 0),
		MaxProofs:     3,
		Note:          "{}",
	})

	v440.BumpInterval(suite.ctx, suite.storageKeeper)

	f, found := suite.storageKeeper.GetFile(suite.ctx, m, o, s)
	suite.Require().True(found)
	suite.Require().Equal(int64(7200), f.ProofInterval)

	err := upgrades.RecoverFiles(suite.ctx, suite.storageKeeper, v440.UpgradeData, 20000, "v4.4.0-test")
	suite.Require().NoError(err)

	i := 0
	suite.storageKeeper.IterateFilesByMerkle(suite.ctx, false, func(_ []byte, val []byte) bool {
		i++
		var f types.UnifiedFile
		err := f.Unmarshal(val)
		suite.Require().NoError(err)

		suite.T().Logf("%x, %s (%d -> %d) | %d || %s", f.Merkle, f.Owner, f.Start, f.Expires, f.FileSize, f.Note)

		return false
	})

	suite.Require().Equal(25001, i)
}
