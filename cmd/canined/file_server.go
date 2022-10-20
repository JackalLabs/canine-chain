package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/rs/cors"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/jackal-dao/canine/x/storage/types"

	"github.com/julienschmidt/httprouter"
	merkletree "github.com/wealdtech/go-merkletree"

	"github.com/spf13/cobra"
)

func writeFileToDisk(clientCtx client.Context, reader io.Reader, file io.ReaderAt, closer io.Closer, size int64) ([]byte, error) {
	h := sha256.New()
	io.Copy(h, reader)
	hashName := h.Sum(nil)

	// This is path which we want to store the file
	direrr := os.MkdirAll(fmt.Sprintf("%s/networkfiles/%s/", clientCtx.HomeDir, fmt.Sprintf("%x", hashName)), os.ModePerm)
	if direrr != nil {
		return hashName, direrr
	}

	var blocksize int64 = 1024
	var i int64 = 0
	for i = 0; i < size; i += blocksize {
		f, err := os.OpenFile(fmt.Sprintf("%s/networkfiles/%s/%d%s", clientCtx.HomeDir, fmt.Sprintf("%x", hashName), i/blocksize, ".jkl"), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return hashName, err
		}

		firstx := make([]byte, blocksize)
		read, err := file.ReadAt(firstx, i)
		fmt.Println(read)
		if err != nil && err != io.EOF {
			return hashName, err
		}
		firstx = firstx[:read]
		// fmt.Printf(": %s :\n", string(firstx))
		read, writeerr := f.Write(firstx)
		fmt.Println(read)
		if writeerr != nil {
			return hashName, err
		}
		f.Close()
	}
	if closer != nil {
		closer.Close()
	}
	return hashName, nil
}

func downloadFileFromURL(clientCtx client.Context, url string, fid string, cid string, db *leveldb.DB) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/d/%s", url, fid))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buff := bytes.NewBuffer([]byte{})
	size, err := io.Copy(buff, resp.Body)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(buff.Bytes())

	hashName, err := writeFileToDisk(clientCtx, reader, reader, nil, size)
	if err != nil {
		return hashName, err
	}

	err = saveToDatabase(hashName, cid, db)
	if err != nil {
		return hashName, err
	}

	return hashName, nil
}

func saveToDatabase(hashName []byte, strcid string, db *leveldb.DB) error {

	err := db.Put(makeDowntimeKey(strcid), []byte(fmt.Sprintf("%d", 0)), nil)
	if err != nil {
		fmt.Printf("Downtime Database Error: %v\n", err)
		return err
	}
	derr := db.Put(makeFileKey(strcid), []byte(fmt.Sprintf("%x", hashName)), nil)
	if derr != nil {
		fmt.Printf("File Database Error: %v\n", derr)
		return err
	}

	fmt.Printf("%s %s\n", fmt.Sprintf("%x", hashName), "Added to database")

	_, cerr := db.Get(makeFileKey(strcid), nil)
	if cerr != nil {
		fmt.Printf("Hash Database Error: %s\n", cerr.Error())
		return err
	}

	return nil

}

func (q *UploadQueue) saveFile(clientCtx client.Context, file multipart.File, handler *multipart.FileHeader, sender string, cmd *cobra.Command, db *leveldb.DB, w *http.ResponseWriter) error {
	size := handler.Size

	hashName, err := writeFileToDisk(clientCtx, file, file, file, size)
	if err != nil {
		fmt.Printf("Write To Disk Error: %v\n", err)
		return err
	}

	info, ierr := clientCtx.Keyring.Key(clientCtx.From)

	if ierr != nil {
		fmt.Printf("Inforing Error: %v\n", ierr)
		return ierr
	}

	ko, err := keyring.MkAccKeyOutput(info)
	if err != nil {
		fmt.Printf("Inforing Error: %v\n", ierr)
		return err
	}

	cidhash := sha256.New()

	fid := fmt.Sprintf("%x", hashName)

	io.WriteString(cidhash, fmt.Sprintf("%s%s%s", sender, ko.Address, fid))
	cid := cidhash.Sum(nil)

	strcid := fmt.Sprintf("%x", cid)

	var wg sync.WaitGroup
	wg.Add(1)

	ctrerr := q.makeContract(cmd, []string{fid, sender, "0"}, &wg)
	if ctrerr != nil {
		fmt.Printf("CONTRACT ERROR: %v\n", ctrerr)
		return ctrerr
	}
	wg.Wait()

	fmt.Printf("%x\n", hashName)

	v := UploadResponse{
		CID: strcid,
		FID: fmt.Sprintf("%x", hashName),
	}

	err = json.NewEncoder(*w).Encode(v)
	if err != nil {
		fmt.Printf("Json Encode Error: %v\n", err)
		return err
	}
	// cidhash := sha256.New()
	// flags := cmd.Flag("from")

	err = saveToDatabase(hashName, strcid, db)
	if err != nil {
		return err
	}

	return nil
}

func (q *UploadQueue) makeContract(cmd *cobra.Command, args []string, wg *sync.WaitGroup) error {

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

	u := Upload{
		Message:  msg,
		Callback: wg,
	}

	q.Queue = append(q.Queue, u)

	return nil
}

func HashData(cmd *cobra.Command, filename string) (string, string, string) {

	clientCtx, qerr := client.GetClientTxContext(cmd)
	if qerr != nil {
		fmt.Printf("%s\n", "can't get client context")
		return "", "", qerr.Error()
	}

	path := fmt.Sprintf("%s/networkfiles/%s/", clientCtx.HomeDir, filename)
	files, err := os.ReadDir(filepath.Clean(path))
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	var size = 0
	var list [][]byte

	for i := 0; i < len(files); i++ {

		path := fmt.Sprintf("%s/networkfiles/%s/%d.jkl", clientCtx.HomeDir, filename, i)

		dat, err := os.ReadFile(filepath.Clean(path))
		if err != nil {
			fmt.Printf("%v\n", err)
		}

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
	router := httprouter.New()

	q := UploadQueue{
		Queue:  make([]Upload, 0),
		Locked: false,
	}

	q.getRoutes(cmd, router, db)
	q.postRoutes(cmd, router, db)

	handler := cors.Default().Handler(router)

	go postProofs(cmd, db)
	go q.startListener(clientCtx, cmd)
	go q.checkStrays(clientCtx, cmd, db)

	fmt.Printf("🌍 Storage Provider: http://0.0.0.0:3333\n")
	err := http.ListenAndServe("0.0.0.0:3333", handler)
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
