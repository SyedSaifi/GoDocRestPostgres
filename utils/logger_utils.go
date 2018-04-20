package utils

import (
	"log"
	"os"
)

func CreateLogger() *log.Logger {

	file, err := os.OpenFile("GoDocRestPostgresApp.log", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		log.Fatal("Cannot open log file")
		os.Exit(1)
	}

	Logger := log.New(file,
		"LOG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	return Logger
}
