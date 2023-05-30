package testutil

import (
	"log"
	"os"
	"path/filepath"
)

func CreateLogger() (log.Logger, *os.File) {
	dir := "~/jackal/canine-chain/testutil"
	logFilePath := filepath.Join(dir, "logs.log")
	f, _ := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(f)

	return *log.Default(), f
}
