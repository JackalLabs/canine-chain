package main

import (
	"sync"

	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const MaxFileSize = 32 << 30

type IndexResponse struct {
	Status  string
	Address string
}

type UploadResponse struct {
	CID string
	FID string
}

type ErrorResponse struct {
	Error string
}

type VersionResponse struct {
	Version string
}

type ListResponse struct {
	Files []string
}

type QueueResponse struct {
	Messages []types.Msg
}

type Message interface {
}

type Upload struct {
	Message  types.Msg
	Callback *sync.WaitGroup
	Err      error
	Response *sdk.TxResponse
}

type UploadQueue struct {
	Queue  []*Upload
	Locked bool
}

type DataBlock struct {
	Key   string
	Value string
}

type DBResponse struct {
	Data []DataBlock
}
