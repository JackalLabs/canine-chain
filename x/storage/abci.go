package storage

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	//err := k.HandleRewardBlock(ctx)
	//if err != nil {
	//	ctx.Logger().Error(err.Error())
	//}

	k.KillOldContracts(ctx)

	var week int64 = (7 * 24 * 60 * 60) / 6

	if ctx.BlockHeight()%week == 0 { // clear out files once a week
		k.ClearDeadFiles(ctx)
	}
}
