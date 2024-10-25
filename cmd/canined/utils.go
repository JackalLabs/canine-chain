package main

import (
	"fmt"

	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"github.com/spf13/cobra"
)

func AddressGenerationCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-address [value]",
		Short: "Generate an address from text, the same system used to generate protocol owned accounts.",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			acc, err := types.GetAccount(args[0])
			if err != nil {
				return err
			}

			fmt.Printf("'%s' becomes: '%s'", args[0], acc.String())

			return nil
		},
	}

	return cmd
}
