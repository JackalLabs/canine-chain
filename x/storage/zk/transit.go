package zk

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (w *WrappedProof) Encode() (*types.ProofPackage, error) {
	var proof bytes.Buffer
	_, err := w.Proof.WriteTo(&proof)
	if err != nil {
		return nil, err
	}
	proof64 := base64.StdEncoding.EncodeToString(proof.Bytes())

	var vkey bytes.Buffer
	_, err = w.VerifyingKey.WriteTo(&vkey)
	if err != nil {
		return nil, err
	}
	key64 := base64.StdEncoding.EncodeToString(vkey.Bytes())

	pp := &types.ProofPackage{
		Proof:        proof64,
		VerifyingKey: key64,
	}

	return pp, err
}

func Decode(proofPackge *types.ProofPackage) (*WrappedProof, error) {
	p := groth16.NewProof(ecc.BN254)
	vk := groth16.NewVerifyingKey(ecc.BN254)
	wp := WrappedProof{
		Proof:        p,
		VerifyingKey: vk,
	}

	proof64, err := base64.StdEncoding.DecodeString(proofPackge.Proof)
	if err != nil {
		return nil, fmt.Errorf("%w: cannot decode proof", err)
	}
	proof := bytes.NewBuffer(proof64)
	_, err = wp.Proof.ReadFrom(proof)
	if err != nil {
		return nil, fmt.Errorf("%w: cannot unmarshal proof", err)
	}

	key64, err := base64.StdEncoding.DecodeString(proofPackge.VerifyingKey)
	if err != nil {
		return nil, fmt.Errorf("%w: cannot decode key", err)
	}
	key := bytes.NewBuffer(key64)
	_, err = wp.VerifyingKey.ReadFrom(key)
	if err != nil {
		return nil, fmt.Errorf("%w: cannot unmarshal key", err)
	}

	return &wp, err
}
