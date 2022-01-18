package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
	"github.com/kennethnym/based-count-bot/internal/server"
)

func increaseBasedCount(env *server.Env, msg *discordgo.MessageCreate) string {
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
			reply += fmt.Sprintf("%s's based count is now %d. <:stonks:768043389005856788>\n", username, val)
		}
	}

	return reply
}

func decreaseBasedCount(env *server.Env, msg *discordgo.MessageCreate) string {
	p := env.Rds.TxPipeline()
	counts := map[string]*redis.IntCmd{}

	for _, user := range msg.Mentions {
		counts[user.Username] = p.Decr(env.Ctx, user.ID)
	}

	_, err := p.Exec(env.Ctx)
	if err != nil {
		log.Fatal(err)
	}

	reply := ""
	for username, c := range counts {
		val := c.Val()
		if val == 0 {
			reply += fmt.Sprintf("%s is no longer based. Ravver cringe innit.", username)
		} else {
			reply += fmt.Sprintf("%s's based count is now %d. <:notstonks:768043459143008277>\n", username, val)
		}
	}

	return reply
}
