package main

import (
	"kafka-go/serviceB/connect"
	"kafka-go/serviceB/utils"
	"log"
)

func main() {
	//создание таблицы для хранения заявок
	createTable := `
		CREATE TABLE IF NOT EXISTS application (
		    id SERIAL PRIMARY KEY,
		    app_id INT NOT NULL,
		    balance INT NOT NULL,
			user_id INT NOT NULL,
			status VARCHAR(10) DEFAULT 'NIL'
		)
	`
	_, err := connect.DB.Exec(createTable)
	if err != nil {
		log.Fatal("Error while create table", err)
	}

	utils.ReadApplicationOutQueue()
}
