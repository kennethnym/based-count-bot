package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/kennethnym/based-count-bot/internal/server"
)

// handleMessage handles new message sent to a channel that the bot has access to.
func handleMessage(env *server.Env) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(sess *discordgo.Session, msg *discordgo.MessageCreate) {
		if msg.Author.ID == sess.State.User.ID {
			return
		}

		mc := strings.TrimSpace(strings.ToLower(msg.Content))
		hasMentions := len(msg.Mentions) > 0
		reply := ""

		if hasMentions && msg.Mentions[0].Username == botUsername {
			if strings.Contains(mc, "how based am i") {
				reply = fetchBasedCount(env, msg.Author.ID)
			} else if strings.Contains(mc, "unbased") {
				reply = "ratio"
			} else if strings.Contains(mc, "based") {
				reply = "I am infinitely based, you're just stating the obvious."
			}
		} else if strings.Contains(mc, "unbased") {
			reply = decreaseBasedCount(env, msg)
		} else if strings.Contains(mc, "based") || strings.Contains(mc, "b a s e d") {
			reply = increaseBasedCount(env, msg)
		}

		sess.ChannelMessageSend(msg.ChannelID, reply)
	}
}
