package types_test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	"github.com/stretchr/testify/require"
	"github.com/wealdtech/go-merkletree"
	"github.com/wealdtech/go-merkletree/sha3"
)

// copied from the provider code, we should probably make this an import for the providers
func GenerateMerkleProof(tree merkletree.MerkleTree, index int, item []byte) (bool, *merkletree.Proof, error) {
	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%d%x", index, item))
	if err != nil {
		panic(err)
	}

	proof, err := tree.GenerateProof(h.Sum(nil), 0)
	if err != nil {
		panic(err)
	}

	valid, err := merkletree.VerifyProofUsing(h.Sum(nil), false, proof, [][]byte{tree.Root()}, sha3.New512())
	if err != nil {
		panic(err)
	}
	return valid, proof, nil
}

func TestFileProof(t *testing.T) {
	i := []byte("Hello world!")

	var hashBuilder strings.Builder
	hashBuilder.WriteString(strconv.FormatInt(0, 10))
	hashBuilder.WriteString(hex.EncodeToString(i))

	hash := sha256.New()
	_, err := io.WriteString(hash, hashBuilder.String())
	require.NoError(t, err)

	data := make([][]byte, 1)

	hashName := hash.Sum(nil)
	data[0] = hashName

	tree, err := merkletree.NewUsing(data, sha3.New512(), false)
	require.NoError(t, err)

	f := &types.UniversalFile{
		Id:            nil,
		Proofs:        make(map[string]*types.FileProof),
		FileSize:      1,
		Start:         0,
		Expires:       0,
		Owner:         "",
		ProofInterval: 0,
		Merkle:        tree.Root(),
	}

	verified, proof, err := GenerateMerkleProof(*tree, 0, i)
	require.NoError(t, err)

	jproof, err := json.Marshal(*proof)
	require.NoError(t, err)

	require.Equal(t, true, verified)

	item := fmt.Sprintf("%x", i)
	proofData := jproof

	var chunk int64

	validProof := f.VerifyProof(proofData, chunk, item)
	require.Equal(t, true, validProof)

	validProof = f.VerifyProof(proofData, chunk+1, item)
	require.Equal(t, false, validProof)
}
