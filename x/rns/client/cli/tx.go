package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

//nolint:unused
var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

//nolint:unused
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

	cmd.AddCommand(CmdRegister())
	cmd.AddCommand(CmdBid())
	cmd.AddCommand(CmdAcceptBid())
	cmd.AddCommand(CmdCancelBid())
	cmd.AddCommand(CmdList())
	cmd.AddCommand(CmdBuy())
	cmd.AddCommand(CmdDelist())
	cmd.AddCommand(CmdTransfer())
	cmd.AddCommand(CmdAddRecord())
	cmd.AddCommand(CmdDelRecord())
	cmd.AddCommand(CmdInit())
	cmd.AddCommand(CmdUpdate())

	return cmd
}
