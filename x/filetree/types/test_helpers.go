package types

import (
	"crypto/sha256"
	"encoding/json"
	fmt "fmt"
	"strings"

	sdkClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	keyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	eciesgo "github.com/ecies/go/v2"
	"github.com/google/uuid"
)

// generate a mock private key using mock keyring
func MakePrivateKey(fromName string) (*eciesgo.PrivateKey, error) {
	var ctx sdkClient.Context
	ctx.Keyring = keyring.NewInMemory()
	ctx.FromName = fromName

	algo := hd.Secp256k1

	_, _, err := ctx.Keyring.NewMnemonic(fromName, keyring.English, sdkTypes.FullFundraiserPath, keyring.DefaultBIP39Passphrase, algo)
	if err != nil {
		return nil, err
	}

	signed, _, err := ctx.Keyring.Sign(ctx.FromName, []byte("jackal_init"))
	if err != nil {
		return nil, err
	}

	k := secp256k1.GenPrivKeyFromSecret(signed)

	newKey := eciesgo.NewPrivateKeyFromBytes(k.Bytes())

	return newKey, nil
}

func CreateMsgMakeRoot(creator string) (*MsgMakeRoot, error) {
	merklePath := MerklePath("s/")
	trackingNumber := uuid.NewString()

	editorIds := strings.Split(creator, ",")
	jsonEditors, err := MakeEditorAccessMap(trackingNumber, editorIds, "place holder key")
	if err != nil {
		return nil, err
	}

	accountHash := HashThenHex(creator)

	msg := NewMsgMakeRoot(
		creator,
		accountHash,
		merklePath,
		"Contents",
		string(jsonEditors),
		"Viewers",
		trackingNumber,
	)

	return msg, nil
}

func CreateRootFolder(creator string) (*Files, error) {
	merklePath := MerklePath("s/")
	trackingNumber := uuid.NewString()

	editorIds := strings.Split(creator, ",")

	jsonEditors, err := MakeEditorAccessMap(trackingNumber, editorIds, "place holder key")
	if err != nil {
		return nil, err
	}

	accountHash := HashThenHex(creator)
	ownerAddress := MakeOwnerAddress(merklePath, accountHash)

	rootFolder := Files{
		Contents:       "Contents", // This won't be used for now, but we're leaving it in as a stub in case it's needed
		Owner:          ownerAddress,
		ViewingAccess:  "NONE", // This won't be used for now, but we're leaving it in as a stub in case it's needed
		EditAccess:     string(jsonEditors),
		Address:        merklePath,
		TrackingNumber: trackingNumber,
	}

	return &rootFolder, nil
}

func CreateFolderOrFile(creator string, editorIds []string, readablePath string) (*Files, error) {
	merklePath := MerklePath(readablePath)
	trackingNumber := uuid.NewString()

	jsonEditors, err := MakeEditorAccessMap(trackingNumber, editorIds, "place holder key")
	if err != nil {
		return nil, err
	}

	accountHash := HashThenHex(creator)
	ownerAddress := MakeOwnerAddress(merklePath, accountHash)

	File := Files{
		Contents:       "Contents: FID goes here",
		Owner:          ownerAddress,
		ViewingAccess:  "Viewers",
		EditAccess:     string(jsonEditors),
		Address:        merklePath,
		TrackingNumber: trackingNumber,
	}

	return &File, nil
}

// Owner address is whoever owns this file/folder
func MakeOwnerAddress(merklePath string, user string) string {
	// make sure that user was already hex(hashed) before it was passed into
	// this function
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("o%s%s", merklePath, user)))
	hash := h.Sum(nil)
	ownerAddress := fmt.Sprintf("%x", hash)

	return ownerAddress
}

func MerkleHelper(argHashpath string) (string, string) {
	// Cut out the / at the end for compatibility with types/merkle-paths.go
	trimPath := strings.TrimSuffix(argHashpath, "/")
	chunks := strings.Split(trimPath, "/")

	parentString := strings.Join(chunks[0:len(chunks)-1], "/")
	childString := (chunks[len(chunks)-1])
	parentHash := MerklePath(parentString)

	h := sha256.New()
	h.Write([]byte(childString))
	childHash := fmt.Sprintf("%x", h.Sum(nil))

	return parentHash, childHash
}

// helper function to hash then hex any given string
func HashThenHex(any string) string {
	H := sha256.New()
	H.Write([]byte(any))
	hash := H.Sum(nil)
	return fmt.Sprintf("%x", hash)
}

func MakeEditorAccessMap(trackingNumber string, editorIds []string, aesKey string) ([]byte, error) {
	editors := make(map[string]string)

	for _, v := range editorIds {
		h := sha256.New()
		h.Write([]byte(fmt.Sprintf("e%s%s", trackingNumber, v)))
		hash := h.Sum(nil)
		addressString := fmt.Sprintf("%x", hash)

		editors[addressString] = fmt.Sprintf("%x", aesKey) // need helper function to generate a placeholder key

	}

	jsonEditors, err := json.Marshal(editors)
	if err != nil {
		return nil, ErrCantMarshall
	}

	return jsonEditors, nil
}

func CreateMsgPostFile(creator string, readablePath string, jsonEditAccess []byte, trackingNumber string) (*MsgPostFile, error) {
	accountHash := HashThenHex(creator)

	parentHash, childHash := MerkleHelper(readablePath)

	msg := NewMsgPostFile(
		creator,
		accountHash,
		parentHash,
		childHash,
		"contents: FID goes here",
		"viewers",
		string(jsonEditAccess),
		trackingNumber,
	)

	return msg, nil
}
