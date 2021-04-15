package server

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Env contains environment the server is running in, including db connection.
type Env struct {
	Ctx context.Context
	Rds *redis.Client
}
