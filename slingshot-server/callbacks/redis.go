package callbacks

import (
	"context"
	"encoding/json"
	"log"
	"slingshot-server/clients"
	"slingshot-server/mem"
	"slingshot-server/slingshot"

	extism "github.com/extism/go-sdk"

)

func InitRedisClient(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {
	/* Expected
	{
		id:""
		uri:""
	}
	*/
	var result = slingshot.StringResult{}
	var record slingshot.RedisConfig
	// Read data from the shared memory
	err := mem.ReadJsonFromMemory(plugin, stack, &record)

	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		_, err := clients.CreateOrGetRedisClient(record)
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
		log.Println("🔴 InitRedisClient, CopyJsonToMemory:", errResult.Error())
	}

}

func RedisSet(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {
	/* Expected
	{ id: "", key: "", value: "" }
	*/
	var result = slingshot.StringResult{}
	var arguments slingshot.RedisRecord

	// Read data from the shared memory
	err := mem.ReadJsonFromMemory(plugin, stack, &arguments)

	// Construct the result
	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {

		redisCli := clients.GetRedisClient(arguments.Id)
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
		log.Println("🔴 RedisSet, CopyJsonToMemory:", errResult.Error())
	}

}

func RedisGet(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {
	/* Expected
	{ id: "", key: "" }
	*/
	var result = slingshot.StringResult{}
	var arguments slingshot.RedisRecord

	// Read data from the shared memory
	err := mem.ReadJsonFromMemory(plugin, stack, &arguments)

	// Construct the result
	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		//fmt.Println("🔵 RedisGet", arguments)
		redisCli := clients.GetRedisClient(arguments.Id)

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
		log.Println("🔴 RedisGet, CopyJsonToMemory:", errResult.Error())
	}
}

func RedisDel(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {
	/* Expected
	{ id: "", key: "" }
	*/
	var result = slingshot.StringResult{}
	var arguments slingshot.RedisRecord

	// Read data from the shared memory
	err := mem.ReadJsonFromMemory(plugin, stack, &arguments)

	// Construct the result
	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		//fmt.Println("🔵 RedisDel", arguments)
		redisCli := clients.GetRedisClient(arguments.Id)

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
		log.Println("🔴 RedisDel, CopyJsonToMemory:", errResult.Error())
	}
}

func RedisFilter(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {
	/* Expected
	{ id: "", key: "*" }
	*/
	var result = slingshot.StringResult{}
	var arguments slingshot.RedisRecord

	// Read data from the shared memory
	err := mem.ReadJsonFromMemory(plugin, stack, &arguments)

	// Construct the result
	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		//fmt.Println("🔵 RedisFilter", arguments)
		redisCli := clients.GetRedisClient(arguments.Id)

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
		log.Println("🔴 RedisFilter, CopyJsonToMemory:", errResult.Error())
	}
}

func RedisPublish(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {
	/* Expected
	{ id: "", channel: "", payload: "" }
	*/
	var result = slingshot.StringResult{}
	var arguments slingshot.RedisMessage

	// Read data from the shared memory
	err := mem.ReadJsonFromMemory(plugin, stack, &arguments)

	// Construct the result
	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		//fmt.Println("🔵 RedisGet", arguments)
		redisCli := clients.GetRedisClient(arguments.Id)

		err := redisCli.Publish(ctx, arguments.Channel, arguments.Payload).Err()

		if err != nil {
			result.Failure = err.Error()
			result.Success = ""
		} else {
			result.Failure = ""
			result.Success = "ok"
		}
	}

	// Copy the result to the memory
	errResult := mem.CopyJsonToMemory(plugin, stack, result)

	if errResult != nil {
		log.Println("🔴 RedisPublish, CopyJsonToMemory:", errResult.Error())
	}
}
