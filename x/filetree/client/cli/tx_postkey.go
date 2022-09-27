package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	eciesgo "github.com/ecies/go/v2"
	"github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPostkey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "postkey",
		Short: "Posts a users generated public key for the encryption model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fmt.Println("clientCtx.GetFromName () is", clientCtx.GetFromName())
			//make a private key
			signed, publicKey, err := clientCtx.Keyring.Sign(clientCtx.GetFromName(), []byte("jackal_init"))
			if err != nil {
				return err
			}

			fmt.Printf("signed is %x\n", signed)
			fmt.Println("public key from signed is", publicKey)

			k := secp256k1.GenPrivKeyFromSecret(signed)

			fmt.Printf("private key is %v\n", *k) //dereference the pointer to show the value it's holding and print the value
			fmt.Printf("private key as hex is %x\n", *k)
			fmt.Printf("k.Bytes() is %v\n", k.Bytes())

			newKey := eciesgo.NewPrivateKeyFromBytes(k.Bytes())
			fmt.Printf("New private key is %v\n", *newKey) //dereference the value that the pointer is holding

			nestedPubKey := newKey.PublicKey
			fmt.Printf("Public key inside is %v\n", *nestedPubKey)

			if err != nil {
				return err
			}

			pubKey := newKey.PublicKey.Bytes(false)

			fmt.Println("public key is", pubKey)
			fmt.Printf("public key as hex is %x", pubKey)
			fmt.Println("from address is ", clientCtx.GetFromAddress().String())
			os.Exit(0)

			msg := types.NewMsgPostkey(
				clientCtx.GetFromAddress().String(),
				fmt.Sprintf("%x", pubKey),
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
