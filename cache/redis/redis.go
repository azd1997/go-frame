package redis

import (
	"github.com/azd1997/go-frame/config"
	"github.com/azd1997/go-frame/logger"
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func InitRedis(conf config.RedisConfig) error {
	client := redis.NewClient(&redis.Options{
		Addr:conf.Addr,
		Password:conf.Password,
		DB:conf.DB,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		return errors.Wrap(err, "InitRedis")
	}

	logger.Logger().Info("Redis ping: ", zap.String("ping", pong))
	return nil
}
