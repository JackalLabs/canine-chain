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

func postProof(cmd *cobra.Command, args []string) (*sdk.TxResponse, error) {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return nil, err
	}

	dex, _ := strconv.Atoi(args[1])

	item, hashlist, err := CreateMerkleForProof(cmd, args[0], dex)
	if err != nil {
		return nil, err
	}

	msg := types.NewMsgPostproof(
		clientCtx.GetFromAddress().String(),
		item,
		hashlist,
		args[2],
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

func postProofs(cmd *cobra.Command, db *leveldb.DB, datedb *leveldb.DB) {
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

		files, _ := os.ReadDir(fmt.Sprintf("%s/networkfiles/", clientCtx.HomeDir))

		for i := 0; i < len(files); i++ {
			nm := files[i].Name()
			if debug {
				fmt.Printf("filename: %s\n", nm)
			}
			cid, cerr := db.Get([]byte(nm), nil)
			if cerr != nil {
				fmt.Printf("Database error: %s\n", cerr.Error())
				os.RemoveAll(fmt.Sprintf("%s/networkfiles/%s", clientCtx.HomeDir, nm))
				continue
			}
			if debug {
				fmt.Printf("CID: %s\n", string(cid))
			}

			ver, verr := checkVerified(cmd, string(cid))
			if verr != nil {
				fmt.Println("Verification error")
				fmt.Printf("ERROR: %v\n", verr)
				fmt.Println(verr.Error())

				val, err := datedb.Get([]byte(nm), nil)
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
					os.RemoveAll(fmt.Sprintf("%s/networkfiles/%s", clientCtx.HomeDir, nm))
					err = db.Delete([]byte(nm), nil)
					if err != nil {
						continue
					}
					err = datedb.Delete([]byte(nm), nil)
					if err != nil {
						continue
					}
				}

				err = datedb.Put([]byte(nm), []byte(fmt.Sprintf("%d", newval)), nil)
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

			var argss = []string{files[i].Name(), block, string(cid)}

			res, err := postProof(cmd, argss)
			if err != nil {
				fmt.Printf("Posting Error: %s\n", err.Error())
				continue
			}

			if res.Code != 0 {
				fmt.Printf("Contract Response Error: %s\n", fmt.Errorf(res.RawLog))
				continue
			}
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}
