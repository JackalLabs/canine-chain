package main

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

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

			res, err := SendTx(clientCtx, cmd.Flags(), upload.Message)
			if err != nil {
				return err
			}

			if res.Code != 0 {
				fmt.Println(fmt.Errorf(res.RawLog))
				return fmt.Errorf(res.RawLog)
			}

			upload.Callback.Done()

			q.Queue = q.Queue[1:]

			q.Locked = false
		}

	}
}
