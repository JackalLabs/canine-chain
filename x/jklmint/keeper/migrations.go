package keeper

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/jklmint/exported"
	v210 "github.com/jackalLabs/canine-chain/v4/x/jklmint/legacy/v210"
	v3 "github.com/jackalLabs/canine-chain/v4/x/jklmint/legacy/v3"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	k              Keeper
	legacySubspace exported.Subspace
}

// NewMigrator returns a new Migrator
func NewMigrator(keeper Keeper, legacy exported.Subspace) Migrator {
	return Migrator{
		k:              keeper,
		legacySubspace: legacy,
	}
}

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v210.MigrateStore(ctx, &m.k.paramSpace)
}

// Migrate3to4 migrates from version 3 to 4.
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	return v3.MigrateStore(ctx, m.legacySubspace, &m.k.paramSpace)
}
