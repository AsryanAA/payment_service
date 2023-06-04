package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kafka-go/src/models"
	"log"
	"os"
)

func ReadConfigFile() string {
	jsonFile, err := os.Open("../config.json")
	defer jsonFile.Close()
	if err != nil {
		log.Fatal("Error while open config file", err)
		return ""
	}

	options, err1 := ioutil.ReadAll(jsonFile)
	if err1 != nil {
		log.Fatal("Error while read config file", err)
		return ""
	}

	var config models.Config
	json.Unmarshal(options, &config)
	fmt.Println(config.PortDev)
	return config.PortDev
}

func ReadConnectConfigFile() models.Connect {
	jsonFile, err := os.Open("../../connect.json")
	defer jsonFile.Close()
	if err != nil {
		log.Fatal("Error while open config file", err)
	}

	options, err1 := ioutil.ReadAll(jsonFile)
	if err1 != nil {
		log.Fatal("Error while read config file", err)
	}

	var connect models.Connect
	json.Unmarshal(options, &connect)
	return connect
}
