package keeper

import (
	"log"
	"os"
)

func createLogger() (log.Logger, *os.File) {
	f, _ := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(f)

	return *log.Default(), f
}
