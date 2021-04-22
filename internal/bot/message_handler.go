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
		reply := ""

		if strings.Contains(mc, "based") {
			reply = increaseBaseCount(env, msg)
		} else if msg.Mentions[0].Username == botUsername {
			if strings.Contains(mc, "how based am i") {
				reply = fetchBasedCount(env, msg.Author.ID)
			}
		}

		sess.ChannelMessageSend(msg.ChannelID, reply)
	}
}
