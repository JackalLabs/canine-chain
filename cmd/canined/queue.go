package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/jackalLabs/canine-chain/x/storage/types"

	ctypes "github.com/cosmos/cosmos-sdk/types"
)

func (q *UploadQueue) append(upload *Upload) {
	q.Queue = append(q.Queue, upload)
}

func (q *UploadQueue) checkStrays(clientCtx client.Context, cmd *cobra.Command, db *leveldb.DB) {
	for {
		time.Sleep(time.Second)

		qClient := types.NewQueryClient(clientCtx)

		res, err := qClient.StraysAll(cmd.Context(), &types.QueryAllStraysRequest{})
		if err != nil {
			fmt.Println(err)
			continue
			// return err
		}

		s := res.Strays

		if len(s) == 0 {
			continue
		}

		for _, stray := range s {
			if _, err := os.Stat(fmt.Sprintf("%s/networkfiles/%s", clientCtx.HomeDir, stray.Fid)); !os.IsNotExist(err) {
				continue
			}

			filesres, err := qClient.FindFile(cmd.Context(), &types.QueryFindFileRequest{Fid: stray.Fid})
			if err != nil {
				fmt.Println(err)
				continue
				// return err
			}
			fmt.Println(filesres.ProviderIps)

			var arr []string
			err = json.Unmarshal([]byte(filesres.ProviderIps), &arr)
			if err != nil {
				fmt.Println(err)
				continue
				// return err
			}

			if len(arr) == 0 {
				err = fmt.Errorf("no providers have the file we want something is wrong")
				fmt.Println(err)
				continue
				// return err
			}

			_, err = downloadFileFromURL(clientCtx, arr[0], stray.Fid, stray.Cid, db)
			if err != nil {
				fmt.Println(err)
				continue
				// return err
			}

			msg := types.NewMsgClaimStray(
				clientCtx.GetFromAddress().String(),
				stray.Cid,
			)
			if err := msg.ValidateBasic(); err != nil {
				fmt.Println(err)
				continue
				// return err
			}

			u := Upload{
				Message:  msg,
				Callback: nil,
				Err:      nil,
				Response: nil,
			}

			q.Queue = append(q.Queue, &u)

			fmt.Println(res)
		}

	}
}

func (q *UploadQueue) startListener(clientCtx client.Context, cmd *cobra.Command) error {
	for {
		time.Sleep(time.Second * 2)

		if q.Locked {
			continue
		}

		l := len(q.Queue)

		if l == 0 {
			continue
		}

		msg := make([]ctypes.Msg, 0)
		uploads := make([]*Upload, 0)
		for i := 0; i < l; i++ {
			upload := q.Queue[i]
			uploads = append(uploads, upload)
			msg = append(msg, upload.Message)
		}

		res, err := SendTx(clientCtx, cmd.Flags(), msg...)
		for _, v := range uploads {
			if err != nil {
				v.Err = err
			} else {
				if res.Code != 0 {
					v.Err = fmt.Errorf(res.RawLog)
				} else {
					v.Response = res
				}
			}
			v.Callback.Done()
		}

		q.Queue = q.Queue[l:]
	}
}
