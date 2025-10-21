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

	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
	"github.com/stretchr/testify/require"
	"github.com/wealdtech/go-merkletree/v2"
	"github.com/wealdtech/go-merkletree/v2/sha3"
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

	f := &types.UnifiedFile{
		Merkle:        tree.Root(),
		Owner:         "",
		Start:         0,
		Expires:       0,
		FileSize:      1,
		ProofInterval: 10,
		ProofType:     0,
		Proofs:        make([]string, 0),
		MaxProofs:     3,
		Note:          "",
	}

	verified, proof, err := GenerateMerkleProof(*tree, 0, i)
	require.NoError(t, err)

	jproof, err := json.Marshal(*proof)
	require.NoError(t, err)

	require.Equal(t, true, verified)

	proofData := jproof

	var chunk int64

	validProof := f.VerifyProof(proofData, chunk, i)
	require.Equal(t, true, validProof)

	validProof = f.VerifyProof(proofData, chunk+1, i)
	require.Equal(t, false, validProof)
}
