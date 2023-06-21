package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackalLabs/canine-chain/v3/x/notifications/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(_ string) *cobra.Command {
	// Group notifications queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// Commenting out some useless queries for now
	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListNotifications())
	// cmd.AddCommand(CmdShowNotifications())

	cmd.AddCommand(CmdListNotiCounter())
	// cmd.AddCommand(CmdShowNotiCounter())
	cmd.AddCommand(CmdListNotificationsByAddress())

	return cmd
}
