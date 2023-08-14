package callbacks

import (
	"context"
	"encoding/json"
	"log"
	"slingshot-server/mem"
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
	Id    string `json:"id"`
	Key   string `json:"key"`
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
	err := mem.ReadJsonFromMemory(plugin, stack, &record)

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
	errResult := mem.CopyJsonToMemory(plugin, stack, result)

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
	err := mem.ReadJsonFromMemory(plugin, stack, &arguments)

	// Construct the result
	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		//fmt.Println("🔵 RedisSet", arguments)
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
	errResult := mem.CopyJsonToMemory(plugin, stack, result)

	if errResult != nil {
		log.Println("🔴 MemorySet, CopyJsonToMemory:", err)
	}

}

func RedisGet(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {
	/* Expected
	{ id: "", key: "" }
	*/
	var result = slingshot.StringResult{}
	var arguments RedisClientArguments

	// Read data from the shared memory
	err := mem.ReadJsonFromMemory(plugin, stack, &arguments)

	// Construct the result
	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		//fmt.Println("🔵 RedisGet", arguments)
		redisCli := GetRedisClient(arguments.Id)

		value, err := redisCli.Get(ctx, string(arguments.Key)).Result()
		if err != nil {
			result.Failure = err.Error()
			result.Success = ""
		} else {
			result.Failure = ""
			result.Success = value
		}
	}

	// Copy the result to the memory
	errResult := mem.CopyJsonToMemory(plugin, stack, result)

	if errResult != nil {
		log.Println("🔴 MemorySet, CopyJsonToMemory:", err)
	}
}

func RedisDel(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {
	/* Expected
	{ id: "", key: "" }
	*/
	var result = slingshot.StringResult{}
	var arguments RedisClientArguments

	// Read data from the shared memory
	err := mem.ReadJsonFromMemory(plugin, stack, &arguments)

	// Construct the result
	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		//fmt.Println("🔵 RedisDel", arguments)
		redisCli := GetRedisClient(arguments.Id)

		_, err := redisCli.Del(ctx, string(arguments.Key)).Result()
		if err != nil {
			result.Failure = err.Error()
			result.Success = ""
		} else {
			result.Failure = ""
			result.Success = arguments.Key
		}
	}

	// Copy the result to the memory
	errResult := mem.CopyJsonToMemory(plugin, stack, result)

	if errResult != nil {
		log.Println("🔴 MemorySet, CopyJsonToMemory:", err)
	}
}

func RedisFilter(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {
	/* Expected
	{ id: "", key: "*" }
	*/
	var result = slingshot.StringResult{}
	var arguments RedisClientArguments

	// Read data from the shared memory
	err := mem.ReadJsonFromMemory(plugin, stack, &arguments)

	// Construct the result
	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		//fmt.Println("🔵 RedisFilter", arguments)
		redisCli := GetRedisClient(arguments.Id)

		keys, err := redisCli.Keys(ctx, string(arguments.Key)).Result()
		jsonArr, err := json.Marshal(keys)
		if err != nil {
			result.Failure = err.Error()
			result.Success = ""
		} else {
			result.Failure = ""
			result.Success = string(jsonArr)
		}
	}

	// Copy the result to the memory
	errResult := mem.CopyJsonToMemory(plugin, stack, result)

	if errResult != nil {
		log.Println("🔴 MemorySet, CopyJsonToMemory:", err)
	}
}
