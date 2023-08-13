package callbacks

import (
	"context"
	"log"
	"slingshot-server/slingshot"
	"sync"

	"github.com/redis/go-redis/v9"

	"github.com/extism/extism"
)

type RedisClientRecord struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

type RedisClientArguments struct {
	Id  string `json:"id"`
	Key string `json:"key"`
	Value string `json:"value"`
}

var redisClients sync.Map

func GetRedisClient(id string) *redis.Client {
	cli, ok := redisClients.Load(id)
	if ok {
		return cli.(*redis.Client)
	} else {
		return nil
	}
}

func CreateOrGetRedisClient(record RedisClientRecord) (*redis.Client, error) {
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


func InitRedisClient(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {
	/* Expected
	{
		id:""
		uri:""
	}
	*/
	var result = slingshot.StringResult{}
	var record RedisClientRecord
	// Read data from the shared memory
	err := slingshot.ReadJsonFromMemory(plugin, stack, &record)

	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		_, err := CreateOrGetRedisClient(record)
		if err != nil {
			result.Failure = err.Error()
			result.Success = ""
		}
		result.Failure = ""
		result.Success = record.Id
	}

	// Copy the result to the memory
	errResult := slingshot.CopyJsonToMemory(plugin, stack, result)

	if errResult != nil {
		log.Println("🔴 MemorySet, CopyJsonToMemory:", err)
	}

}

func RedisSet(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {
	/* Expected
	{ id: "", key: "", value: "" }
	*/
	var result = slingshot.StringResult{}
	var arguments RedisClientArguments

	// Read data from the shared memory
	err := slingshot.ReadJsonFromMemory(plugin, stack, &arguments)

	// Construct the result
	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		redisCli := GetRedisClient(arguments.Id)

		err = redisCli.Set(ctx, string(arguments.Key), string(arguments.Value), 0).Err() 
		if err != nil {
			result.Failure = err.Error()
			result.Success = ""
		} else {
			result.Failure = ""
			result.Success = arguments.Key
		}
	}

	// Copy the result to the memory
	errResult := slingshot.CopyJsonToMemory(plugin, stack, result)

	if errResult != nil {
		log.Println("🔴 MemorySet, CopyJsonToMemory:", err)
	}

}
