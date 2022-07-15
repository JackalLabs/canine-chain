package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackal-dao/canine/x/storage/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdPostContract())
	cmd.AddCommand(CmdCreateContracts())
	cmd.AddCommand(CmdUpdateContracts())
	cmd.AddCommand(CmdDeleteContracts())
	cmd.AddCommand(CmdCreateProofs())
	cmd.AddCommand(CmdUpdateProofs())
	cmd.AddCommand(CmdDeleteProofs())
	cmd.AddCommand(CmdItem())
	cmd.AddCommand(CmdPostproof())
	cmd.AddCommand(CmdCreateActiveDeals())
	cmd.AddCommand(CmdUpdateActiveDeals())
	cmd.AddCommand(CmdDeleteActiveDeals())
	cmd.AddCommand(CmdSignContract())
	cmd.AddCommand(CmdCreateMiners())
	cmd.AddCommand(CmdUpdateMiners())
	cmd.AddCommand(CmdDeleteMiners())
	cmd.AddCommand(CmdSetMinerIp())
	cmd.AddCommand(CmdSetMinerTotalspace())
	cmd.AddCommand(CmdInitMiner())
	// this line is used by starport scaffolding # 1

	return cmd
}
