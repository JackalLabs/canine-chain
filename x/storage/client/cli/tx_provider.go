package cli

import (
	"strconv"

	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdProvider() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider [command]",
		Short: "Provider subcommands",
		Long:  "The provider category for the canine cli tool.",
	}

	cmd.AddCommand(CmdSetProviderIP())
	cmd.AddCommand(CmdSetProviderTotalspace())
	cmd.AddCommand(CmdInitProvider())
	cmd.AddCommand(CmdSetProviderKeybase())

	return cmd
}
