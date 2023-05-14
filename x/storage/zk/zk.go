package zk

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/hash"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
)

func HashData(data []byte, ccs constraint.ConstraintSystem) (*WrappedProof, []byte, error) {
	element, err := fr.Hash(data, []byte("storage:"), 1)
	if err != nil {
		return nil, nil, err
	}
	bz := element[0].Bytes() // make addressable memory on heap

	h := hash.MIMC_BN254.New()
	_, err = h.Write(bz[:])
	if err != nil {
		return nil, nil, err
	}
	hash := h.Sum(nil)
	// create proof
	assignment := Circuit{Secret: bz[:], Hash: hash}
	wit, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		return nil, hash, err
	}

	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		return nil, hash, err
	}

	// *Prover
	proof, err := groth16.Prove(ccs, pk, wit)
	if err != nil {
		return nil, hash, err
	}

	publicAssignment := Circuit{Hash: hash}
	publicWitness, err := frontend.NewWitness(&publicAssignment, ecc.BN254.ScalarField())
	if err != nil {
		return nil, hash, err
	}
	// *Verifier
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		return nil, hash, err
	}

	return &WrappedProof{
		Proof:        proof,
		VerifyingKey: vk,
	}, hash, nil
}

func VerifyHash(wp *WrappedProof, hash []byte) bool {
	publicAssignment := Circuit{Hash: hash}
	fmt.Println(hash)

	publicWitness, err := frontend.NewWitness(&publicAssignment, ecc.BN254.ScalarField())
	if err != nil {
		return false
	}

	// *Verifier
	err = groth16.Verify(wp.Proof, wp.VerifyingKey, publicWitness)
	if err != nil {
		return false
	}

	return true
}
