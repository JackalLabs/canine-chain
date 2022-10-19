package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackal-dao/canine/x/storage/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group storage queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListContracts())
	cmd.AddCommand(CmdShowContracts())
	cmd.AddCommand(CmdListProofs())
	cmd.AddCommand(CmdShowProofs())
	cmd.AddCommand(CmdListActiveDeals())
	cmd.AddCommand(CmdShowActiveDeals())
	cmd.AddCommand(CmdListProviders())
	cmd.AddCommand(CmdShowProviders())
	cmd.AddCommand(CmdFreespace())

	cmd.AddCommand(CmdFindFile())

	cmd.AddCommand(CmdListPayBlocks())
	cmd.AddCommand(CmdShowPayBlocks())
	cmd.AddCommand(CmdListClientUsage())
	cmd.AddCommand(CmdShowClientUsage())
	cmd.AddCommand(CmdListStrays())
	cmd.AddCommand(CmdShowStrays())
	cmd.AddCommand(CmdGetClientFreeSpace())

	cmd.AddCommand(CmdListFidCid())
	cmd.AddCommand(CmdShowFidCid())
	// this line is used by starport scaffolding # 1

	return cmd
}
