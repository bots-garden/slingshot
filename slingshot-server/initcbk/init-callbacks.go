package initcbk

import (
	"slingshot-server/callbacks"
	"slingshot-server/hof"
)

func LoadHostFunctionCallBacks() {

	print_string := hof.DefineHostFunctionCallBack(
		"hostPrint",
		callbacks.Print,
	)

	log_string := hof.DefineHostFunctionCallBack(
		"hostLog",
		callbacks.Log,
	)

	get_message := hof.DefineHostFunctionCallBack(
		"hostGetMessage",
		callbacks.GetMessage,
	)

	memory_set := hof.DefineHostFunctionCallBack(
		"hostMemorySet",
		callbacks.MemorySet,
	)

	memory_get := hof.DefineHostFunctionCallBack(
		"hostMemoryGet",
		callbacks.MemoryGet,
	)

	get_env := hof.DefineHostFunctionCallBack(
		"hostGetEnv",
		callbacks.GetEnv,
	)

	init_redis_cli := hof.DefineHostFunctionCallBack(
		"hostInitRedisClient",
		callbacks.InitNatsConnection,
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
	redis_publish := hof.DefineHostFunctionCallBack(
		"hostRedisPublish",
		callbacks.RedisPublish,
	)

	init_nats_connection := hof.DefineHostFunctionCallBack(
		"hostInitNatsConnection",
		callbacks.InitNatsConnection,
	)

	nats_publish := hof.DefineHostFunctionCallBack(
		"hostNatsPublish",
		callbacks.NatsPublish,
	)

	hof.AppendHostFunction(get_message)
	hof.AppendHostFunction(print_string)
	hof.AppendHostFunction(log_string)
	hof.AppendHostFunction(memory_set)
	hof.AppendHostFunction(memory_get)
	hof.AppendHostFunction(get_env)
	hof.AppendHostFunction(init_redis_cli)
	hof.AppendHostFunction(redis_set)
	hof.AppendHostFunction(redis_get)
	hof.AppendHostFunction(redis_del)
	hof.AppendHostFunction(redis_filter)
	hof.AppendHostFunction(redis_publish)

	hof.AppendHostFunction(init_nats_connection)

	hof.AppendHostFunction(nats_publish)
}
