package zk

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/hash/mimc"
)

func GetCircuit() (constraint.ConstraintSystem, error) {
	// *Verifier
	var mimcCircuit Circuit
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &mimcCircuit)
	if err != nil {
		return nil, err
	}

	return ccs, err
}

func (circuit *Circuit) Define(api frontend.API) error {
	mimc, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}

	mimc.Write(circuit.Secret)
	api.AssertIsEqual(mimc.Sum(), circuit.Hash)

	return nil
}
