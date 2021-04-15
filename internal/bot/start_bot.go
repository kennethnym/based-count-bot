package bot

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/kennethnym/based-count-bot/internal/server"
)

// StartBot starts the bot, and will run indefinitely until interrupted.
func StartBot(env *server.Env) {
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(handleMessage(env))
	// we only care about receiving message events
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bot is running.")

	c := make(chan os.Signal, 1)
	<-c

	discord.Close()
}
