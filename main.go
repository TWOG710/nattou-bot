package main

import (
	"log"

	"github.com/TWOG710/nattou-bot/util"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {

	logFile, err := util.SetLogDir()
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	conf, err := util.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := linebot.New(conf.ChannelSecret, conf.ChannelToken)
	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage(conf.Message)
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
