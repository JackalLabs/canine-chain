package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	eciesgo "github.com/ecies/go/v2"
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
