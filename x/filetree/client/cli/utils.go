package cli

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	eciesgo "github.com/ecies/go/v2"
	filetypes "github.com/jackal-dao/canine/x/filetree/types"
)

func MakePrivateKey(clientCtx client.Context) (*eciesgo.PrivateKey, error) {
	signed, _, err := clientCtx.Keyring.Sign(clientCtx.GetFromName(), []byte("jackal_init"))
	if err != nil {
		return nil, err
	}

	k := secp256k1.GenPrivKeyFromSecret(signed)

	newKey := eciesgo.NewPrivateKeyFromBytes(k.Bytes())

	return newKey, nil
}

func merkleHelper(argHashpath string) (string, string) {

	//Cut out the / at the end for compatibility with types/merkle-paths.go
	trimPath := strings.TrimSuffix(argHashpath, "/")
	chunks := strings.Split(trimPath, "/")

	parentString := strings.Join(chunks[0:len(chunks)-1], "/")
	childString := string(chunks[len(chunks)-1])
	parentHash := filetypes.MerklePath(parentString)

	h := sha256.New()
	h.Write([]byte(childString))
	childHash := fmt.Sprintf("%x", h.Sum(nil))

	return parentHash, childHash

}

// Owner address is whoever owns this file/folder
func MakeOwnerAddress(merklePath string, user string) string {
	//make sure that user was already hex(hashed) before it was passed into
	//this function
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("o%s%s", merklePath, user)))
	hash := h.Sum(nil)
	ownerAddress := fmt.Sprintf("%x", hash)

	return ownerAddress
}
