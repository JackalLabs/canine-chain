package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(_ string) *cobra.Command {
	// Group storage queries under a subcommand
	cmd := &cobra.Command{
		Use: types.ModuleName,
		Short: fmt.Sprintf(
			"Querying commands for the %s module",
			types.ModuleName,
		),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdGetStorageStats(),
		CmdListActiveProviders(),
		CmdShowReportForms(),
		CmdListReportForms(),
		CmdShowAttestForms(),
		CmdListAttestForms(),
		CmdCheckPrice(),
		CmdQueryAddress(),
		CmdShowStoragePaymentInfo(),
		CmdListStoragePaymentInfo(),
		CmdFileUploadCheck(),
		CmdGetPayData(),
		CmdGetClientFreeSpace(),
		CmdFreespace(),
		CmdShowProviders(),
		CmdListProviders(),
		CmdShowActiveDeals(),
		CmdListActiveDeals(),
		CmdQueryParams(),
		CmdShowProof(),
		CmdFindFile(),
		CmdOpenFiles(),
	)

	return cmd
}
