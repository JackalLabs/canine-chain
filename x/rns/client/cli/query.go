package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(_ string) *cobra.Command {
	// Group rns queries under a subcommand
	cmd := &cobra.Command{
		Use: types.ModuleName,
		Short: fmt.Sprintf(
			"Querying commands for the %s module",
			types.ModuleName,
		),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListNames())
	cmd.AddCommand(CmdShowNames())
	cmd.AddCommand(CmdListBids())
	cmd.AddCommand(CmdShowBids())
	cmd.AddCommand(CmdListForsale())
	cmd.AddCommand(CmdShowForsale())
	cmd.AddCommand(CmdListInit())
	cmd.AddCommand(CmdShowInit())
	cmd.AddCommand(CmdListOwnedNames())

	return cmd
}
