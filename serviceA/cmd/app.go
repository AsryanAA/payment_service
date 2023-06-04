package main

import (
	"kafka-go/serviceA/controllers"
	"kafka-go/utils"
	"log"
	"net/http"
)

func main() {
	port := utils.ReadConfigFile()

	http.HandleFunc("/api/v1/fund_transfer", controllers.ApplicationRequest)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error while start server", err)
	}
}
