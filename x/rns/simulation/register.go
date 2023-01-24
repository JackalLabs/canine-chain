package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackalLabs/canine-chain/x/rns/keeper"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

func StringWithCharset(r *rand.Rand, length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}

func SimulateMsgRegister(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRegister{
			Creator: simAccount.Address.String(),
		}
		// generating a random name
		// generating a random TLD
		tldIndex := simtypes.RandIntBetween(r, 0, len(types.SupportedTLDs))
		tld := types.SupportedTLDs[tldIndex]
		// generating a random name
		nameLength := simtypes.RandIntBetween(r, 1, 10)
		name := StringWithCharset(r, nameLength, charset)

		// calculating the owner
		owner, err := sdk.AccAddressFromBech32(string(simAccount.Address))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Can't parse sender"), nil, nil
		}

		// generating a random time in years
		numYears := simtypes.RandIntBetween(r, 1, 15)
		blockHeight := ctx.BlockTime()
		time := int64(numYears) * 5733818

		// checking if the domain already exists on-chain
		whois, isFound := k.GetNames(ctx, name, tld)
		if isFound {
			if whois.Value == owner.String() {
				time = whois.Expires + time
			} else if blockHeight.Unix() < whois.Expires {
				return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Name is already registered"), nil, nil
			}
		} else {
			time += blockHeight.Unix()
		}

		// calculating the necessary costs to rent the domain
		domainPrice, err := GetCostOfName(name, tld)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Grabbing cost of name"), nil, nil
		}
		var cost sdk.Int
		cost = sdk.NewInt(domainPrice)

		// ensuring the account has enough coins to buy the domain
		jBalance := bk.GetBalance(ctx, simAccount.Address, "ujkl")
		if jBalance.Amount.LTE(cost) {
			c := sdk.NewCoin("ujkl", cost)

			err := bk.MintCoins(ctx, types.ModuleName, sdk.NewCoins(c))
			if err != nil {
				return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unabled to fund account"), nil, err
			}

			err = bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, simAccount.Address, sdk.NewCoins(c))
			if err != nil {
				return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unabled to fund account"), nil, err
			}
		}

		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		coins, hasNeg := spendable.SafeSub(sdk.NewCoins(sdk.NewCoin("ujkl", cost)))

		var fees sdk.Coins

		if !hasNeg {
			var err error
			fees, err = simtypes.RandomFees(r, ctx, coins)
			if err != nil {
				return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
			}
		}

		// generating the transaction
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
