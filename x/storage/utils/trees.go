package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	"github.com/wealdtech/go-merkletree/v2"
	"github.com/wealdtech/go-merkletree/v2/sha3"
)

func BuildTree(buf io.Reader, chunkSize int64) ([]byte, []byte, [][]byte, int, error) {
	tree, chunks, size, err := BuildJustTree(buf, chunkSize)
	if err != nil {
		return nil, nil, nil, 0, err
	}

	r := tree.Root()

	exportedTree, err := json.Marshal(tree)
	if err != nil {
		return nil, nil, nil, 0, err
	}

	return r, exportedTree, chunks, size, nil
}

func BuildJustTree(buf io.Reader, chunkSize int64) (*merkletree.MerkleTree, [][]byte, int, error) {
	size := 0

	data := make([][]byte, 0)
	chunks := make([][]byte, 0)

	index := 0

	for {
		b := make([]byte, chunkSize)
		read, _ := buf.Read(b)

		if read == 0 {
			break
		}

		b = b[:read]

		size += read

		chunks = append(chunks, b)

		hexedData := hex.EncodeToString(b)

		hash := sha256.New()
		hash.Write([]byte(fmt.Sprintf("%d%s", index, hexedData))) // appending the index and the data
		hashName := hash.Sum(nil)

		data = append(data, hashName)

		index++
	}

	tree, err := merkletree.NewUsing(data, sha3.New512(), false)
	if err != nil {
		return nil, nil, 0, err
	}

	return tree, chunks, size, nil
}

// VerifyProof checks whether a proof is valid against a merkle
func VerifyProof(merkle []byte, proofData []byte, chunk int64, item []byte) bool {
	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%d%x", chunk, item))
	if err != nil {
		return false
	}
	hashName := h.Sum(nil)

	var proof merkletree.Proof // unmarshal proof
	err = json.Unmarshal(proofData, &proof)
	if err != nil {
		return false
	}

	verified, err := merkletree.VerifyProofUsing(hashName, false, &proof, [][]byte{merkle}, sha3.New512())
	if err != nil {
		return false
	}

	return verified
}
