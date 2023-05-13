package zk

import (
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/frontend"
)

type Circuit struct {
	Secret frontend.Variable
	Hash   frontend.Variable `gnark:",public"`
}

type WrappedProof struct {
	Proof         groth16.Proof
	WitnessPublic witness.Witness
	VerifyingKey  groth16.VerifyingKey
}
