package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
	"github.com/kennethnym/based-count-bot/internal/server"
)

func increaseBaseCount(env *server.Env, msg *discordgo.MessageCreate) string {
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

	return reply
}
