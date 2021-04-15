package server

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

// InitServer initializes server components, including environment variables.
func InitServer() (*Env, error) {
	_ = godotenv.Load()
	rds, err := initRedis()

	if err != nil {
		return nil, err
	}

	env := Env{
		Rds: rds,
		Ctx: context.Background(),
	}

	return &env, nil
}

func initRedis() (*redis.Client, error) {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))

	if err != nil {
		return nil, err
	}

	rds := redis.NewClient(opt)

	return rds, nil
}
