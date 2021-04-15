package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
	"github.com/kennethnym/based-count-bot/internal/server"
)

// handleMessage handles new message sent to a channel that the bot has access to.
func handleMessage(env *server.Env) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(sess *discordgo.Session, msg *discordgo.MessageCreate) {
		if msg.Author.ID == sess.State.User.ID {
			return
		}

		mc := strings.ToLower(msg.Content)

		if !strings.Contains(mc, "based") {
			return
		}

		p := env.Rds.TxPipeline()
		counts := map[string]*redis.IntCmd{}

		for _, basedUser := range msg.Mentions {
			counts[basedUser.Username] = p.Incr(env.Ctx, basedUser.ID)
		}

		_, err := p.Exec(env.Ctx)

		if err != nil {
			log.Fatal(err)
		}

		reply := ""

		for username, c := range counts {
			val := c.Val()
			if val == 1 {
				reply += fmt.Sprintf("%s is now officially based!\n", username)
			} else {
				reply += fmt.Sprintf("%s's based count is now %d\n", username, val)
			}
		}

		sess.ChannelMessageSend(msg.ChannelID, reply)
	}
}
