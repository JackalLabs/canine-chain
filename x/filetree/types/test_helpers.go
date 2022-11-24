package types

import (
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	keyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	eciesgo "github.com/ecies/go/v2"
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
