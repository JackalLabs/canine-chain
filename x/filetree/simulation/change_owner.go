package simulation

import (
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func SimulateMsgChangeOwner(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		simBob, _ := simtypes.RandomAcc(r, accs)
		address := simAccount.Address.String()
		bob := simBob.Address.String()
		accountHash := types.HashThenHex(address)
		bobHash := types.HashThenHex(address)

		/*
			1. create share<address> file at s/home/
			2. choose another account to transfer ownership to
			3. transfer ownership
		*/
		homeFolder, err := types.CreateFolderOrFile(address, strings.Split(address, ","), strings.Split(address, ","), "s/home/")
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgChangeOwner, "unable to create home folder"), nil, err
		}

		_, found := k.GetFiles(ctx, homeFolder.Address, homeFolder.Owner)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgChangeOwner, "unable to find s/home/"), nil, nil
		}

		shareFilePath := "s/home/share" + simAccount.Address.String()
		shareFile, err := types.CreateFolderOrFile(
			address,
			strings.Split(address, ","),
			strings.Split(address, ","),
			shareFilePath,
		)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgChangeOwner, "unable to create share file"), nil, err
		}
		k.SetFiles(ctx, *shareFile)

		bobOwnerAddr := types.MakeOwnerAddress(shareFile.Address, bobHash)
		_, found = k.GetFiles(ctx, shareFile.Address, bobOwnerAddr)
		if found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgChangeOwner, "file already shared"), nil, nil
		}

		msg := &types.MsgChangeOwner{
			Creator:   address,
			Address:   shareFile.Address,
			FileOwner: accountHash,
			NewOwner:  types.HashThenHex(bob),
		}

		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgChangeOwner, "failed to generate fee"), nil, err
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
