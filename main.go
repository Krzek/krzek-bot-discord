package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/krzek/krzek-bot-discord/bot"
	"github.com/krzek/krzek-bot-discord/config"
)

func main() {

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	bot.Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
