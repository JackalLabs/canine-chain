package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/jackal-dao/canine/x/storage/types"

	merkletree "github.com/wealdtech/go-merkletree"

	"github.com/spf13/cobra"
)

func CreateMerkleForProof(cmd *cobra.Command, filename string, index int) (string, string, error) {

	clientCtx, qerr := client.GetClientTxContext(cmd)
	if qerr != nil {
		return "", "", qerr
	}
	files, _ := os.ReadDir(fmt.Sprintf("%s/networkfiles/%s/", clientCtx.HomeDir, filename))

	var data [][]byte

	var item []byte

	for i := 0; i < len(files); i += 1 {
		f, err := os.ReadFile(fmt.Sprintf("%s/networkfiles/%s/%d%s", clientCtx.HomeDir, filename, i, ".jkl"))
		if err != nil {
			fmt.Printf("Error can't open file!\n")
			return "", "", err
		}

		if i == index {
			item = f
		}

		h := sha256.New()
		io.WriteString(h, fmt.Sprintf("%d%x", i, f))
		hashName := h.Sum(nil)

		data = append(data, hashName)
	}

	tree, err := merkletree.New(data)
	if err != nil {
		return "", "", err
	}

	h := sha256.New()
	io.WriteString(h, fmt.Sprintf("%d%x", index, item))
	ditem := h.Sum(nil)

	proof, err := tree.GenerateProof(ditem)
	if err != nil {
		return "", "", err
	}

	jproof, err := json.Marshal(*proof)
	if err != nil {
		return "", "", err
	}

	e := hex.EncodeToString(tree.Root())

	k, _ := hex.DecodeString(e)

	verified, err := merkletree.VerifyProof(ditem, proof, k)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if !verified {
		fmt.Printf("%s\n", "Cannot verify")
	}

	return fmt.Sprintf("%x", item), string(jproof), nil

}

func postProof(cmd *cobra.Command, cid string, block string, db *leveldb.DB) (*sdk.TxResponse, error) {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return nil, err
	}

	dex, err := strconv.Atoi(block)
	if err != nil {
		return nil, err
	}

	data, err := db.Get(makeFileKey(cid), nil)
	if err != nil {
		return nil, err
	}

	item, hashlist, err := CreateMerkleForProof(cmd, string(data), dex)
	if err != nil {
		return nil, err
	}

	msg := types.NewMsgPostproof(
		clientCtx.GetFromAddress().String(),
		item,
		hashlist,
		cid,
	)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	res, err := SendTx(clientCtx, cmd.Flags(), msg)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func postProofs(cmd *cobra.Command, db *leveldb.DB) {
	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		return
	}
	interval, err := cmd.Flags().GetUint16("interval")
	if err != nil {
		return
	}

	clientCtx, qerr := client.GetClientTxContext(cmd)
	if qerr != nil {
		return
	}
	for {

		iter := db.NewIterator(nil, nil)
		for iter.Next() {
			cid := string(iter.Key())
			value := string(iter.Value())

			if cid[:len(FILE_KEY)] != FILE_KEY {
				continue
			}

			cid = cid[len(FILE_KEY):]

			if debug {
				fmt.Printf("filename: %s\n", value)
			}

			if debug {
				fmt.Printf("CID: %s\n", cid)
			}

			ver, verr := checkVerified(cmd, cid)
			if verr != nil {
				fmt.Println("Verification error")
				fmt.Printf("ERROR: %v\n", verr)
				fmt.Println(verr.Error())

				val, err := db.Get(makeDowntimeKey(cid), nil)
				newval := 0
				if err == nil {
					newval, err = strconv.Atoi(string(val))
					if err != nil {
						continue
					}
				}
				fmt.Printf("filemissdex: %d\n", newval)
				newval += 1

				if newval > 8 {
					os.RemoveAll(fmt.Sprintf("%s/networkfiles/%s", clientCtx.HomeDir, cid))
					err = db.Delete(makeFileKey(cid), nil)
					if err != nil {
						continue
					}
					err = db.Delete(makeDowntimeKey(cid), nil)
					if err != nil {
						continue
					}
					// err = db.Delete(makeUptimeKey(cid), nil)
					// if err != nil {
					// 	continue
					// }
					continue
				}

				err = db.Put(makeDowntimeKey(cid), []byte(fmt.Sprintf("%d", newval)), nil)
				if err != nil {
					continue
				}
				continue
			}

			if ver {
				if debug {
					fmt.Printf("%s\n", "Skipping file as it's already verified.")
				}
				continue
			}

			block, berr := queryBlock(cmd, string(cid))
			if berr != nil {
				fmt.Printf("Query Error: %v\n", berr)
				continue
			}

			res, err := postProof(cmd, cid, block, db)
			if err != nil {
				fmt.Printf("Posting Error: %s\n", err.Error())
				continue
			}

			if res.Code != 0 {
				fmt.Printf("Contract Response Error: %s\n", fmt.Errorf(res.RawLog))
				continue
			}
		}
		iter.Release()
		err = iter.Error()

		time.Sleep(time.Duration(interval) * time.Second)
	}
}
