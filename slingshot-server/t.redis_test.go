package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"slingshot-server/callbacks"
	"slingshot-server/slingshot"
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

	config := slingshot.GetPluginConfig()
	manifest := slingshot.GetManifest(wasmFilePath)

	// Add an host function
	get_env := slingshot.DefineHostFunctionCallBack(
		"hostGetEnv",
		callbacks.GetEnv,
	)
	init_redis_cli := slingshot.DefineHostFunctionCallBack(
		"hostInitRedisClient",
		callbacks.InitRedisClient,
	)

	redis_set := slingshot.DefineHostFunctionCallBack(
		"hostRedisSet",
		callbacks.RedisSet,
	)

	redis_get := slingshot.DefineHostFunctionCallBack(
		"hostRedisGet",
		callbacks.RedisGet,
	)
	redis_del := slingshot.DefineHostFunctionCallBack(
		"hostRedisDel",
		callbacks.RedisDel,
	)
	/*
	redis_filter := slingshot.DefineHostFunctionCallBack(
		"hostRedisFilter",
		callbacks.RedisFilter,
	)
	*/

	slingshot.AppendHostFunction(get_env)
	slingshot.AppendHostFunction(init_redis_cli)
	slingshot.AppendHostFunction(redis_set)
	slingshot.AppendHostFunction(redis_get)
	slingshot.AppendHostFunction(redis_del)
	//slingshot.AppendHostFunction(redis_filter)

	err := slingshot.InitializePluging(ctx, pluginId, manifest, config, slingshot.GetHostFunctions())
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

	plugin, err := slingshot.GetPlugin("slingshotRedisplug")
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

	plugin, err := slingshot.GetPlugin("slingshotRedisplug")
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

	plugin, err := slingshot.GetPlugin("slingshotRedisplug")
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

	plugin, err := slingshot.GetPlugin("slingshotRedisplug")
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
