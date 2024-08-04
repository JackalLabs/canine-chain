package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackalLabs/canine-chain/v4/x/filetree/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	_ = queryRoute
	// Group filetree queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())

	cmd.AddCommand(CmdListFiles())
	cmd.AddCommand(CmdShowFiles())
	cmd.AddCommand(CmdGetKeys())

	cmd.AddCommand(CmdListPubkey())
	cmd.AddCommand(CmdShowPubkey())
	cmd.AddCommand(CmdShowFileFromPath())
	// this line is used by starport scaffolding # 1

	return cmd
}
