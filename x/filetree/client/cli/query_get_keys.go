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
	eciesgo "github.com/ecies/go/v2"
	"github.com/jackal-dao/canine/x/filetree/keeper"
	"github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetKeys() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-keys [file path] [file owner]",
		Short: "get the encryption keys from a file tree entry",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqHashpath := args[0]
			reqOwner := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pathString := types.MerklePath(reqHashpath)

			h := sha256.New()
			h.Write([]byte(fmt.Sprintf("o%s%s", pathString, reqOwner))) //May not need this in future
			hash := h.Sum(nil)
			ownerString := fmt.Sprintf("%x", hash)

			params := &types.QueryGetFilesRequest{
				Address:      pathString,
				OwnerAddress: ownerString,
			}

			res, err := queryClient.Files(context.Background(), params)
			if err != nil {
				fmt.Println("cannot find file")
				return err
			}

			viewers := res.Files.ViewingAccess
			var m map[string]string

			jerr := json.Unmarshal([]byte(viewers), &m)
			if jerr != nil {
				fmt.Println("cannot unmarshall viewers")
				return jerr
			}

			addressString := keeper.MakeViewerAddress(reqHashpath, clientCtx.GetFromAddress().String())
			fmt.Println(addressString)
			fmt.Println(m)
			todec := m[addressString]
			fmt.Printf("%v\n", todec)
			hexMessage, err := hex.DecodeString(todec)
			if err != nil {

				return err
			}

			key, err := MakePrivateKey(clientCtx)
			if err != nil {
				return err
			}

			decrypt, err := eciesgo.Decrypt(key, hexMessage)
			if err != nil {
				fmt.Printf("%v\n", hexMessage)
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
