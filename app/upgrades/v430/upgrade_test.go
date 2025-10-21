package v430_test

import (
	"github.com/jackalLabs/canine-chain/v5/app/upgrades"
	v410 "github.com/jackalLabs/canine-chain/v5/app/upgrades/v410"
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
)

func (suite *UpgradeTestKeeper) TestUpgrade() {
	suite.SetupSuite()
	setupMsgServer(suite)

	err := upgrades.RecoverFiles(suite.ctx, suite.storageKeeper, v410.UpgradeData, 20000, "v4.1.0-test")
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

	suite.Require().Equal(25000, i)
}
