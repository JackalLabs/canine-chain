package zk

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/hash"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/hash/mimc"
)

func (circuit *Circuit) Define(api frontend.API) error {
	mimc, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}

	mimc.Write(circuit.Secret)
	api.AssertIsEqual(mimc.Sum(), circuit.Hash)

	return nil
}

func HashData(data []byte) (*WrappedProof, error) {
	element, err := fr.Hash(data, []byte("storage:"), 1)
	if err != nil {
		return nil, err
	}
	bz := element[0].Bytes() // make addressable memory on heap

	h := hash.MIMC_BN254.New()
	_, err = h.Write(bz[:])
	if err != nil {
		return nil, err
	}
	hash := h.Sum(nil)
	// debug
	fmt.Printf("Hash(public): %s\n", big.NewInt(0).SetBytes(hash).String())
	fmt.Printf("PreImage(secret): %s\n", big.NewInt(0).SetBytes(bz[:]).String())

	// *Prover
	// create proof
	assignment := Circuit{Secret: bz[:], Hash: hash}
	wit, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		return nil, err
	}
	witnessPublic, err := wit.Public()
	if err != nil {
		return nil, err
	}

	// *Verifier
	var mimcCircuit Circuit
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &mimcCircuit)
	if err != nil {
		panic(err)
	}

	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		panic(err)
	}

	// *Prover
	proof, err := groth16.Prove(ccs, pk, wit)
	if err != nil {
		panic(err)
	}

	return &WrappedProof{
		Proof:         proof,
		WitnessPublic: witnessPublic,
		VerifyingKey:  vk,
	}, nil
}

func VerifyHash(wp *WrappedProof) bool {
	// *Verifier
	err := groth16.Verify(wp.Proof, wp.VerifyingKey, wp.WitnessPublic)
	if err != nil {
		return false
	}

	return true
}
