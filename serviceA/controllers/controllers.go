package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"kafka-go/src/models"
	"kafka-go/utils"
	"log"
	"net/http"
)

func ApplicationRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "This is not GET Request endpoint")
	case "POST":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalln("Error while reading request body", err)
		}

		var app models.Application
		json.Unmarshal(body, &app)

		utils.WriteApplicationInQueue(body)

		fmt.Fprintf(w, fmt.Sprintf("Ваша заявка принята. Id вашей заявки %s", app.AppId))
	default:
		fmt.Fprintf(w, "Ooops... Что то пошло не так")
	}
}
