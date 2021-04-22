package bot

import (
	"fmt"

	"github.com/kennethnym/based-count-bot/internal/server"
)

func fetchBasedCount(env *server.Env, userID string) string {
	bc := env.Rds.Get(env.Ctx, userID).Val()
	return fmt.Sprintf("Your based count is %s. Not as based as me though. Try harder.", bc)
}
