package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdProvider() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider [command]",
		Short: "Provider subcommands",
		Long:  "The provider category for the canine cli tool.",
	}

	cmds := []*cobra.Command{
		CmdSetProviderTotalspace(),
		CmdSetProviderIP(),
		CmdSetProviderKeybase(),
		CmdInitProvider(),
	}

	for _, c := range cmds {
		flags.AddTxFlagsToCmd(c)
		cmd.AddCommand(c)
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
