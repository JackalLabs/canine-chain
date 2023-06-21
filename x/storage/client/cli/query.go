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
	cmd.AddCommand(CmdListContracts())
	cmd.AddCommand(CmdShowContracts())
	cmd.AddCommand(CmdListActiveDeals())
	cmd.AddCommand(CmdShowActiveDeals())
	cmd.AddCommand(CmdListProviders())
	cmd.AddCommand(CmdShowProviders())
	cmd.AddCommand(CmdFreespace())

	cmd.AddCommand(CmdFindFile())

	cmd.AddCommand(CmdListStrays())
	cmd.AddCommand(CmdShowStrays())
	cmd.AddCommand(CmdGetClientFreeSpace())

	cmd.AddCommand(CmdListFidCid())
	cmd.AddCommand(CmdShowFidCid())
	cmd.AddCommand(CmdGetPayData())

	cmd.AddCommand(CmdFileUploadCheck())

	cmd.AddCommand(CmdListStoragePaymentInfo())
	cmd.AddCommand(CmdShowStoragePaymentInfo())

	cmd.AddCommand(CmdCheckPrice())

	cmd.AddCommand(CmdListAttestForms())
	cmd.AddCommand(CmdShowAttestForms())

	cmd.AddCommand(CmdListActiveProviders())
	// this line is used by starport scaffolding # 1
	cmd.AddCommand(CmdGetStorageStats())

	return cmd
}
