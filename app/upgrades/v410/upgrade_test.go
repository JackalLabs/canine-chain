package v410_test

import (
	v410 "github.com/jackalLabs/canine-chain/v4/app/upgrades/v410"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (suite *UpgradeTestKeeper) TestUpgrade() {
	suite.SetupSuite()
	setupMsgServer(suite)

	err := v410.RecoverFiles(suite.ctx, suite.storageKeeper, 420)
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

	suite.Require().Equal(3, i)
}
