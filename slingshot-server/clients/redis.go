package clients

import (
	"slingshot-server/slingshot"
	"sync"

	"github.com/redis/go-redis/v9"
)

var redisClients sync.Map

// GetRedisClient returns a redis client
func GetRedisClient(id string) *redis.Client {
	cli, ok := redisClients.Load(id)
	if ok {
		return cli.(*redis.Client)
	} else {
		return nil
	}
}

func CreateOrGetRedisClient(record slingshot.RedisConfig) (*redis.Client, error) {
	var redisDbCli *redis.Client

	cli, _ := redisClients.Load(record.Id)
	if cli == nil {
		addr, err := redis.ParseURL(record.Uri)
		if err != nil {
			return nil, err
		}
		redisDbCli = redis.NewClient(addr)
		redisClients.Store(record.Id, redisDbCli)
	} else {
		redisDbCli = cli.(*redis.Client)
		return redisDbCli, nil
	}
	return redisDbCli, nil
}
