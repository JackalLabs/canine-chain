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

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/cosmos/cosmos-sdk/client"
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

	if qerr != nil {
		fmt.Printf("Client Context Error: %v\n", qerr)
		v := ErrorResponse{
			Error: qerr.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(v)
		return
	}

	file, handler, err := r.FormFile("file") // Retrieve the file from form data

	sender := r.Form.Get("sender")

	if err != nil {
		fmt.Printf("Error with form file!\n")
		v := ErrorResponse{
			Error: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(v)
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
		v := ErrorResponse{
			Error: direrr.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(v)
		return
	}

	var blocksize int64 = 1024
	var i int64 = 0
	for i = 0; i < size; i += blocksize {
		f, err := os.OpenFile(fmt.Sprintf("%s/networkfiles/%s/%d%s", clientCtx.HomeDir, fmt.Sprintf("%x", hashName), i/blocksize, ".jkl"), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf("Error can't open file!\n")
			v := ErrorResponse{
				Error: err.Error(),
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(v)
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

	res, ctrerr := makeContract(cmd, []string{fmt.Sprintf("%x", hashName), sender, "0"})
	if ctrerr != nil {
		fmt.Printf("CONTRACT ERROR: %v\n", ctrerr)
		v := ErrorResponse{
			Error: ctrerr.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(v)
		return
	}

	if res.Code != 0 {
		fmt.Println(fmt.Errorf(res.RawLog))
		v := ErrorResponse{
			Error: res.RawLog,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(v)
		return
	}
	// cidhash := sha256.New()
	// flags := cmd.Flag("from")

	info, ierr := clientCtx.Keyring.Key(clientCtx.From)

	if ierr != nil {
		fmt.Printf("Inforing Error: %v\n", ierr)
		v := ErrorResponse{
			Error: ierr.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(v)
		return
	}

	ko, err := keyring.MkAccKeyOutput(info)
	if err != nil {
		fmt.Printf("Inforing Error: %v\n", ierr)
		v := ErrorResponse{
			Error: ierr.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(v)
		return
	}

	cidhash := sha256.New()
	io.WriteString(cidhash, ko.Address+fmt.Sprintf("%x", hashName))
	cid := cidhash.Sum(nil)

	strcid := fmt.Sprintf("%x", cid)

	err = datedb.Put([]byte(fmt.Sprintf("%x", hashName)), []byte(fmt.Sprintf("%d", 0)), nil)
	if err != nil {
		fmt.Printf("Database Error: %v\n", err)
		v := ErrorResponse{
			Error: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(v)
		return
	}
	derr := db.Put([]byte(fmt.Sprintf("%x", hashName)), []byte(strcid), nil)
	if derr != nil {
		fmt.Printf("Database Error: %v\n", derr)
		v := ErrorResponse{
			Error: derr.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(v)
		return
	}

	fmt.Printf("%s %s\n", fmt.Sprintf("%x", hashName), "Added to database")

	_, cerr := db.Get([]byte(fmt.Sprintf("%x", hashName)), nil)
	if cerr != nil {
		fmt.Printf("ERROR: %s\n", cerr.Error())
		v := ErrorResponse{
			Error: cerr.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(v)
		return
	}

	v := UploadResponse{
		CID: strcid,
		FID: fmt.Sprintf("%x", hashName),
	}
	json.NewEncoder(w).Encode(v)
}

func indexres(cmd *cobra.Command, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	clientCtx := client.GetClientContextFromCmd(cmd)

	v := IndexResponse{
		Status:  "online",
		Address: clientCtx.GetFromAddress().String(),
	}
	json.NewEncoder(w).Encode(v)
}

func makeContract(cmd *cobra.Command, args []string) (*sdk.TxResponse, error) {
	fmt.Printf("%s\n", args[0])

	merkleroot, filesize, fid := HashData(cmd, args[0])

	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return SendTx(clientCtx, cmd.Flags(), msg)
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

	return res.ActiveDeals.Blocktoprove, nil
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
			w.Write([]byte("cannot find file"))
			return
		}

		data = append(data, f...)
	}

	w.Write(data)
}

func checkVersion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	v := VersionResponse{
		Version: "1.0.0",
	}
	json.NewEncoder(w).Encode(v)
}

func getRoutes(cmd *cobra.Command, router *httprouter.Router) {
	dfil := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		downfil(cmd, w, r, ps)
	}

	ires := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		indexres(cmd, w, r, ps)
	}

	router.GET("/version", checkVersion)
	router.GET("/v", checkVersion)
	router.GET("/download/:file", dfil)
	router.GET("/d/:file", dfil)
	router.GET("/", ires)
}

func postRoutes(cmd *cobra.Command, router *httprouter.Router, db *leveldb.DB, datedb *leveldb.DB) {
	upfil := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		FileUpload(w, r, ps, cmd, db, datedb)
	}

	router.POST("/upload", upfil)
	router.POST("/u", upfil)
}

func StartFileServer(cmd *cobra.Command) {
	clientCtx, qerr := client.GetClientTxContext(cmd)
	if qerr != nil {
		fmt.Println(qerr)
		return
	}

	fmt.Println(cmd.Flags().GetString(flags.FlagHome))

	db, dberr := leveldb.OpenFile(fmt.Sprintf("%s/contracts/contractsdb", clientCtx.HomeDir), nil)
	if dberr != nil {
		fmt.Println(dberr)
		return
	}
	datedb, dberr := leveldb.OpenFile(fmt.Sprintf("%s/contracts/datesdb", clientCtx.HomeDir), nil)
	if dberr != nil {
		fmt.Println(dberr)
		return
	}
	router := httprouter.New()

	getRoutes(cmd, router)
	postRoutes(cmd, router, db, datedb)

	go postProofs(cmd, db, datedb)

	fmt.Printf("🌍 Storage Provider: http://0.0.0.0:3333\n")
	err := http.ListenAndServe("0.0.0.0:3333", router)
	if err != nil {
		fmt.Println(err)
		return
	}

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Storage Provider Closed\n")
		return
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
