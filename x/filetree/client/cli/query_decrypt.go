package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDecrypt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decrypt [message]",
		Short: "Decrypt some message with your private key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqMessage := args[0]

			hexMessage, err := hex.DecodeString(reqMessage)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.From

			decrypt, _, err := clientCtx.Keyring.Decrypt(from, hexMessage)
			if err != nil {
				return err
			}

			fmt.Printf("DECRYPTED:\n%s\n", string(decrypt))

			return nil
		},
	}
	cmd.Flags().String(flags.FlagFrom, "", "Name or address of private key with which to sign")

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
