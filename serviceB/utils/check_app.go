package utils

import (
	"encoding/json"
	"fmt"
	"kafka-go/serviceB/connect"
	"kafka-go/src/models"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func CheckApplication(message []byte) {
	var app models.Application
	json.Unmarshal(message, &app)

	checkDuplicate(app.AppId, app.Sum, app.UserId)
}

// проверка на существование заявки с данным айди
func checkDuplicate(appId, sum, userId string) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM application WHERE id=%s AND user_id=%s", appId, userId)
	var count int
	err := connect.DB.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatal("Error while query application", err)
	} else {
		if count != 0 {
			fmt.Println("Заявка с таким id уже существует")
		} else if count > 1 {
			checkBalance(appId, sum, userId)
		} else {
			queryInsert := fmt.Sprintf("INSERT INTO application (app_id, balance, user_id) VALUES (%s, %s, %s)",
				appId, sum, userId)
			_, err := connect.DB.Exec(queryInsert)
			if err != nil {
				log.Fatal("Error while insert new record", err)
			}
			updateApplicationStatus(appId, userId)
		}
	}
}

func checkBalance(appId, sum, userId string) {
	query := fmt.Sprintf("SELECT balance FROM application WHERE user_id=%s ORDER BY app_id DESC LIMIT 1", userId)
	var balance int
	err := connect.DB.QueryRow(query).Scan(&balance)
	if err != nil {
		fmt.Println("Error while query balance", err)
	} else {
		amount, _ := strconv.Atoi(sum)
		if balance >= amount {
			queryInsert := fmt.Sprintf("INSERT INTO application (app_id, balance, user_id)"+
				"VALUES (%s, %d, %s)", appId, balance-amount, userId)
			_, err := connect.DB.Exec(queryInsert)
			if err != nil {
				log.Fatal("Error while insert new application", err)
			}
			fmt.Println("Деньги успешно списаны, статус обновиться в течение 30 секунд")
			updateApplicationStatus(appId, userId)
		} else {
			fmt.Println("Недостаточно средств на балансе у пользователя с id", userId)
		}
	}
}

func updateApplicationStatus(appId, userId string) {
	go func(id string) {
		time.Sleep(time.Second * 30)
		state := RandState()
		queryUpdateStatus := fmt.Sprintf("UPDATE application SET status='%s' WHERE app_id=%s AND user_id=%s",
			state, appId, userId)
		_, err := connect.DB.Exec(queryUpdateStatus)
		if err != nil {
			log.Fatal("Error while update status for application", err)
		}
	}(appId)
}

func RandState() string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(2)
	if num == 1 {
		return "УСПЕШНО"
	}
	return "НЕ УСПЕШНО"
}
