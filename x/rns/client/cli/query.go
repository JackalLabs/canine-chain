package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackal-dao/canine/x/rns/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group rns queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
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

	// this line is used by starport scaffolding # 1

	return cmd
}
