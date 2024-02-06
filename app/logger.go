package app

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func InitLogger() {
	path := "logs/"
	fileName := "ica_host.log"

	// Create directory if it doesn't exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile(path+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Exported function for info logging
func LogInfo(v ...interface{}) {
	InfoLogger.Println(v...)
}

// Exported function for err logging
func LogError(v ...interface{}) {
	ErrorLogger.Println(v...)
}
