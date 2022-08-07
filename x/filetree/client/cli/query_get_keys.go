package cli

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetKeys() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-keys [hashpath]",
		Short: "get the encryption keys from a file tree entry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqHashpath := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetFilesRequest{
				Address: reqHashpath,
			}

			res, err := queryClient.Files(context.Background(), params)
			if err != nil {
				return err
			}

			viewers := res.Files.ViewingAccess
			var m map[string]string

			json.Unmarshal([]byte(viewers), &m)

			h := sha256.New()
			h.Write([]byte(clientCtx.GetFromAddress().String()))
			hash := h.Sum(nil)
			addressString := fmt.Sprintf("%x", hash)

			hexMessage, err := hex.DecodeString(m[addressString])
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
