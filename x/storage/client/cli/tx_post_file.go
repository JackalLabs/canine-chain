package cli

import (
	"bufio"
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/cosmos/cosmos-sdk/client/input"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/utils"
	"github.com/spf13/pflag"
	"github.com/tendermint/tendermint/libs/rand"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

type ErrorResponse struct {
	Error string `json:"error"`
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

func uploadFile(ip string, r io.Reader, merkle []byte, start int64, address string) error {
	cli := http.DefaultClient

	u, err := url.Parse(ip)
	if err != nil {
		return err
	}

	u = u.JoinPath("v2/upload")

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	defer writer.Close()

	err = writer.WriteField("sender", address)
	if err != nil {
		return err
	}

	fmt.Println(hex.EncodeToString(merkle))

	err = writer.WriteField("merkle", hex.EncodeToString(merkle))
	if err != nil {
		return err
	}

	err = writer.WriteField("start", fmt.Sprintf("%d", start))
	if err != nil {
		return err
	}

	fileWriter, err := writer.CreateFormFile("file", hex.EncodeToString(merkle))
	if err != nil {
		return err
	}

	_, err = io.Copy(fileWriter, r)
	if err != nil {
		return err
	}
	writer.Close()

	req, _ := http.NewRequest("POST", u.String(), &b)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	res, err := cli.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 202 {

		var errRes ErrorResponse

		err := json.NewDecoder(res.Body).Decode(&errRes)
		if err != nil {
			return err
		}

		return fmt.Errorf("upload failed with code %d | %s", res.StatusCode, errRes.Error)
	}

	bb, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(bb))

	return nil
}

func createMerkleRoot(file io.Reader, chunkSize int64) ([]byte, error) {
	root, _, _, _, err := utils.BuildTree(file, chunkSize)
	return root, err
}

func postFileToChain(ctx client.Context, flags *pflag.FlagSet, merkle []byte, fileSize, maxProofs, expires int64) (startat int64, err error) {
	msg := types.NewMsgPostFile(
		ctx.GetFromAddress().String(),
		merkle,
		fileSize,
		40,
		0,
		maxProofs,
		`{"note":"Uploaded with canined"}`)
	msg.Expires = expires
	if err := msg.ValidateBasic(); err != nil {
		return 0, err
	}

	res, err := GenerateOrBroadcastTx(ctx, flags, msg)
	if err != nil {
		return 0, err
	}
	if res != nil {
		fmt.Println(res.RawLog)
	}
	if res.Code != 0 {
		return 0, errors.New("tx failed")
	}

	startatStr := ""
find:
	for _, event := range res.Events {
		if event.Type != "post_file" {
			continue
		}

		for _, attr := range event.Attributes {
			if string(attr.Key) == "start" {
				startatStr = string(attr.Value)
				break find
			}
		}
	}

	if startatStr == "" {
		panic(errors.New("start block event attribute not found in tx response"))
	}

	startat, err = strconv.ParseInt(startatStr, 10, 64)
	if err != nil {
		panic(err)
	}

	return startat, nil
}

func handlePost(ctx client.Context, flags *pflag.FlagSet, filename string, ips []string, expires, maxProofs int64) error {
	query := types.NewQueryClient(ctx)
	params, err := query.Params(context.Background(), &types.QueryParams{})
	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	root, err := createMerkleRoot(file, params.Params.ChunkSize)
	if err != nil {
		return errors.Join(errors.New("failed to create merkle root of file"), err)
	}

	startat, err := postFileToChain(ctx, flags, root, stat.Size(), maxProofs, expires)
	if err != nil {
		return err
	}

	for _, ip := range ips {
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			panic(err)
		}
		err = uploadFile(ip, file, root, startat, ctx.GetFromAddress().String())
		if err != nil {
			fmt.Printf("failed to upload file to provider: %v", err)
		}
	}

	return err
}

func getProvidersToUpload(ctx client.Context, dest string, count int64) (ips []string, err error) {
	query := types.NewQueryClient(ctx)
	if dest != "" {
		ips = append(ips, dest)
	}
	ips = append(ips, []string{
		"https://mprov01.jackallabs.io",
		"https://mprov02.jackallabs.io",
		"https://jklstorage1.squirrellogic.com",
		"https://jklstorage2.squirrellogic.com",
		"https://jklstorage3.squirrellogic.com",
	}...)

	if len(ips) > int(count) {
		return ips[:count], nil
	}

	res, err := query.ActiveProviders(
		context.Background(),
		&types.QueryActiveProviders{})
	if err != nil {
		return nil, errors.Join(errors.New("failed to find providers"), err)
	}
	if len(res.Providers) == 0 {
		return nil, errors.New("there are no active providers on chain")
	}

	info, err := ctx.Client.ABCIInfo(context.Background())
	if err != nil {
		return nil, err
	}
	r := rand.NewRand()
	r.Seed(info.Response.LastBlockHeight)

	fill := min((int(count) - len(ips)), len(res.Providers))

	// randomly pick active providers
	for range fill {
		i := r.Int() % len(res.Providers)
		pick := res.Providers[i]
		res.Providers = append(res.Providers[:i], res.Providers[i+1:]...)

		prov, err := query.Provider(
			context.Background(),
			&types.QueryProvider{Address: pick.Address})
		if err != nil {
			return nil, err
		}

		ips = append(ips, prov.Provider.Ip)
	}

	return ips, nil
}

func CmdPostFile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post [file-path] [expire-block]",
		Short: "Post file to chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			filePath := args[0]
			expireBlock, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return errors.Join(errors.New("invalid expire block"), err)
			}

			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			info, err := ctx.Client.ABCIInfo(context.Background())
			if err != nil {
				return err
			}
			if expireBlock < info.Response.LastBlockHeight {
				return errors.New("expire block is earlier than current block height")
			}

			var dest string
			if cmd.Flags().Changed("dest") {
				dest, err = cmd.Flags().GetString("dest")
				if err != nil {
					panic(err)
				}
			}
			count, err := cmd.Flags().GetInt64("max_proofs")
			if err != nil {
				panic(err)
			}

			ips, err := getProvidersToUpload(ctx, dest, count)
			if err != nil {
				return err
			}

			return handlePost(ctx, cmd.Flags(), filePath, ips, expireBlock, count)
		},
	}
	cmd.Flags().String("dest", "", "upload file to a specific ip address")
	cmd.Flags().Int64("max_proofs", 3, "max proofs")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
