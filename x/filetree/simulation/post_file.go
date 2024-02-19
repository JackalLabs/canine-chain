package simulation

import (
	"math/rand"

	"github.com/google/uuid"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func SimulateMsgPostFile(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgPostFile{
			Creator: simAccount.Address.String(),
		}

		file, err := types.CreateRootFolder(simAccount.Address.String())
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate root folder"), nil, err
		}

		_, found := k.GetFiles(ctx, file.Address, file.Owner)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to find root folder"), nil, nil
		}

		paths := GetDirectory()

		for _, path := range paths {
			_, childHash := types.MerkleHelper(path)
			ownerAddr := types.MakeOwnerAddress(childHash, types.HashThenHex(msg.Creator))
			_, found := k.GetFiles(ctx, childHash, ownerAddr)
			if !found {
				trackingNum := uuid.NewString()
				editorAccess, err := types.MakeEditorAccessMap(
					trackingNum, []string{simAccount.Address.String()}, "place holder key")
				if err != nil {
					return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate editor access map"), nil, err
				}
				msg, err = types.CreateMsgPostFile(
					simAccount.Address.String(), path, editorAccess, trackingNum)
				if err != nil {
					return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to create PostFile message"), nil, err
				}

				spendable := bk.SpendableCoins(ctx, simAccount.Address)
				fees, err := simtypes.RandomFees(r, ctx, spendable)
				if err != nil {
					return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
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

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "all files are posted"), nil, nil
	}
}
