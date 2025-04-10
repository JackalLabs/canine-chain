package testutil

import (
	"log"
	"os"
)

func CreateLogger() (log.Logger, *os.File) {
	f, _ := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(f)

	return *log.Default(), f
}
