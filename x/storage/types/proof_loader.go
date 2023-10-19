package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type ProofLoader interface {
	SetProof(ctx sdk.Context, Proof FileProof)
	GetProofWithBuiltKey(
		ctx sdk.Context,
		key []byte,
	) (val FileProof, found bool)
	RemoveProofWithBuiltKey(
		ctx sdk.Context,
		key []byte,
	)
}
