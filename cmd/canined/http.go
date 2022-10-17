package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/julienschmidt/httprouter"

	"github.com/spf13/cobra"
)

func indexres(cmd *cobra.Command, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		fmt.Println(err)
		return
	}

	address := clientCtx.GetFromAddress()

	v := IndexResponse{
		Status:  "online",
		Address: address.String(),
	}
	json.NewEncoder(w).Encode(v)
}

func checkVersion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	v := VersionResponse{
		Version: "1.0.0",
	}
	json.NewEncoder(w).Encode(v)
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

func listFiles(cmd *cobra.Command, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	clientCtx, qerr := client.GetClientTxContext(cmd)
	if qerr != nil {
		return
	}

	files, _ := os.ReadDir(fmt.Sprintf("%s/networkfiles/%s/", clientCtx.HomeDir, ps.ByName("file")))

	var fileNames []string = make([]string, 0)

	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}

	v := ListResponse{
		Files: fileNames,
	}

	json.NewEncoder(w).Encode(v)
}

func (q *UploadQueue) getRoutes(cmd *cobra.Command, router *httprouter.Router) {
	dfil := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		downfil(cmd, w, r, ps)
	}

	ires := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		indexres(cmd, w, r, ps)
	}

	lres := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		listFiles(cmd, w, r, ps)
	}

	queue := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		q.listqueue(cmd, w, r, ps)
	}

	router.GET("/version", checkVersion)
	router.GET("/v", checkVersion)
	router.GET("/download/:file", dfil)
	router.GET("/d/:file", dfil)
	router.GET("/list", lres)
	router.GET("/l", lres)
	router.GET("/queue", queue)
	router.GET("/q", queue)
	router.GET("/", ires)

}

func (q *UploadQueue) postRoutes(cmd *cobra.Command, router *httprouter.Router, db *leveldb.DB, datedb *leveldb.DB) {
	upfil := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		q.fileUpload(&w, r, ps, cmd, db, datedb)
	}

	router.POST("/upload", upfil)
	router.POST("/u", upfil)
}

func (q *UploadQueue) listqueue(cmd *cobra.Command, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var messages = make([]types.Msg, 0)

	for _, v := range q.Queue {
		messages = append(messages, v.Message)
	}

	v := QueueResponse{
		Messages: messages,
	}

	json.NewEncoder(w).Encode(v)
}

// This function returns the filename(to save in database) of the saved file
// or an error if it occurs
func (q *UploadQueue) fileUpload(w *http.ResponseWriter, r *http.Request, ps httprouter.Params, cmd *cobra.Command, db *leveldb.DB, datedb *leveldb.DB) {
	// ParseMultipartForm parses a request body as multipart/form-data
	r.ParseMultipartForm(MaxFileSize) // MAX file size lives here

	clientCtx, qerr := client.GetClientTxContext(cmd)

	if qerr != nil {
		fmt.Printf("Client Context Error: %v\n", qerr)
		v := ErrorResponse{
			Error: qerr.Error(),
		}
		(*w).WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(*w).Encode(v)
		return
	}

	file, handler, err := r.FormFile("file") // Retrieve the file from form data

	sender := r.Form.Get("sender")

	if err != nil {
		fmt.Printf("Error with form file!\n")
		v := ErrorResponse{
			Error: err.Error(),
		}
		(*w).WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(*w).Encode(v)
		return
	}

	err = q.saveFile(clientCtx, file, handler, sender, cmd, db, datedb, w)
	if err != nil {
		v := ErrorResponse{
			Error: err.Error(),
		}
		(*w).WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(*w).Encode(v)
	}

}
