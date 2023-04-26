package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackalLabs/canine-chain/x/amm/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListPool())
	cmd.AddCommand(CmdShowPool())
	cmd.AddCommand(CmdShowProviderRecord())
	cmd.AddCommand(CmdEstimatePoolJoin())
	cmd.AddCommand(CmdEstimateSwapOut())

	cmd.AddCommand(CmdEstimateSwapIn())

	cmd.AddCommand(CmdEstimatePoolJoin())

	cmd.AddCommand(CmdEstimatePoolExit())

	cmd.AddCommand(CmdListRecordsFromPool())

	return cmd
}
