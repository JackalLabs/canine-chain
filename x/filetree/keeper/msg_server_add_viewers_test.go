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
	testutil "github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgAddViewers() {
	logger, logFile := testutil.CreateLogger()
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

	alice, err := sdkTypes.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	bob, err := sdkTypes.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	// Let it be that bob has posted a public key after signing
	// with his keyring backend using the CLI. Below is a simulation of a CLI client environment

	var ctx sdkClient.Context
	ctx.Keyring = keyring.NewInMemory() // temporary keyring

	// clientCtx.FromName in the CLI will be bob's keyring ID (bob), not the full bech32 address
	// The canined (daemon) will initialize clientCtx.FromName with bob's keyring ID taken directly from the keyring backend
	// using the --from flag
	ctx.FromName = "bob"
	algo := hd.Secp256k1

	_, _, error := ctx.Keyring.NewMnemonic(ctx.FromName, keyring.English, sdkTypes.FullFundraiserPath, keyring.DefaultBIP39Passphrase, algo)
	suite.Require().NoError(error)

	signed, _, err := ctx.Keyring.Sign(ctx.FromName, []byte("jackal_init"))
	suite.Require().NoError(error)

	k := secp256k1.GenPrivKeyFromSecret(signed)

	bobPrivateKey := eciesgo.NewPrivateKeyFromBytes(k.Bytes())

	bobPubKey := bobPrivateKey.PublicKey.Bytes(false) // to hex
	pubKeyStruct := types.Pubkey{
		Address: bob.String(),
		Key:     fmt.Sprintf("%x", bobPubKey),
	}
	suite.filetreeKeeper.SetPubkey(suite.ctx, pubKeyStruct)

	// set root folder for alice
	aliceRootFolder, err := types.CreateRootFolder(alice.String())
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceRootFolder)

	// set home folder for alice
	aliceHomeFolder, err := types.CreateFolderOrFile(alice.String(), strings.Split(alice.String(), ","), "s/home/")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceHomeFolder)

	// put pepe in home
	pepejpg, err := types.CreateFolderOrFile(alice.String(), strings.Split(alice.String(), ","), "s/home/pepe.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *pepejpg)
	pepeMerklePath := types.MerklePath("s/home/pepe.jpg")
	aliceAccountHash := types.HashThenHex(alice.String())
	ownerAddress := types.MakeOwnerAddress(pepeMerklePath, aliceAccountHash)
	bobViewerAddress := keeper.MakeViewerAddress(pepejpg.TrackingNumber, bob.String())

	// Let it be that alice has encrypted pepe using the web client, and she now has an AES IV and Key for pepe

	pepeAESKeyAndIV := "{ key: mock key, IV: mock initialisation vector } "

	// If alice wants to share pepe with bob, she will take bob's public key from chain and use ECIES to encrypt pepeAESKeyAndIV

	pubKeyReq := types.QueryGetPubkeyRequest{
		Address: bob.String(),
	}

	res, err := suite.queryClient.Pubkey(suite.ctx.Context(), &pubKeyReq)
	suite.Require().NoError(err)
	// bob uploaded his public key in hex format so we decode it from hex format
	pkey, err := eciesgo.NewPublicKeyFromHex(res.Pubkey.Key)
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
					alice.String(),
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
					alice.String(),
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
					bob.String(),
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

				fileReq := types.QueryGetFilesRequest{
					Address:      pepeMerklePath,
					OwnerAddress: ownerAddress,
				}

				res, err := suite.queryClient.Files(suite.ctx.Context(), &fileReq)
				suite.Require().NoError(err)
				logger.Println(res)

				validViewer, err := keeper.HasViewingAccess(res.Files, bob.String())
				suite.Require().NoError(err)
				suite.Require().Equal(validViewer, true)

				logger.Println(res.Files)

				// Let's see if Bob can decrypt and recover pepe's AES iv and key using his private key

				viewers := res.Files.ViewingAccess
				var m map[string]string

				error := json.Unmarshal([]byte(viewers), &m)
				suite.Require().NoError(error)

				// It was posted in hex format so lets decode it from hex first
				bytesFromHexString, err := hex.DecodeString(m[bobViewerAddress])
				suite.Require().NoError(err)

				// Just as bob used his keyring to sign and create a private-public key pair to post on chain,
				// He will use his keyring to sign again

				signed, _, err := ctx.Keyring.Sign(ctx.FromName, []byte("jackal_init"))
				suite.Require().NoError(error)

				k := secp256k1.GenPrivKeyFromSecret(signed)

				bobPrivateKeyAgain := eciesgo.NewPrivateKeyFromBytes(k.Bytes())

				decrypted, err := eciesgo.Decrypt(bobPrivateKeyAgain, bytesFromHexString)
				suite.Require().NoError(err)
				suite.Require().EqualValues(fmt.Sprintf("%s", decrypted), "{ key: mock key, IV: mock initialisation vector } ")

				logger.Println(fmt.Sprintf("%s", decrypted))

			}
		})
	}
	logFile.Close()
}
