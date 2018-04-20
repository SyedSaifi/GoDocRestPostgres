package main

import (
	"GoDocRestPostgres/config"
	"GoDocRestPostgres/handler"
	"GoDocRestPostgres/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	logger := utils.CreateLogger()

	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatal("Not able to connect database.", err)
	}

	fmt.Println("Succesfully Connected to DB...")
	conf := config.NewConfiguration(db, logger)
	router := mux.NewRouter()

	handler.InitializeHandler(router, conf)

	fmt.Println("Recipes Application is starting on :: ", getPort())
	err = http.ListenAndServe(":"+getPort(), router)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Recipes Application running on port %s :: ", getPort())
	}
}

func getPort() string {
	if os.Getenv("PORT") == "" {
		return "6000"
	}
	return os.Getenv("PORT")
}
