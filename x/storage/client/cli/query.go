package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(_ string) *cobra.Command {
	// Group storage queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListActiveDeals())
	cmd.AddCommand(CmdShowActiveDeals())
	cmd.AddCommand(CmdListProviders())
	cmd.AddCommand(CmdShowProviders())
	cmd.AddCommand(CmdFreespace())

	cmd.AddCommand(CmdFindFile())

	cmd.AddCommand(CmdGetClientFreeSpace())

	cmd.AddCommand(CmdGetPayData())

	cmd.AddCommand(CmdFileUploadCheck())

	cmd.AddCommand(CmdListStoragePaymentInfo())
	cmd.AddCommand(CmdShowStoragePaymentInfo())

	cmd.AddCommand(CmdCheckPrice())

	cmd.AddCommand(CmdListAttestForms())
	cmd.AddCommand(CmdShowAttestForms())

	cmd.AddCommand(CmdListReportForms())
	cmd.AddCommand(CmdShowReportForms())

	cmd.AddCommand(CmdListActiveProviders())
	cmd.AddCommand(CmdGetStorageStats())

	return cmd
}
