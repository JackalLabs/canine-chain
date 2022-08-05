package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"

	"github.com/jackal-dao/canine/x/storage/types"

	"github.com/julienschmidt/httprouter"
	merkletree "github.com/wealdtech/go-merkletree"

	"github.com/spf13/cobra"
)

// This function returns the filename(to save in database) of the saved file
// or an error if it occurs
func FileUpload(w http.ResponseWriter, r *http.Request, ps httprouter.Params, cmd *cobra.Command, db *leveldb.DB, datedb *leveldb.DB) {
	// ParseMultipartForm parses a request body as multipart/form-data
	r.ParseMultipartForm(32 << 20)

	clientCtx, qerr := client.GetClientTxContext(cmd)

	file, handler, err := r.FormFile("file") // Retrieve the file from form data

	sender := r.Form.Get("sender")

	if err != nil {
		fmt.Printf("Error with form file!\n")
		return
	}

	size := handler.Size
	h := sha256.New()
	io.Copy(h, file)
	hashName := h.Sum(nil)
	file.Close()

	// This is path which we want to store the file
	direrr := os.MkdirAll(fmt.Sprintf("%s/networkfiles/%s/", clientCtx.HomeDir, fmt.Sprintf("%x", hashName)), os.ModePerm)
	if direrr != nil {
		fmt.Printf("Error directory can't be made!\n")
		return
	}

	var blocksize int64 = 1024
	var i int64 = 0
	for i = 0; i < size; i += blocksize {
		f, err := os.OpenFile(fmt.Sprintf("%s/networkfiles/%s/%d%s", clientCtx.HomeDir, fmt.Sprintf("%x", hashName), i/blocksize, ".jkl"), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf("Error can't open file!\n")
			return
		}

		file, handler, err = r.FormFile("file") // Retrieve the file from form data
		if err != nil {
			fmt.Printf("Error with form file!\n")
			return
		}

		firstx := make([]byte, blocksize)
		file.ReadAt(firstx, i)
		file.Close()
		// fmt.Printf(": %s :\n", string(firstx))
		_, writeerr := f.Write(firstx)
		if writeerr != nil {
			fmt.Printf("Error can't write to file!\n")
		}
		f.Close()
	}

	ctrerr := makeContract(cmd, []string{fmt.Sprintf("%x", hashName), sender, "0"})
	if ctrerr != nil {
		fmt.Printf("CONTRACT ERROR: %v\n", ctrerr)
		return
	}
	// cidhash := sha256.New()
	// flags := cmd.Flag("from")

	if qerr != nil {
		fmt.Printf("Client Context Error: %v\n", qerr)
		return
	}

	info, ierr := clientCtx.Keyring.Key(clientCtx.From)

	if ierr != nil {
		fmt.Printf("Inforing Error: %v\n", ierr)
		return
	}

	ko, err := keyring.MkAccKeyOutput(info)
	if err != nil {
		fmt.Printf("Inforing Error: %v\n", ierr)
		return
	}

	cidhash := sha256.New()
	io.WriteString(cidhash, ko.Address+fmt.Sprintf("%x", hashName))
	cid := cidhash.Sum(nil)

	strcid := fmt.Sprintf("%x", cid)

	err = datedb.Put([]byte(fmt.Sprintf("%x", hashName)), []byte(fmt.Sprintf("%d", 0)), nil)
	if err != nil {
		fmt.Printf("Database Error: %v\n", err)
		return
	}
	derr := db.Put([]byte(fmt.Sprintf("%x", hashName)), []byte(strcid), nil)
	if derr != nil {
		fmt.Printf("Database Error: %v\n", derr)
		return
	}

	fmt.Printf("%s %s\n", fmt.Sprintf("%x", hashName), "Added to database")

	_, cerr := db.Get([]byte(fmt.Sprintf("%x", hashName)), nil)
	if cerr != nil {
		fmt.Printf("ERROR: %s\n", cerr.Error())
	}

	type uploadResponse struct {
		CID string
		FID string
	}

	v := uploadResponse{
		CID: strcid,
		FID: fmt.Sprintf("%x", hashName),
	}
	json.NewEncoder(w).Encode(v)
}

func StartServer() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "start-miner",
		Short: "start jackal storage miner",
		Long:  `Start jackal storage miner`,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			StartFileServer(cmd)
			return nil
		},
	}

	cmd.SetOut(io.Discard)
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func indexres(cmd *cobra.Command, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	clientCtx := client.GetClientContextFromCmd(cmd)

	type indexResponse struct {
		Status  string
		Address string
	}

	v := indexResponse{
		Status:  "online",
		Address: clientCtx.GetFromAddress().String(),
	}
	json.NewEncoder(w).Encode(v)
}

func CreateMerkleForProof(cmd *cobra.Command, filename string, index int) (string, string) {

	clientCtx, qerr := client.GetClientTxContext(cmd)
	if qerr != nil {
		return "", qerr.Error()
	}
	files, _ := os.ReadDir(fmt.Sprintf("%s/networkfiles/%s/", clientCtx.HomeDir, filename))

	var data [][]byte

	var item []byte

	for i := 0; i < len(files); i += 1 {
		f, err := os.ReadFile(fmt.Sprintf("%s/networkfiles/%s/%d%s", clientCtx.HomeDir, filename, i, ".jkl"))
		if err != nil {
			fmt.Printf("Error can't open file!\n")
			return "", ""
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
		panic(err)
	}

	h := sha256.New()
	io.WriteString(h, fmt.Sprintf("%d%x", index, item))
	ditem := h.Sum(nil)

	proof, err := tree.GenerateProof(ditem)
	if err != nil {
		panic(err)
	}

	jproof, err := json.Marshal(*proof)
	if err != nil {
		panic(err)
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

	return fmt.Sprintf("%x", item), string(jproof)

}

func SubmitProof() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "submit-proof [filename] [index] [contract-id]",
		Short: "Submit merkle proof of file",
		Long:  `Submit merkle proof of file`,
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			return postProof(cmd, args)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func makeContract(cmd *cobra.Command, args []string) error {
	fmt.Printf("%s\n", args[0])

	merkleroot, filesize, fid := HashData(cmd, args[0])

	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	msg := types.NewMsgPostContract(
		clientCtx.GetFromAddress().String(),
		args[1],
		args[2],
		filesize,
		fid,
		merkleroot,
	)
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
}

func CreateTree() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "create-contract [filename] [signee] [duration]",
		Short: "Creates a contract",
		Long:  `Creates a contract`,
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return makeContract(cmd, args)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func HashData(cmd *cobra.Command, filename string) (string, string, string) {

	clientCtx, qerr := client.GetClientTxContext(cmd)
	if qerr != nil {
		fmt.Printf("%s\n", "can't get client context")
		return "", "", qerr.Error()
	}

	files, _ := os.ReadDir(fmt.Sprintf("%s/networkfiles/%s/", clientCtx.HomeDir, filename))
	fmt.Printf("%s\n", fmt.Sprintf("%s/networkfiles/%s/", clientCtx.HomeDir, filename))
	fmt.Printf("Found %d\n", len(files))
	var size = 0
	var list [][]byte

	for i := 0; i < len(files); i++ {
		dat, _ := os.ReadFile(fmt.Sprintf("%s/networkfiles/%s/%d.jkl", clientCtx.HomeDir, filename, i))

		size = size + len(dat)

		h := sha256.New()
		io.WriteString(h, fmt.Sprintf("%d%x", i, dat))
		hashName := h.Sum(nil)

		list = append(list, hashName)

	}

	t, err := merkletree.New(list)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	return hex.EncodeToString(t.Root()), fmt.Sprintf("%d", size), filename

}

func postProof(cmd *cobra.Command, args []string) error {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	dex, _ := strconv.Atoi(args[1])

	item, hashlist := CreateMerkleForProof(cmd, args[0], dex)

	fmt.Printf("%s, %s", item, hashlist)

	msg := types.NewMsgPostproof(
		clientCtx.GetFromAddress().String(),
		item,
		hashlist,
		args[2],
	)
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
}

func queryBlock(cmd *cobra.Command, cid string) (string, error) {
	clientCtx := client.GetClientContextFromCmd(cmd)

	queryClient := types.NewQueryClient(clientCtx)

	argCid := cid

	params := &types.QueryGetActiveDealsRequest{
		Cid: argCid,
	}

	res, err := queryClient.ActiveDeals(context.Background(), params)
	if err != nil {
		return "", err
	}

	return res.ActiveDeals.Blocktoprove, clientCtx.PrintProto(res)
}

func checkVerified(cmd *cobra.Command, cid string) (bool, error) {
	clientCtx := client.GetClientContextFromCmd(cmd)

	queryClient := types.NewQueryClient(clientCtx)

	argCid := cid

	params := &types.QueryGetActiveDealsRequest{
		Cid: argCid,
	}

	res, err := queryClient.ActiveDeals(context.Background(), params)
	if err != nil {
		return false, err
	}

	ver, err := strconv.ParseBool(res.ActiveDeals.Proofverified)
	if err != nil {
		return false, err
	}

	return ver, nil
}

func postProofs(cmd *cobra.Command, db *leveldb.DB, datedb *leveldb.DB) {
	clientCtx, qerr := client.GetClientTxContext(cmd)
	if qerr != nil {
		return
	}
	for {

		files, _ := os.ReadDir(fmt.Sprintf("%s/networkfiles/", clientCtx.HomeDir))

		for i := 0; i < len(files); i++ {
			nm := files[i].Name()
			fmt.Printf("filename: %s\n", nm)
			cid, cerr := db.Get([]byte(nm), nil)
			if cerr != nil {
				fmt.Printf("ERROR: %s\n", cerr.Error())
				continue
			}
			fmt.Printf("CID: %s\n", string(cid))

			ver, verr := checkVerified(cmd, string(cid))
			if verr != nil {
				fmt.Printf("ERROR: %v\n", verr)
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
				fmt.Printf("%s\n", "Skipping file as it's already verified.")
				continue
			}

			block, berr := queryBlock(cmd, string(cid))
			if berr != nil {
				fmt.Printf("ERROR: %v\n", berr)
				continue
			}

			var argss = []string{files[i].Name(), block, string(cid)}

			err := postProof(cmd, argss)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				continue
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func downfil(cmd *cobra.Command, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	clientCtx, qerr := client.GetClientTxContext(cmd)
	if qerr != nil {
		return
	}

	files, _ := os.ReadDir(fmt.Sprintf("%s/networkfiles/%s/", clientCtx.HomeDir, ps.ByName("file")))

	var data []byte

	for i := 0; i < len(files); i += 1 {
		f, err := os.ReadFile(fmt.Sprintf("%s/networkfiles/%s/%d%s", clientCtx.HomeDir, ps.ByName("file"), i, ".jkl"))
		if err != nil {
			fmt.Printf("Error can't open file!\n")
			return
		}

		data = append(data, f...)
	}

	w.Write(data)
}

func checkVersion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	type versionResponse struct {
		Version string
	}

	v := versionResponse{
		Version: "1.0.0",
	}
	json.NewEncoder(w).Encode(v)
}

func StartFileServer(cmd *cobra.Command) {
	clientCtx, qerr := client.GetClientTxContext(cmd)
	if qerr != nil {
		return
	}

	db, dberr := leveldb.OpenFile(fmt.Sprintf("%s/contracts/contractsdb", clientCtx.HomeDir), nil)
	if dberr != nil {
		fmt.Println(dberr)
	}
	datedb, dberr := leveldb.OpenFile(fmt.Sprintf("%s/contracts/datesdb", clientCtx.HomeDir), nil)
	if dberr != nil {
		fmt.Println(dberr)
	}
	router := httprouter.New()
	upfil := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		FileUpload(w, r, ps, cmd, db, datedb)
	}

	dfil := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		downfil(cmd, w, r, ps)
	}

	ires := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		indexres(cmd, w, r, ps)
	}

	router.GET("/version", checkVersion)
	router.GET("/v", checkVersion)
	router.POST("/upload", upfil)
	router.POST("/u", upfil)
	router.GET("/download/:file", dfil)
	router.GET("/d/:file", dfil)
	router.GET("/", ires)

	go postProofs(cmd, db, datedb)

	fmt.Printf("now listening!\n")
	err := http.ListenAndServe("0.0.0.0:3333", router)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("yay!")

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
