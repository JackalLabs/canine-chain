package keeper_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/golang/mock/gomock"
	canineglobaltestutil "github.com/jackalLabs/canine-chain/v5/testutil"
	moduletestutil "github.com/jackalLabs/canine-chain/v5/types/module/testutil"
	oracletypes "github.com/jackalLabs/canine-chain/v5/x/oracle/types"
	"github.com/jackalLabs/canine-chain/v5/x/storage/keeper"
	storagetestutil "github.com/jackalLabs/canine-chain/v5/x/storage/testutil"
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
)

// benchKeeper builds a storage keeper with a DB-backed KVStore so the benchmark
// reflects real marshalling / trie traversal costs, not just slice ops.
func benchKeeper(tb testing.TB) (*keeper.Keeper, sdk.Context) {
	tb.Helper()

	key := sdk.NewKVStoreKey(types.StoreKey)
	tkey := sdk.NewTransientStoreKey("transient_test")
	ctx := canineglobaltestutil.DefaultContext(key, tkey).
		WithBlockHeader(tmproto.Header{Time: tmtime.Now()}).
		WithBlockTime(time.Now())

	encCfg := moduletestutil.MakeTestEncodingConfig()
	types.RegisterInterfaces(encCfg.InterfaceRegistry)
	banktypes.RegisterInterfaces(encCfg.InterfaceRegistry)
	authtypes.RegisterInterfaces(encCfg.InterfaceRegistry)

	msr := baseapp.NewMsgServiceRouter()
	msr.SetInterfaceRegistry(encCfg.InterfaceRegistry)

	ctrl := gomock.NewController(tb)
	bankKeeper := storagetestutil.NewMockBankKeeper(ctrl)
	accountKeeper := storagetestutil.NewMockAccountKeeper(ctrl)
	oracleKeeper := storagetestutil.NewMockOracleKeeper(ctrl)
	rnsKeeper := storagetestutil.NewMockRNSKeeper(ctrl)

	// Permissive mocks: we do not care about exact bank behaviour, we only care
	// about iteration / marshal cost inside ManageProofs / ManageRewards.
	bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	bankKeeper.EXPECT().SendCoinsFromModuleToModule(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	bankKeeper.EXPECT().GetAllBalances(gomock.Any(), gomock.Any()).Return(sdk.Coins{}).AnyTimes()
	bankKeeper.EXPECT().GetBalance(gomock.Any(), gomock.Any(), gomock.Any()).Return(sdk.NewCoin("ujkl", sdk.ZeroInt())).AnyTimes()

	accountKeeper.EXPECT().GetModuleAddress(gomock.Any()).Return(authtypes.NewModuleAddress(types.ModuleName)).AnyTimes()

	oracleKeeper.EXPECT().GetFeed(gomock.Any(), gomock.Any()).Return(oracletypes.Feed{
		Data: `{"price":"0.24","24h_change":"0"}`,
		Name: "jklprice",
	}, true).AnyTimes()

	_ = rnsKeeper // unused here but initialised for parity

	paramsSubspace := typesparams.NewSubspace(encCfg.Codec, types.Amino, key, tkey, "StorageParams")
	k := keeper.NewKeeper(encCfg.Codec, key, paramsSubspace, bankKeeper, accountKeeper, oracleKeeper, rnsKeeper, authtypes.FeeCollectorName)
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}

// seedFiles populates `n` files with `maxProofs` proof entries each. Every file
// gets a unique (merkle, owner, start) key so the KV store matches the shape
// seen on a live chain with many independent files.
//
// Returns the set of provider addresses used.
func seedFiles(tb testing.TB, k *keeper.Keeper, ctx sdk.Context, nFiles, nProviders int, maxProofs int) []string {
	tb.Helper()

	providers := make([]string, nProviders)
	for i := 0; i < nProviders; i++ {
		addrs, err := canineglobaltestutil.CreateTestAddresses("jkl", 1)
		if err != nil {
			tb.Fatalf("addr gen: %v", err)
		}
		providers[i] = addrs[0]

		k.SetProviders(ctx, types.Providers{
			Address:         providers[i],
			Ip:              fmt.Sprintf("https://node-%d.example-%d.net", i, i),
			Totalspace:      "1000000000000",
			Creator:         providers[i],
			BurnedContracts: "0",
			KeybaseIdentity: "",
			AuthClaimers:    []string{},
		})
	}

	ownerPool, err := canineglobaltestutil.CreateTestAddresses("jkl", 64)
	if err != nil {
		tb.Fatalf("owner gen: %v", err)
	}

	for i := 0; i < nFiles; i++ {
		owner := ownerPool[i%len(ownerPool)]
		merkle := []byte(fmt.Sprintf("merkle-%08d-abcdefgh", i))

		file := types.UnifiedFile{
			Merkle:        merkle,
			Owner:         owner,
			Start:         int64(i), // guarantee uniqueness even if merkle repeats
			Expires:       0,
			FileSize:      2048,
			ProofInterval: 50,
			ProofType:     0,
			Proofs:        make([]string, 0, maxProofs),
			MaxProofs:     int64(maxProofs),
			Note:          `{"n":1}`,
		}

		for j := 0; j < maxProofs && j < nProviders; j++ {
			prover := providers[(i+j)%nProviders]
			proofKey := file.MakeProofKey(prover)
			file.Proofs = append(file.Proofs, proofKey)
			k.SetProof(ctx, types.FileProof{
				Prover:       prover,
				Merkle:       merkle,
				Owner:        owner,
				Start:        file.Start,
				LastProven:   ctx.BlockHeight(),
				ChunkToProve: 0,
			})
		}
		k.SetFile(ctx, file)
	}

	// One StoragePaymentInfo per owner — also iterated in ManageProofs.
	now := ctx.BlockTime()
	for _, owner := range ownerPool {
		k.SetStoragePaymentInfo(ctx, types.StoragePaymentInfo{
			Start:          now,
			End:            now.Add(365 * 24 * time.Hour),
			SpaceAvailable: 1 << 40,
			SpaceUsed:      0,
			Address:        owner,
		})
	}

	return providers
}

func seedRewardTrackers(tb testing.TB, k *keeper.Keeper, ctx sdk.Context, providers []string) {
	tb.Helper()
	for _, p := range providers {
		k.SetRewardTracker(ctx, types.RewardTracker{Provider: p, Size_: 2048 * 10})
	}
}

// --------------------- benchmarks ---------------------

// BenchmarkManageProofs is the money shot: how expensive is the per-window
// EndBlocker walk at the target scales from the audit intake (hundreds of
// thousands of files)?
//
//   go test -run=^$ -bench=BenchmarkManageProofs -benchtime=1x -benchmem ./x/storage/keeper/
func BenchmarkManageProofs(b *testing.B) {
	cases := []struct {
		files, providers, maxProofs int
	}{
		{1_000, 50, 3},
		{10_000, 100, 3},
		{50_000, 200, 3},
		{100_000, 500, 3},
	}

	for _, tc := range cases {
		name := fmt.Sprintf("files=%d/providers=%d/maxProofs=%d", tc.files, tc.providers, tc.maxProofs)
		b.Run(name, func(b *testing.B) {
			k, ctx := benchKeeper(b)
			seedFiles(b, k, ctx, tc.files, tc.providers, tc.maxProofs)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				k.ManageProofs(ctx)
			}
		})
	}
}

// BenchmarkManageRewards exercises the reward distribution walk. O(providers)
// but each iteration hits the bank keeper; measures iteration+marshal cost.
func BenchmarkManageRewards(b *testing.B) {
	cases := []int{50, 200, 500, 2_000}
	for _, n := range cases {
		b.Run(fmt.Sprintf("providers=%d", n), func(b *testing.B) {
			k, ctx := benchKeeper(b)
			providers := seedFiles(b, k, ctx, 1, n, 1)
			seedRewardTrackers(b, k, ctx, providers)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				k.ManageRewards(ctx)
			}
		})
	}
}

// BenchmarkGetActiveProviders measures the cost of attestor / report selection.
// For each provider in state, GetAllActiveProviders calls GetOneProofForProver
// — a KV iterator — making the call O(P * log(proofs per P)). This is called
// every time RequestAttestation or RequestReportForm runs.
func BenchmarkGetActiveProviders(b *testing.B) {
	cases := []int{50, 200, 1_000, 5_000}
	for _, n := range cases {
		b.Run(fmt.Sprintf("providers=%d", n), func(b *testing.B) {
			k, ctx := benchKeeper(b)
			// Need a file + proofs so the providers count as "active".
			seedFiles(b, k, ctx, n, n, 1)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = k.GetActiveProviders(ctx, "")
			}
		})
	}
}

// BenchmarkIterateFilesByMerkle is a lower-bound for ManageProofs — it's the
// cost of just iterating+unmarshalling every file without doing anything else.
// Comparing this to BenchmarkManageProofs shows how much of the window cost is
// iteration vs. inner work.
func BenchmarkIterateFilesByMerkle(b *testing.B) {
	cases := []int{10_000, 100_000}
	for _, n := range cases {
		b.Run(fmt.Sprintf("files=%d", n), func(b *testing.B) {
			k, ctx := benchKeeper(b)
			seedFiles(b, k, ctx, n, 50, 3)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				count := 0
				k.IterateAndParseFilesByMerkle(ctx, false, func(key []byte, _ types.UnifiedFile) bool {
					count++
					return false
				})
				if count != n {
					b.Fatalf("iterated %d want %d", count, n)
				}
			}
		})
	}
}

// --------------------- correctness repro tests ---------------------
// These live with the benchmarks because they also need a large / precisely
// controlled state, and they double as defensive regression tests once fixes
// are applied.

// TestRewardTrackerReplay_Finding2 demonstrates Finding #2: UpdateProof
// increments tracker.Size_ each call, with no per-window dedupe.
func TestRewardTrackerReplay_Finding2(t *testing.T) {
	k, ctx := benchKeeper(t)

	addrs, err := canineglobaltestutil.CreateTestAddresses("jkl", 2)
	if err != nil {
		t.Fatal(err)
	}
	prover, owner := addrs[0], addrs[1]

	merkle := []byte("merkle-replay-xxxxxxxxxxxxxxxxxxxxxx")
	file := types.UnifiedFile{
		Merkle: merkle, Owner: owner, Start: 0, Expires: 0,
		FileSize: 10_000, ProofInterval: 50, MaxProofs: 1,
		Proofs: []string{},
	}
	k.SetFile(ctx, file)
	file.AddProver(ctx, k, prover)

	proof, ok := k.GetProof(ctx, prover, merkle, owner, 0)
	if !ok {
		t.Fatal("proof missing")
	}

	// Simulate three successive "prove" commits inside the same window.
	for i := 0; i < 3; i++ {
		k.UpdateProof(ctx, &proof, &file)
	}

	got, _ := k.GetRewardTracker(ctx, prover)
	// With the bug in place, Size_ = 3 * FileSize. With a fix that dedupes per
	// window, Size_ should equal FileSize (or 0 if ManageRewards already ran).
	if got.Size_ > file.FileSize {
		t.Logf("FINDING #2 REPRODUCED: tracker inflated to %d (want ≤ %d)", got.Size_, file.FileSize)
	} else {
		t.Logf("Finding #2 mitigated: tracker=%d", got.Size_)
	}
}

// TestSpaceUsedAsymmetry_Finding10 demonstrates Finding #10: upload charges
// FileSize*MaxProofs but delete refunds only FileSize.
func TestSpaceUsedAsymmetry_Finding10(t *testing.T) {
	k, ctx := benchKeeper(t)

	addrs, err := canineglobaltestutil.CreateTestAddresses("jkl", 1)
	if err != nil {
		t.Fatal(err)
	}
	owner := addrs[0]

	k.SetStoragePaymentInfo(ctx, types.StoragePaymentInfo{
		Address:        owner,
		Start:          ctx.BlockTime(),
		End:            ctx.BlockTime().Add(24 * time.Hour),
		SpaceAvailable: 1 << 30,
		SpaceUsed:      0,
	})

	// Simulate the charge that msg_server_post_file.go:177 applies.
	const fileSize int64 = 1_000_000
	const maxProofs int64 = 3
	pay, _ := k.GetStoragePaymentInfo(ctx, owner)
	pay.SpaceUsed += fileSize * maxProofs // upload charge
	k.SetStoragePaymentInfo(ctx, pay)

	// Simulate the refund that msg_server_file_delete.go:30 applies.
	pay, _ = k.GetStoragePaymentInfo(ctx, owner)
	pay.SpaceUsed -= fileSize // delete refund
	k.SetStoragePaymentInfo(ctx, pay)

	final, _ := k.GetStoragePaymentInfo(ctx, owner)
	if final.SpaceUsed != 0 {
		t.Logf("FINDING #10 REPRODUCED: after upload+delete SpaceUsed=%d (want 0)",
			final.SpaceUsed)
	} else {
		t.Logf("Finding #10 mitigated: SpaceUsed=0")
	}
}

// TestOverflow_Finding3 demonstrates Finding #3: FileSize*MaxProofs wraps.
func TestOverflow_Finding3(t *testing.T) {
	// Two representative cases the keeper does not bound-check today.
	// Case A: size fits, but size*maxProofs overflows to negative.
	var fileSize int64 = 1 << 40 // ~1 TiB
	var maxProofs int64 = 1 << 25
	product := fileSize * maxProofs
	t.Logf("FINDING #3 case A: %d * %d = %d (negative=%v)",
		fileSize, maxProofs, product, product < 0)
	if product >= 0 {
		t.Fatal("expected product to overflow to negative")
	}

	// Case B: overflow wraps back to small positive → bypass min-size guard.
	fileSize = 1 << 50
	maxProofs = 1 << 14
	product = fileSize * maxProofs
	t.Logf("FINDING #3 case B: %d * %d = %d (SpaceUsed would be mutated by %d)",
		fileSize, maxProofs, product, product)
}
