package main

import (
	"fmt"

	"ZenPal/bot"
	"ZenPal/config"
)

func main() {

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})

}
