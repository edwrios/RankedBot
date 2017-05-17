package main

import (
	"RankedBot/bot"
	"RankedBot/config"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	//<-make(chan struct{})
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	return
}
