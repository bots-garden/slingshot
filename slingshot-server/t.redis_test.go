package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"slingshot-server/callbacks"
	"slingshot-server/clients"
	"slingshot-server/hof"
	"slingshot-server/plg"
	"slingshot-server/slingshot"
	"strings"

	//"strings"
	"testing"
)

func TestCreateRedisClient(t *testing.T) {
	fmt.Println("===[TestCreateRedisClient]===")
	record := slingshot.RedisConfig{
		Id:  "redis_cli",
		Uri: os.Getenv("REDIS_URI"),
	}

	redisCli, err := clients.CreateOrGetRedisClient(record)
	if err != nil {
		fmt.Println("🔴", "TestCreateRedisClient", err)
	} else {
		fmt.Println("🟢 TestCreateRedisClient, redisCli: ", redisCli)
	}

	if clients.GetRedisClient(record.Id) != nil {
		fmt.Println("🟢", "TestCreateRedisClient, GetRedisClient: ", clients.GetRedisClient(record.Id))

	} else {
		fmt.Println("🔴", "TestCreateRedisClient")
		t.Errorf("Redis client is null")
	}

}

func initPluginForRedis(wasmFilePath string, pluginId string) {
	ctx := context.Background()

	config := plg.GetPluginConfig("info")
	manifest := plg.GetManifest(wasmFilePath, "*", "{}", "{}")

	print_string := hof.DefineHostFunctionCallBack(
		"hostPrint",
		callbacks.Print,
	)

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

	hof.AppendHostFunction(print_string)
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
	fmt.Println("===[TestRedisInit]===")

	wasmFilePath := "../plugins/tests/use-redis/use-redis.wasm"
	wasmFunctionName := "init_redis_cli" // will return the id of the redis client

	expected := "redis-cli-wasm"

	fmt.Println("✋ REDIS_URI: ", os.Getenv("REDIS_URI"))

	initPluginForRedis(wasmFilePath, "slingshotRedisplug")

	extismPlugin, err := plg.GetPlugin("slingshotRedisplug")
	fmt.Println("✋ extismPlugin: ", extismPlugin)

	if err != nil {
		log.Println("🔴 !!! Error when getting the plugin", err)
		os.Exit(1)
	}

	_, out, err := extismPlugin.Plugin.Call(wasmFunctionName, nil)

	result := string(out)
	fmt.Println("🔵 result: ", result)

	if result != expected {
		fmt.Println("🔴", "TestRedisInit")
		t.Errorf("Result %q, Expected %q", result, expected)
	} else {
		fmt.Println("🟢", "TestRedisInit")
	}

}

func TestRedisSet(t *testing.T) {
	fmt.Println("===[TestRedisSet]===")

	//redisCli, err := clients.CreateOrGetRedisClient(record)

	wasmFilePath := "../plugins/tests/use-redis/use-redis.wasm"
	expected := "001"
	//redisClientId := "redis-cli-wasm"

	fmt.Println("✋ REDIS_URI: ", os.Getenv("REDIS_URI"))

	initPluginForRedis(wasmFilePath, "slingshotRedisplug")

	extismPlugin, err := plg.GetPlugin("slingshotRedisplug")
	fmt.Println("✋ extismPlugin: ", extismPlugin)

	if err != nil {
		log.Println("🔴 !!! Error when getting the plugin", err)
		os.Exit(1)
	}

	// First, initialize the Redis client
	_, out, err := extismPlugin.Plugin.Call("init_redis_cli", nil)
	if err != nil {
		log.Println("🔴 !!! Error when calling init_redis_cli", err)
		os.Exit(1)
	}

	result := string(out)
	fmt.Println("🔵 init_redis_cli (redis client id):", result)

	_, out, err = extismPlugin.Plugin.Call("redis_set", nil)
	result = string(out)

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
	_, out, err := plugin.Plugin.Call("init_redis_cli", nil)

	result := string(out)
	fmt.Println("🟠 init_redis_cli (redis client id):", result)

	_, out, err = plugin.Plugin.Call("redis_get", nil)
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
	_, out, err := plugin.Plugin.Call("init_redis_cli", nil)

	result := string(out)
	fmt.Println("🟠 init_redis_cli (redis client id):", result)

	_, out, err = plugin.Plugin.Call("redis_del", nil)
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
	_, out, err := plugin.Plugin.Call("init_redis_cli", nil)

	result := string(out)
	fmt.Println("🟠 init_redis_cli (redis client id):", result)

	_, out, err = plugin.Plugin.Call("redis_filter", nil)
	result = string(out)
	fmt.Println("🟠 redis_filter (keys):", result)

	if strings.Contains(result, "001") && strings.Contains(result, "002") && strings.Contains(result, "003") {
		fmt.Println("🟢", "TestRedisFilter")
	} else {
		fmt.Println("🔴", "TestRedisFilter")
		t.Errorf("Result %q, Expected %q", result, expected)
	}

}
