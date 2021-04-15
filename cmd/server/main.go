package main

import (
	"log"

	"github.com/kennethnym/based-count-bot/internal/bot"
	"github.com/kennethnym/based-count-bot/internal/server"
)

func main() {
	env, err := server.InitServer()

	if err != nil {
		log.Fatal(err)
	}

	bot.StartBot(env)
}
