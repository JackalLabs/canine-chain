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

var Scalar = ecc.BN254.ScalarField

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
	hashValue := h.Sum(nil)
	// create proof
	assignment := Circuit{Secret: bz[:], Hash: hashValue}
	wit, err := frontend.NewWitness(&assignment, Scalar())
	if err != nil {
		return nil, hashValue, err
	}

	fmt.Printf("building hash: %x\n", hashValue)

	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		return nil, hashValue, err
	}

	// *Prover
	proof, err := groth16.Prove(ccs, pk, wit)
	if err != nil {
		return nil, hashValue, err
	}

	publicWitness, err := wit.Public()
	if err != nil {
		return nil, hashValue, err
	}

	// *Verifier
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		return nil, hashValue, err
	}

	return &WrappedProof{
		Proof:        proof,
		VerifyingKey: vk,
	}, hashValue, nil
}

func VerifyHash(wp *WrappedProof, hash []byte) (bool, error) {
	publicAssignment := Circuit{Hash: hash}
	fmt.Printf("verification hash: %x\n", hash)

	publicWitness, err := frontend.NewWitness(&publicAssignment, Scalar(), frontend.PublicOnly())
	if err != nil {
		return false, fmt.Errorf("%w: Cannot create a new witness", err)
	}

	// *Verifier
	err = groth16.Verify(wp.Proof, wp.VerifyingKey, publicWitness)
	if err != nil {
		return false, fmt.Errorf("%w: Cannot verify proof", err)
	}

	return true, nil
}
