package cli

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPostFile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post-file [hashpath] [contents] [keys] [viewers] [editors]",
		Short: "post a new file to your file explorer",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argHashpath := args[0]
			argContents := args[1]
			argKeys := args[2]
			argViewers := args[3]
			argEditors := args[4]

			_ = argViewers

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			key, err := sdk.AccAddressFromBech32(clientCtx.GetFromAddress().String())
			if err != nil {
				return err
			}

			queryClient := authtypes.NewQueryClient(clientCtx)
			res, err := queryClient.Account(cmd.Context(), &authtypes.QueryAccountRequest{Address: key.String()})
			if err != nil {
				return err
			}

			fmt.Println(res.Account.TypeUrl)

			var acc authtypes.BaseAccount

			err = acc.Unmarshal(res.Account.Value)
			if err != nil {
				return err
			}
			var pkey secp256k1.PubKey

			err = pkey.Unmarshal(acc.PubKey.Value)
			if err != nil {
				return err
			}

			fmt.Printf("ACCOUNT INFO:\n%x\n", pkey.Key)

			encrypted, err := clientCtx.Keyring.Encrypt(pkey.Key, []byte(argKeys))
			if err != nil {
				return err
			}

			h := sha256.New()
			h.Write([]byte(argHashpath))
			hash := h.Sum(nil)

			pathString := fmt.Sprintf("%x", hash)

			viewers := make(map[string]string)

			h = sha256.New()
			h.Write([]byte(clientCtx.GetFromAddress().String()))
			hash = h.Sum(nil)

			addressString := fmt.Sprintf("%x", hash)

			viewers[addressString] = fmt.Sprintf("%x", encrypted)

			jsonViewers, err := json.Marshal(viewers)
			if err != nil {
				return err
			}

			msg := types.NewMsgPostFile(
				clientCtx.GetFromAddress().String(),
				pathString,
				argContents,
				string(jsonViewers),
				argEditors,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
