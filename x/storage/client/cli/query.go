package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
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
		CmdListActiveDealsByMerkle(),
		CmdListActiveDeals(),
		CmdQueryParams(),
		CmdShowProof(),
		CmdFindFile(),
		CmdOpenFiles(),
	)

	return cmd
}
