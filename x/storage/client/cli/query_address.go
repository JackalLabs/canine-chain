package cli

import (
	"fmt"

	allTypes "github.com/jackalLabs/canine-chain/v4/types"

	"github.com/spf13/cobra"
)

func CmdQueryAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "address",
		Short: "get an accounts address",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			acc, err := allTypes.GetAccount(args[0])
			if err != nil {
				return err
			}

			fmt.Printf("Account address:\n\t%s -> %s\n", args[0], acc.String())
			return nil
		},
	}

	return cmd
}
