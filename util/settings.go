package util

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
)

const layout = "2006-01-02"

type ConfigFile struct {
	ChannelSecret string `json:"channelSecret"`
	ChannelToken  string `json:"channelToken"`
	Message       string `json:"message"`
}

func LoadConfig() (*ConfigFile, error) {

	configFileName, err := os.Open("./json/config.json")
	if err != nil {
		log.Println("Error : LoadConfig : os.Open : ", err)
		return nil, err
	}
	defer configFileName.Close()

	raw, err := io.ReadAll(configFileName)
	if err != nil {
		log.Println("Error : LoadConfig : io.ReadAll : ", err)
		return nil, err
	}

	var configFile ConfigFile
	json.Unmarshal(raw, &configFile)

	return &configFile, nil
}

func SetLogDir() (*os.File, error) {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	logDir := currentDir + "/log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.Mkdir(logDir, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	now := time.Now()
	logfileName := logDir + "/log_" + now.Format(layout) + ".txt"

	logFile, err := os.OpenFile(logfileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logFile)

	return logFile, nil
}
