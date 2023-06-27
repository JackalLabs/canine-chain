package simulation

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	eciesgo "github.com/ecies/go/v2"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func SimulateMsgAddEditors(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		merklePath := types.MerklePath("s/home/")
		simAccount, _ := simtypes.RandomAcc(r, accs)
		simBob, _ := simtypes.RandomAcc(r, accs)
		address, _ := sdk.Bech32ifyAddressBytes("jkl", simAccount.Address)
		bob, _ := sdk.Bech32ifyAddressBytes("jkl", simBob.Address)
		accountHash := types.HashThenHex(address)
		ownerAddress := types.MakeOwnerAddress(merklePath, accountHash)

		// root folder
		rootFolder, err := types.CreateRootFolder(address)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddEditors, "unable to create root folder"), nil, err
		}
		k.SetFiles(ctx, *rootFolder)

		// home folder
		homeFolder, err := types.CreateFolderOrFile(address, strings.Split(address, ","), strings.Split(address, ","), "s/home/")
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddEditors, "unable to create home folder"), nil, err
		}
		k.SetFiles(ctx, *homeFolder)

		// bob as an editor for home folder
		// get editor id and editor key for bob
		editIds := keeper.MakeEditorAddress(homeFolder.TrackingNumber, bob)

		mockKeyAndIV := "{ key: mock key, IV: mock initialisation vector } "
		pkeyHex := fmt.Sprintf("%x", simBob.PubKey.Bytes())
		pkey, _ := eciesgo.NewPublicKeyFromHex(pkeyHex)
		encryptedKeyAndIV, _ := eciesgo.Encrypt(pkey, []byte(mockKeyAndIV))
		editKeys := fmt.Sprintf("%x", encryptedKeyAndIV)

		msg := &types.MsgAddEditors{
			Creator:    simAccount.Address.String(),
			EditorIds:  editIds,
			EditorKeys: editKeys,
			Address:    merklePath,
			Fileowner:  ownerAddress,
		}

		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddEditors, "failed to generate fee"), nil, err
		}

		txCtx := simulation.OperationInput{
			R:             r,
			App:           app,
			TxGen:         simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:           nil,
			Msg:           msg,
			MsgType:       msg.Type(),
			Context:       ctx,
			SimAccount:    simAccount,
			AccountKeeper: ak,
			ModuleName:    types.ModuleName,
		}

		return simulation.GenAndDeliverTx(txCtx, fees)
	}
}
