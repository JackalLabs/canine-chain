package keeper

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/exported"
	"github.com/jackalLabs/canine-chain/v4/x/storage/legacy/paramUpgrade"
	v2 "github.com/jackalLabs/canine-chain/v4/x/storage/legacy/v2"
	v4 "github.com/jackalLabs/canine-chain/v4/x/storage/legacy/v4"
	v5 "github.com/jackalLabs/canine-chain/v4/x/storage/legacy/v5"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	k              Keeper
	legacySubspace exported.Subspace
}

// NewMigrator returns a new Migrator
func NewMigrator(keeper Keeper, legacySubspace exported.Subspace) Migrator {
	return Migrator{
		k:              keeper,
		legacySubspace: legacySubspace,
	}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, &m.k.paramStore)
}

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, &m.k.paramStore)
}

// Migrate3to4 migrates from version 3 to 4.
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	return paramupgrade.MigrateStore(ctx, &m.k.paramStore)
}

// Migrate4to5 migrates from version 4 to 5.
func (m Migrator) Migrate4to5(ctx sdk.Context) error {
	return v4.MigrateStore(ctx, m.legacySubspace, &m.k.paramStore)
}

// Migrate5to6 migrates from version 5 to 6.
func (m Migrator) Migrate5to6(ctx sdk.Context) error {
	return v5.MigrateStore(ctx, m.legacySubspace, &m.k.paramStore)
}
