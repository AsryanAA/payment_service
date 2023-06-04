package connect

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"kafka-go/src/models"
	"kafka-go/utils"
	"log"
)

func connectDB() *sql.DB {
	var connect models.Connect = utils.ReadConnectConfigFile()
	// connection string
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		connect.HostDB, connect.PortDB, connect.User, connect.Password, connect.DataBase)

	// open database
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		log.Fatal("Error while connect DB", err)
	}

	return db
}

var DB = connectDB()
