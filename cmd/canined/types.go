package main

import (
	"sync"

	"github.com/jackal-dao/canine/x/storage/types"
)

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

type Upload struct {
	Message  *types.MsgPostContract
	Callback *sync.WaitGroup
}

type UploadQueue struct {
	Queue  []Upload
	Locked bool
}
