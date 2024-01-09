package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdPostFile())
	cmd.AddCommand(CmdPostFilePublic())
	cmd.AddCommand(CmdAddViewers())
	cmd.AddCommand(CmdPostkey())
	cmd.AddCommand(CmdDeleteFile())
	cmd.AddCommand(CmdRemoveViewers())
	cmd.AddCommand(CmdMakeRoot())
	cmd.AddCommand(CmdMakeRootV2())
	cmd.AddCommand(CmdAddEditors())
	cmd.AddCommand(CmdRemoveEditors())
	cmd.AddCommand(CmdResetEditors())
	cmd.AddCommand(CmdResetViewers())
	cmd.AddCommand(CmdChangeOwner())
	// this line is used by starport scaffolding # 1

	return cmd
}
