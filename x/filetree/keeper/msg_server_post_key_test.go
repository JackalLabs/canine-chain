package keeper_test

import (
	"fmt"

	sdkClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	keyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	eciesgo "github.com/ecies/go/v2"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgPostKey() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	alice, err := sdkTypes.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	privateKey, err := makePrivateKey("alice") //clientCtx.FromName in the CLI will be alice's keyring name (alice), not the full account address
	suite.Require().NoError(err)

	pubKey := privateKey.PublicKey.Bytes(false)

	cases := []struct {
		preRun    func() *types.MsgPostkey
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgPostkey {
				return types.NewMsgPostkey(
					alice.String(),
					fmt.Sprintf("%x", pubKey),
				)
			},
			expErr: false,
			name:   "post key success",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.Postkey(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgPostkeyResponse{}, *res)

			}
		})
	}

}

// generate a mock private key using mock keyring
func makePrivateKey(fromName string) (*eciesgo.PrivateKey, error) {
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
