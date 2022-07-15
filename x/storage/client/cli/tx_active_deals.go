package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/storage/types"
	"github.com/spf13/cobra"
)

func CmdCreateActiveDeals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-active-deals [cid] [signee] [miner] [startblock] [endblock] [filesize] [proofverified] [proofsmissed] [blocktoprove]",
		Short: "Create a new active_deals",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexCid := args[0]

			// Get value arguments
			argSignee := args[1]
			argMiner := args[2]
			argStartblock := args[3]
			argEndblock := args[4]
			argFilesize := args[5]
			argProofverified := args[6]
			argProofsmissed := args[7]
			argBlocktoprove := args[8]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateActiveDeals(
				clientCtx.GetFromAddress().String(),
				indexCid,
				argSignee,
				argMiner,
				argStartblock,
				argEndblock,
				argFilesize,
				argProofverified,
				argProofsmissed,
				argBlocktoprove,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateActiveDeals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-active-deals [cid] [signee] [miner] [startblock] [endblock] [filesize] [proofverified] [proofsmissed] [blocktoprove]",
		Short: "Update a active_deals",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexCid := args[0]

			// Get value arguments
			argSignee := args[1]
			argMiner := args[2]
			argStartblock := args[3]
			argEndblock := args[4]
			argFilesize := args[5]
			argProofverified := args[6]
			argProofsmissed := args[7]
			argBlocktoprove := args[8]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateActiveDeals(
				clientCtx.GetFromAddress().String(),
				indexCid,
				argSignee,
				argMiner,
				argStartblock,
				argEndblock,
				argFilesize,
				argProofverified,
				argProofsmissed,
				argBlocktoprove,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteActiveDeals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-active-deals [cid]",
		Short: "Delete a active_deals",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexCid := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteActiveDeals(
				clientCtx.GetFromAddress().String(),
				indexCid,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
