package cli

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/filetree/types"
	filetypes "github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdInitAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init-account [root-hashpath] [account] [editors]",
		Short: "Broadcast message InitAccount",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRootHashpath := args[0]
			argAccount := args[1]
			argEditors := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			newKey, err := MakePrivateKey(clientCtx)

			if err != nil {
				return err
			}

			pubKey := newKey.PublicKey.Bytes(false)
			//In the keeper, the merklePath function will trim the trailing slash for us but let's just do it anyways to be safe.
			trimMerklePath := strings.TrimSuffix(argRootHashpath, "/")
			merklePath := types.MerklePath(trimMerklePath)
			fmt.Println("The merkle path is", merklePath)
			editors := make(map[string]string)
			editorAddresses := strings.Split(argEditors, ",")
			editorAddresses = append(editorAddresses, clientCtx.GetFromAddress().String())

			//Getting the tracker from the client side safe? By the time your transaction is done, the tracker would have been incremented by many other transactions
			queryClient := filetypes.NewQueryClient(clientCtx)
			res, err := queryClient.Tracker(cmd.Context(), &filetypes.QueryGetTrackerRequest{})
			if err != nil {
				return types.ErrTrackerNotFound
			}
			trackingNumber := res.Tracker.TrackingNumber
			fmt.Println("Tracking number is", trackingNumber)

			for _, v := range editorAddresses {
				if len(v) < 1 {
					continue
				}

				//This root folder is not supposed to hold the file's AES key, so there is nothing to encrypt. The purpose
				//Of the list of editors is to allow a user to invite others to write to their root folder.

				h := sha256.New()
				h.Write([]byte(fmt.Sprintf("e%d%s", trackingNumber, v)))
				hash := h.Sum(nil)

				addressString := fmt.Sprintf("%x", hash)

				editors[addressString] = fmt.Sprintf("%x", "NoKeyHere") //No need to store a key here
			}

			jsonEditors, err := json.Marshal(editors)
			if err != nil {
				return err
			}

			fmt.Println("argAccount bech32 is", argAccount)

			h := sha256.New()
			h.Write([]byte(fmt.Sprintf("%s", argAccount)))
			hash := h.Sum(nil)

			accountHash := fmt.Sprintf("%x", hash)

			fmt.Println("accountHash is", accountHash)

			msgInitRoot := types.NewMsgInitAccount(
				clientCtx.GetFromAddress().String(),
				accountHash,
				merklePath,
				string(jsonEditors),
				fmt.Sprintf("%x", pubKey),
				trackingNumber,
			)

			if err := msgInitRoot.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgInitRoot)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
