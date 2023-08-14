package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"slingshot-server/callbacks"
	"slingshot-server/hof"
	"slingshot-server/plg"
	"strings"
	"testing"
)

func TestCreateRedisClient(t *testing.T) {

	record := callbacks.RedisClientRecord{
		Id:  "redis_cli",
		Uri: os.Getenv("REDIS_URI"),
	}

	redisCli, err := callbacks.CreateOrGetRedisClient(record)
	if err != nil {
		fmt.Println("🔴", "TestCreateRedisClient", err)
	}
	fmt.Println("🟠", redisCli)

	if callbacks.GetRedisClient(record.Id) != nil {
		fmt.Println("🟢", "TestCreateRedisClient")

	} else {
		fmt.Println("🔴", "TestCreateRedisClient")
		t.Errorf("Redis client is null")
	}

}

func initPluginForRedis(wasmFilePath string, pluginId string) {
	ctx := context.Background()

	config := plg.GetPluginConfig()
	manifest := plg.GetManifest(wasmFilePath)

	// Add an host function
	get_env := hof.DefineHostFunctionCallBack(
		"hostGetEnv",
		callbacks.GetEnv,
	)
	init_redis_cli := hof.DefineHostFunctionCallBack(
		"hostInitRedisClient",
		callbacks.InitRedisClient,
	)

	redis_set := hof.DefineHostFunctionCallBack(
		"hostRedisSet",
		callbacks.RedisSet,
	)

	redis_get := hof.DefineHostFunctionCallBack(
		"hostRedisGet",
		callbacks.RedisGet,
	)
	redis_del := hof.DefineHostFunctionCallBack(
		"hostRedisDel",
		callbacks.RedisDel,
	)

	redis_filter := hof.DefineHostFunctionCallBack(
		"hostRedisFilter",
		callbacks.RedisFilter,
	)

	hof.AppendHostFunction(get_env)
	hof.AppendHostFunction(init_redis_cli)
	hof.AppendHostFunction(redis_set)
	hof.AppendHostFunction(redis_get)
	hof.AppendHostFunction(redis_del)
	hof.AppendHostFunction(redis_filter)

	err := plg.InitializePluging(ctx, pluginId, manifest, config, hof.GetHostFunctions())
	if err != nil {
		log.Println("🔴 !!! Error when loading the plugin", err)
		os.Exit(1)
	}

}

func TestRedisInit(t *testing.T) {
	wasmFilePath := "../plugins/tests/use-redis/use-redis.wasm"
	wasmFunctionName := "init_redis_cli" // will return the id of the redis client
	expected := "redis-cli-wasm"

	fmt.Println("🟠", os.Getenv("REDIS_URI"))

	initPluginForRedis(wasmFilePath, "slingshotRedisplug")

	plugin, err := plg.GetPlugin("slingshotRedisplug")
	if err != nil {
		log.Println("🔴 !!! Error when getting the plugin", err)
		os.Exit(1)
	}
	_, out, err := plugin.Call(wasmFunctionName, nil)

	result := string(out)
	fmt.Println("🟠", result)
	if result != expected {
		fmt.Println("🔴", "TestRedisInit")
		t.Errorf("Result %q, Expected %q", result, expected)
	} else {
		fmt.Println("🟢", "TestRedisInit")
	}

}

func TestRedisSet(t *testing.T) {
	wasmFilePath := "../plugins/tests/use-redis/use-redis.wasm"
	expected := "001"
	//redisClientId := "redis-cli-wasm"

	fmt.Println("🟠", os.Getenv("REDIS_URI"))

	initPluginForRedis(wasmFilePath, "slingshotRedisplug")

	plugin, err := plg.GetPlugin("slingshotRedisplug")
	if err != nil {
		log.Println("🔴 !!! Error when getting the plugin", err)
		os.Exit(1)
	}
	// First, initialize the Redis client
	_, out, err := plugin.Call("init_redis_cli", nil)

	result := string(out)
	fmt.Println("🟠 init_redis_cli (redis client id):", result)

	_, out, err = plugin.Call("redis_set", nil)
	result = string(out)
	fmt.Println("🟠 redis_set (key):", result)

	if result != expected {
		fmt.Println("🔴", "TestRedisSet")
		t.Errorf("Result %q, Expected %q", result, expected)
	} else {
		fmt.Println("🟢", "TestRedisSet")
	}

}

func TestRedisGet(t *testing.T) {
	wasmFilePath := "../plugins/tests/use-redis/use-redis.wasm"
	expected := "zero zero one"
	//redisClientId := "redis-cli-wasm"

	fmt.Println("🟠", os.Getenv("REDIS_URI"))

	initPluginForRedis(wasmFilePath, "slingshotRedisplug")

	plugin, err := plg.GetPlugin("slingshotRedisplug")
	if err != nil {
		log.Println("🔴 !!! Error when getting the plugin", err)
		os.Exit(1)
	}
	// First, initialize the Redis client
	_, out, err := plugin.Call("init_redis_cli", nil)

	result := string(out)
	fmt.Println("🟠 init_redis_cli (redis client id):", result)

	_, out, err = plugin.Call("redis_get", nil)
	result = string(out)
	fmt.Println("🟠 redis_get (value):", result)

	if result != expected {
		fmt.Println("🔴", "TestRedisGet")
		t.Errorf("Result %q, Expected %q", result, expected)
	} else {
		fmt.Println("🟢", "TestRedisGet")
	}

}

func TestRedisDel(t *testing.T) {
	wasmFilePath := "../plugins/tests/use-redis/use-redis.wasm"
	expected := "001"
	//redisClientId := "redis-cli-wasm"

	fmt.Println("🟠", os.Getenv("REDIS_URI"))

	initPluginForRedis(wasmFilePath, "slingshotRedisplug")

	plugin, err := plg.GetPlugin("slingshotRedisplug")
	if err != nil {
		log.Println("🔴 !!! Error when getting the plugin", err)
		os.Exit(1)
	}
	// First, initialize the Redis client
	_, out, err := plugin.Call("init_redis_cli", nil)

	result := string(out)
	fmt.Println("🟠 init_redis_cli (redis client id):", result)

	_, out, err = plugin.Call("redis_del", nil)
	result = string(out)
	fmt.Println("🟠 redis_del (key):", result)

	if result != expected {
		fmt.Println("🔴", "TestRedisDel")
		t.Errorf("Result %q, Expected %q", result, expected)
	} else {
		fmt.Println("🟢", "TestRedisDel")
	}

}

func TestRedisFilter(t *testing.T) {
	wasmFilePath := "../plugins/tests/use-redis/use-redis.wasm"
	expected := `["003","001","002"]`
	//redisClientId := "redis-cli-wasm"

	fmt.Println("🟠", os.Getenv("REDIS_URI"))

	initPluginForRedis(wasmFilePath, "slingshotRedisplug")

	plugin, err := plg.GetPlugin("slingshotRedisplug")
	if err != nil {
		log.Println("🔴 !!! Error when getting the plugin", err)
		os.Exit(1)
	}
	// First, initialize the Redis client
	_, out, err := plugin.Call("init_redis_cli", nil)

	result := string(out)
	fmt.Println("🟠 init_redis_cli (redis client id):", result)

	_, out, err = plugin.Call("redis_filter", nil)
	result = string(out)
	fmt.Println("🟠 redis_filter (keys):", result)

	if strings.Contains(result, "001") && strings.Contains(result, "002") && strings.Contains(result, "003") {
		fmt.Println("🟢", "TestRedisFilter")
	} else {
		fmt.Println("🔴", "TestRedisFilter")
		t.Errorf("Result %q, Expected %q", result, expected)
	}

}
