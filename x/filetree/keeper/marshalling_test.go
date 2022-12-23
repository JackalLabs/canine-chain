package keeper_test

import (
	"encoding/json"
	"fmt"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (suite *KeeperTestSuite) TestJSONMarshalling() {
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

	alice, err := sdkTypes.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	bob, err := sdkTypes.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	// set root folder for alice
	aliceRootFolder, err := types.CreateRootFolder(alice.String())
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceRootFolder)

	aliceViewerID := strings.Split(alice.String(), ",")
	aliceEditorID := aliceViewerID

	// set home folder for alice
	aliceHomeFolder, err := types.CreateFolderOrFile(alice.String(), aliceEditorID, aliceViewerID, "s/home/")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceHomeFolder)

	// put pepe in home
	// pepe's viewing access is a marshalled slice which means the keeper will fail to unmarshall
	pepejpg, err := CreateBadFile(alice.String(), aliceEditorID, aliceViewerID, "s/home/pepe.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *pepejpg)
	pepeMerklePath := types.MerklePath("s/home/pepe.jpg")
	aliceAccountHash := types.HashThenHex(alice.String())
	ownerAddress := types.MakeOwnerAddress(pepeMerklePath, aliceAccountHash)
	bobViewerAddress := keeper.MakeViewerAddress(pepejpg.TrackingNumber, bob.String())

	cases := []struct {
		preRun    func() *types.MsgAddViewers
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // alice fails to add a viewer
			preRun: func() *types.MsgAddViewers {
				return types.NewMsgAddViewers(
					alice.String(),
					bobViewerAddress,
					fmt.Sprintf("%x", "encryptedPepeAESKeyAndIV"),
					pepeMerklePath,
					ownerAddress,
				)
			},
			expErr:    true,
			name:      "alice fails to add a viewer",
			expErrMsg: "cannot unmarshall data from json",
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

			}
		})
	}
}

func CreateBadFile(creator string, editorIds []string, viewerIds []string, readablePath string) (*types.Files, error) {
	merklePath := types.MerklePath(readablePath)
	trackingNumber := uuid.NewString()

	jsonEditors, err := types.MakeEditorAccessMap(trackingNumber, editorIds, "place holder key")
	if err != nil {
		return nil, err
	}

	viewers := make([]string, 10, 10)

	for i := range viewerIds {
		viewers[i] = fmt.Sprintf("%x", "aesKey")
	}

	jsonViewers, err := json.Marshal(viewers)
	if err != nil {
		return nil, types.ErrCantMarshall
	}

	accountHash := types.HashThenHex(creator)
	ownerAddress := types.MakeOwnerAddress(merklePath, accountHash)

	File := types.Files{
		Contents:       "Contents: FID goes here",
		Owner:          ownerAddress,
		ViewingAccess:  string(jsonViewers),
		EditAccess:     string(jsonEditors),
		Address:        merklePath,
		TrackingNumber: trackingNumber,
	}

	return &File, nil
}
