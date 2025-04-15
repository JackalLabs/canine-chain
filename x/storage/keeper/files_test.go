package keeper_test

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/stretchr/testify/require"
	"testing"

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

func benchNewVOld(count int, b *testing.B, new bool) {
	name := "old"
	if new {
		name = "new"
	}

	r := require.New(b)

	storageKeeper, _, _, encCfg, ctx := setupStorageKeeperBench(b)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, storageKeeper)
	queryClient := types.NewQueryClient(queryHelper)
	_ = queryClient

	for i := 0; i < count; i++ {
		s := sha512.New()
		s.Write([]byte(fmt.Sprintf("s=%d", i)))
		file := types.UnifiedFile{
			Merkle:        s.Sum(nil),
			Owner:         "marston",
			Start:         1,
			Expires:       100,
			FileSize:      2000,
			ProofInterval: 720,
			ProofType:     0,
			Proofs:        []string{},
			MaxProofs:     3,
			Note:          `{"test": "yes"}`,
		}

		if new {
			err := storageKeeper.SetFile(ctx, file)
			r.NoError(err)
		} else {
			storageKeeper.SetFileOld(ctx, file)
		}
	}

	b.Run(fmt.Sprintf("%s_list_file_count_%d", name, count), func(b *testing.B) {
		if new {
			files := storageKeeper.GetAllFileByMerkle()
			r.Equal(count, len(files))
		} else {
			files := storageKeeper.GetAllFileByMerkleOld(ctx)
			r.Equal(count, len(files))
		}
	})

	b.Run(fmt.Sprintf("%s_file_size_count_%d", name, count), func(b *testing.B) {
		if new {
			_, err := storageKeeper.GetTotalFileSize()
			r.NoError(err)
		} else {
			var totalSize int64
			storageKeeper.IterateFilesByMerkleOld(ctx, false, func(_ []byte, val []byte) bool {
				var file types.UnifiedFile
				encCfg.Codec.MustUnmarshal(val, &file)

				s := file.FileSize * int64(len(file.Proofs))
				totalSize += s

				return false
			})
		}
	})

}

func BenchmarkFileBaseAgainstKV(b *testing.B) {

	counts := []int{10, 100, 1_000, 10_000, 100_000, 1_000_000}

	for _, v := range counts {

		for i := 0; i < b.N; i++ {
			benchNewVOld(v, b, true)
		}

		for i := 0; i < b.N; i++ {
			benchNewVOld(v, b, false)
		}
	}
}
