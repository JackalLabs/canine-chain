package main

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/spf13/cobra"
)

func StartServer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-provider",
		Short: "start jackal storage provider",
		Long:  `Start jackal storage provider`,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			StartFileServer(cmd)
			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Bool("debug", false, "allow printing info messages from the storage provider daemon")
	cmd.Flags().Uint16("interval", 30, "the interval in seconds for which to check proofs")

	return cmd
}

func SubmitProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-proof [filename] [index] [contract-id]",
		Short: "Submit merkle proof of file",
		Long:  `Submit merkle proof of file`,
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Cannot manually submit proofs.")
			// _, err := postProof(cmd, args)
			// return err
			return nil
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
