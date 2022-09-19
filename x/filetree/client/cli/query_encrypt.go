package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	eciesgo "github.com/ecies/go/v2"
	"github.com/spf13/cobra"

	filtypes "github.com/jackal-dao/canine/x/filetree/types"
)

var _ = strconv.Itoa(0)

func CmdEncrypt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encrypt [address] [message]",
		Short: "Encrypt a message with a users address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]
			reqMessage := args[1]
			_ = reqMessage

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			key, err := sdk.AccAddressFromBech32(reqAddress)
			if err != nil {
				return err
			}

			queryClient := filtypes.NewQueryClient(clientCtx)
			res, err := queryClient.Pubkey(cmd.Context(), &filtypes.QueryGetPubkeyRequest{Address: key.String()})
			if err != nil {
				return err
			}

			fmt.Printf("ACCOUNT INFO:\n%x\n", res.Pubkey.Key)

			pkey, err := eciesgo.NewPublicKeyFromHex(res.Pubkey.Key)
			if err != nil {
				return err
			}

			encrypted, err := clientCtx.Keyring.Encrypt(pkey.Bytes(false), []byte(reqMessage))
			if err != nil {
				return err
			}

			fmt.Printf("Encrypted:\n%x\n", encrypted)

			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
