package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
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

			encrypted, err := clientCtx.Keyring.Encrypt(pkey.Key, []byte(reqMessage))
			if err != nil {
				return err
			}

			fmt.Printf("ACCOUNT INFO:\n%x\n", encrypted)

			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
