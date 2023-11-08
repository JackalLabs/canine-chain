package keeper_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	sdkClient "github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	eciesgo "github.com/ecies/go/v2"
	"github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgAddViewers() {
	suite.SetupSuite()
	msgSrvr, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	alice := testAddresses[0]
	bob := testAddresses[1]

	// Let it be that bob has posted a public key after signing
	// with his keyring backend using the CLI. Below is a simulation of a CLI client environment

	var ctx sdkClient.Context
	ctx.Keyring = keyring.NewInMemory() // temporary keyring

	// clientCtx.FromName in the CLI will be bob's keyring ID (bob), not the full bech32 address
	// The canined (daemon) will initialize clientCtx.FromName with bob's keyring ID taken directly from the keyring backend
	// using the --from flag
	ctx.FromName = "bob"
	algo := hd.Secp256k1

	_, _, err = ctx.Keyring.NewMnemonic(ctx.FromName, keyring.English, sdkTypes.FullFundraiserPath, keyring.DefaultBIP39Passphrase, algo)
	suite.Require().NoError(err)

	signed, _, err := ctx.Keyring.Sign(ctx.FromName, []byte("jackal_init"))
	suite.Require().NoError(err)

	k := secp256k1.GenPrivKeyFromSecret(signed)

	bobPrivateKey := eciesgo.NewPrivateKeyFromBytes(k.Bytes())

	bobPubKey := bobPrivateKey.PublicKey.Bytes(false) // to hex
	pubKeyStruct := types.Pubkey{
		Address: bob,
		Key:     fmt.Sprintf("%x", bobPubKey),
	}
	suite.filetreeKeeper.SetPubkey(suite.ctx, pubKeyStruct)

	// set root folder for alice
	aliceRootFolder, err := types.CreateRootFolder(alice)
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceRootFolder)

	aliceViewerID := strings.Split(alice, ",")
	aliceEditorID := aliceViewerID

	// set home folder for alice
	aliceHomeFolder, err := types.CreateFolderOrFile(alice, aliceEditorID, aliceViewerID, "s/home/")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceHomeFolder)

	// put pepe in home
	pepejpg, err := types.CreateFolderOrFile(alice, aliceEditorID, aliceViewerID, "s/home/pepe.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *pepejpg)
	pepeMerklePath := types.MerklePath("s/home/pepe.jpg")
	aliceAccountHash := types.HashThenHex(alice)
	ownerAddress := types.MakeOwnerAddress(pepeMerklePath, aliceAccountHash)
	bobViewerAddress := keeper.MakeViewerAddress(pepejpg.TrackingNumber, bob)

	// Let it be that alice has encrypted pepe using the web client, and she now has an AES IV and Key for pepe

	pepeAESKeyAndIV := "{ key: mock key, IV: mock initialisation vector } "

	// If alice wants to share pepe with bob, she will take bob's public key from chain and use ECIES to encrypt pepeAESKeyAndIV

	pubKeyReq := types.QueryPubKeyRequest{
		Address: bob,
	}

	res, err := suite.queryClient.PubKey(suite.ctx.Context(), &pubKeyReq)
	suite.Require().NoError(err)
	// bob uploaded his public key in hex format so we decode it from hex format
	pkey, err := eciesgo.NewPublicKeyFromHex(res.PubKey.Key)
	suite.Require().NoError(err)

	encryptedPepeAESKeyAndIV, err := eciesgo.Encrypt(pkey, []byte(pepeAESKeyAndIV)) // convert to hex
	suite.Require().NoError(err)

	cases := []struct {
		preRun    func() *types.MsgAddViewers
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // alice adds a viewer
			preRun: func() *types.MsgAddViewers {
				return types.NewMsgAddViewers(
					alice,
					bobViewerAddress,
					fmt.Sprintf("%x", encryptedPepeAESKeyAndIV),
					pepeMerklePath,
					ownerAddress,
				)
			},
			expErr: false,
			name:   "alice adds a viewer",
		},
		{ // alice fails to add a viewer to a non existent file
			preRun: func() *types.MsgAddViewers {
				return types.NewMsgAddViewers(
					alice,
					bobViewerAddress,
					fmt.Sprintf("%x", encryptedPepeAESKeyAndIV),
					types.MerklePath("s/home/ghost.jpg"),
					ownerAddress,
				)
			},
			expErr:    true,
			expErrMsg: "file not found",
			name:      "alice cannot share ghosts",
		},
		{ // bob doesn't own pepe so he can't share it
			preRun: func() *types.MsgAddViewers {
				return types.NewMsgAddViewers(
					bob,
					bobViewerAddress,
					fmt.Sprintf("%x", encryptedPepeAESKeyAndIV),
					pepeMerklePath,
					ownerAddress,
				)
			},
			expErr:    true,
			expErrMsg: "Unathorized. Only the owner can add a viewer.",
			name:      "bob can't share what he doesn't own",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.AddViewers(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {

				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgAddViewersResponse{}, *res)
				// Let's confirm that bob is a viewer

				fileReq := types.QueryFileRequest{
					Address:      pepeMerklePath,
					OwnerAddress: ownerAddress,
				}

				res, err := suite.queryClient.Files(suite.ctx.Context(), &fileReq)
				suite.Require().NoError(err)

				validViewer, err := keeper.HasViewingAccess(res.Files, bob)
				suite.Require().NoError(err)
				suite.Require().Equal(validViewer, true)

				// Let's see if Bob can decrypt and recover pepe's AES iv and key using his private key

				viewers := res.Files.ViewingAccess
				var m map[string]string

				err = json.Unmarshal([]byte(viewers), &m)
				suite.Require().NoError(err)

				// It was posted in hex format so lets decode it from hex first
				bytesFromHexString, err := hex.DecodeString(m[bobViewerAddress])
				suite.Require().NoError(err)

				// Just as bob used his keyring to sign and create a private-public key pair to post on chain,
				// He will use his keyring to sign again

				signed, _, err := ctx.Keyring.Sign(ctx.FromName, []byte("jackal_init"))
				suite.Require().NoError(err)

				k := secp256k1.GenPrivKeyFromSecret(signed)

				bobPrivateKeyAgain := eciesgo.NewPrivateKeyFromBytes(k.Bytes())

				decrypted, err := eciesgo.Decrypt(bobPrivateKeyAgain, bytesFromHexString)
				suite.Require().NoError(err)
				suite.Require().EqualValues(string(decrypted), "{ key: mock key, IV: mock initialisation vector } ")

			}
		})
	}
}
