package cli

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/cosmos/cosmos-sdk/client/input"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

type ErrorResponse struct {
	Error string `json:"error"`
}

type UploadResponse struct {
	CID string `json:"cid"`
	FID string `json:"fid"`
}

func prepareFactory(clientCtx client.Context, txf tx.Factory) (tx.Factory, error) {
	from := clientCtx.GetFromAddress()

	if err := txf.AccountRetriever().EnsureExists(clientCtx, from); err != nil {
		return txf, err
	}

	initNum, initSeq := txf.AccountNumber(), txf.Sequence()
	if initNum == 0 || initSeq == 0 {
		num, seq, err := txf.AccountRetriever().GetAccountNumberSequence(clientCtx, from)
		if err != nil {
			return txf, err
		}

		if initNum == 0 {
			txf = txf.WithAccountNumber(num)
		}

		if initSeq == 0 {
			txf = txf.WithSequence(seq)
		}
	}

	return txf, nil
}

// GenerateOrBroadcastTx is some dumb wrapper I had to make cause the sdk assumes I don't want to programmatically handle the
// response but instead print it out like a doofus
func GenerateOrBroadcastTx(clientCtx client.Context, flags *pflag.FlagSet, msgs ...sdk.Msg) (*sdk.TxResponse, error) {
	txf := tx.NewFactoryCLI(clientCtx, flags)

	for _, msg := range msgs {
		if err := msg.ValidateBasic(); err != nil {
			return nil, err
		}
	}

	txf, err := prepareFactory(clientCtx, txf)
	if err != nil {
		return nil, err
	}

	if txf.SimulateAndExecute() || clientCtx.Simulate {
		_, adjusted, err := tx.CalculateGas(clientCtx, txf, msgs...)
		if err != nil {
			return nil, err
		}

		txf = txf.WithGas(adjusted)
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", tx.GasEstimateResponse{GasEstimate: txf.Gas()})
	}

	if clientCtx.Simulate {
		return nil, err
	}

	txn, err := tx.BuildUnsignedTx(txf, msgs...)
	if err != nil {
		return nil, err
	}

	if !clientCtx.SkipConfirm {
		out, err := clientCtx.TxConfig.TxJSONEncoder()(txn.GetTx())
		if err != nil {
			return nil, err
		}

		_, _ = fmt.Fprintf(os.Stderr, "%s\n\n", out)

		buf := bufio.NewReader(os.Stdin)
		ok, err := input.GetConfirmation("confirm transaction before signing and broadcasting", buf, os.Stderr)

		if err != nil || !ok {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", "cancelled transaction")
			return nil, err
		}
	}

	txn.SetFeeGranter(clientCtx.GetFeeGranterAddress())
	err = tx.Sign(txf, clientCtx.GetFromName(), txn, true)
	if err != nil {
		return nil, err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(txn.GetTx())
	if err != nil {
		return nil, err
	}

	// broadcast to a Tendermint node
	return clientCtx.BroadcastTxCommit(txBytes)
}

func uploadFile(ip string, r io.Reader, address string) (string, error) {
	cli := http.DefaultClient

	u, err := url.Parse(ip)
	if err != nil {
		return "", err
	}

	u = u.JoinPath("upload")

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	defer writer.Close()

	err = writer.WriteField("sender", address)
	if err != nil {
		return "", err
	}

	fileWriter, err := writer.CreateFormFile("file", "file")
	if err != nil {
		return "", err
	}

	_, err = io.Copy(fileWriter, r)
	if err != nil {
		return "", err
	}
	writer.Close()

	req, _ := http.NewRequest("POST", u.String(), &b)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	res, err := cli.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {

		var errRes ErrorResponse

		err := json.NewDecoder(res.Body).Decode(&errRes)
		if err != nil {
			return "", err
		}

		return "", fmt.Errorf("upload failed with code %d | %s", res.StatusCode, errRes.Error)
	}

	var ures UploadResponse
	err = json.NewDecoder(res.Body).Decode(&ures)
	if err != nil {
		return "", err
	}

	return ures.CID, nil
}

func postFile(fileData []byte, cmd *cobra.Command) {
	buf := bytes.NewBuffer(fileData)
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		panic(err)
	}

	address := clientCtx.GetFromAddress().String()

	ip := "http://198.244.165.95:3333/"

	uploadBuffer := bytes.NewBuffer(buf.Bytes())
	cid, err := uploadFile(ip, uploadBuffer, address)
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(time.Second * 1)

	msg := types.NewMsgSignContract(
		address,
		cid,
		true,
	)
	if err := msg.ValidateBasic(); err != nil {
		panic(err)
	}

	res, err := GenerateOrBroadcastTx(clientCtx, cmd.Flags(), msg)
	if err != nil {
		panic(err)
	}

	if res != nil {
		fmt.Println(res.RawLog)
		return
	}
	if res.Code != 0 {
		panic("tx failed!")
	}
}

func CmdPostRandomFile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post-random [p-count]",
		Short: "Post random file to chain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			countArg := args[0]
			count, err := strconv.ParseInt(countArg, 10, 64)
			if err != nil {
				panic(err)
			}

			url := fmt.Sprintf("https://baconipsum.com/api/?type=meat-and-filler&paras=%d&format=text", count)
			hcli := http.DefaultClient
			resp, err := hcli.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return nil
			}
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			postFile(bodyBytes, cmd)

			return nil
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdPostFile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post [file-path]",
		Short: "Post file to chain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			filePath := args[0]

			file, err := os.ReadFile(filePath)
			if err != nil {
				return err
			}

			postFile(file, cmd)

			return nil
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
