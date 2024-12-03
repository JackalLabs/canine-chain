package upgrades

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	allTypes "github.com/jackalLabs/canine-chain/v4/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	storageKeeper "github.com/jackalLabs/canine-chain/v4/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

const AccountName = "STORAGE_ESCROW"

func RecoverFiles(ctx sdk.Context, keeper *storageKeeper.Keeper, merkles string, planHeight int64, name string) error {
	account, err := allTypes.GetAccount(AccountName) // creating account to hold the files
	if err != nil {
		return err
	}

	for i, line := range strings.Split(strings.TrimSuffix(merkles, "\n"), "\n") {
		ctx.Logger().Info(line)
		data := strings.Split(line, ",")
		if len(data) < 3 {
			continue
		}
		merkle := data[0]
		sizeString := data[1]
		fid := data[2]
		size, err := strconv.ParseInt(sizeString, 10, 64)
		if err != nil {
			ctx.Logger().Error("cannot decode %s, skipping...", merkle)
			return err
		}

		merkleBytes, err := hex.DecodeString(merkle)
		if err != nil {
			ctx.Logger().Error("cannot decode %s, skipping...", merkle)
			return err
		}
		start := planHeight - int64(i/32) // 32 files per block
		if start < 0 {
			start = 0
		}
		f := types.UnifiedFile{
			Merkle:        merkleBytes,
			Owner:         account.String(),
			Start:         start,
			Expires:       planHeight + ((200 * 365 * 24 * 60 * 60) / 6),
			FileSize:      size,
			ProofInterval: 3600,
			ProofType:     0,
			Proofs:        make([]string, 0),
			MaxProofs:     3,
			Note:          fmt.Sprintf("{\"memo\":\"Recovered during %s\", \"fid\":\"%s\"}", name, fid),
		}

		keeper.SetFile(ctx, f)

	}
	return nil
}
