package main

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
