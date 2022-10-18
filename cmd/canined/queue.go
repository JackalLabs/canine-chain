package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/jackal-dao/canine/x/storage/types"
	stypes "github.com/jackal-dao/canine/x/storage/types"
)

func (q *UploadQueue) checkStrays(clientCtx client.Context, cmd *cobra.Command, db *leveldb.DB, datedb *leveldb.DB) {
	for {
		time.Sleep(time.Second)

		qClient := stypes.NewQueryClient(clientCtx)

		res, err := qClient.StraysAll(cmd.Context(), &stypes.QueryAllStraysRequest{})
		if err != nil {
			fmt.Println(err)
			continue
			// return err
		}

		s := res.Strays

		if len(s) == 0 {
			continue
		}

		stray := s[0]

		filesres, err := qClient.FindFile(cmd.Context(), &stypes.QueryFindFileRequest{Fid: stray.Fid})
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

		_, err = downloadFileFromURL(clientCtx, arr[0], stray.Fid, stray.Cid, db, datedb)
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
		}

		q.Queue = append(q.Queue, u)

		fmt.Println(res)

	}
}

func (q *UploadQueue) startListener(clientCtx client.Context, cmd *cobra.Command) error {
	for {
		time.Sleep(time.Second)

		if q.Locked {
			continue
		}

		if len(q.Queue) > 0 {
			fmt.Println(q.Queue)

			q.Locked = true
			upload := q.Queue[0]
			q.Queue = q.Queue[1:]

			res, err := SendTx(clientCtx, cmd.Flags(), upload.Message)
			if err != nil {
				fmt.Println(err)
			} else {
				if res.Code != 0 {
					fmt.Println(fmt.Errorf(res.RawLog))
				}
			}

			if upload.Callback != nil {
				upload.Callback.Done()
			}

			q.Locked = false
		}

	}
}
