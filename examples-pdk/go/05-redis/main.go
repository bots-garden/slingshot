package main

import (
	"errors"
	"strings"

	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func redisHandler(input []byte) []byte {

	redisURI := slingshot.GetEnv("REDIS_URI")
	idRedisClient, errInit := slingshot.InitRedisClient("redisDb", redisURI)
	if errInit != nil {
		slingshot.Println("😡 " + errInit.Error())
	} else {
		slingshot.Println("🙂 " + idRedisClient)
	}

	k1, errSet1 := slingshot.RedisSet("redisDb", "001", "Huey 😀")
	k2, errSet2 := slingshot.RedisSet("redisDb", "002", "Dewey 😄")
	k3, errSet3 := slingshot.RedisSet("redisDb", "003", "Louie 😆")

	allSetErrs := errors.Join(errSet1, errSet2, errSet3)
	if allSetErrs != nil {
		slingshot.Println("😡 " + allSetErrs.Error())
	} else {
		slingshot.Println("🙂 " + strings.Join([]string{k1, k2, k3}, ","))
	}

	v1, errGet1 := slingshot.RedisGet("redisDb", "001")
	v2, errGet2 := slingshot.RedisGet("redisDb", "002")
	v3, errGet3 := slingshot.RedisGet("redisDb", "003")

	allGetErrs := errors.Join(errGet1, errGet2, errGet3)
	if allGetErrs != nil {
		slingshot.Println("😡 " + allSetErrs.Error())
	} else {
		slingshot.Println("🙂 " + strings.Join([]string{v1, v2, v3}, ","))
	}

	key, errDel := slingshot.RedisDel("redisDb", "002")
	if errDel != nil {
		slingshot.Println("😡 " + errDel.Error())
	} else {
		slingshot.Println("🙂 " + key)
	}

	keys, errKeys := slingshot.RedisFilter("redisDb", "00*")
	if errKeys != nil {
		slingshot.Println("😡 " + errKeys.Error())
	} else {
		slingshot.Println("🙂 " + keys)
	}

	/* output:
	🙂 001,002,003
	🙂 Huey,Dewey,Louie
	🙂 002
	🙂 ["003","001"]
	*/

	return nil
}

//export callHandler
func callHandler() {
	slingshot.Println("👋 callHandler function")
	slingshot.ExecHandler(redisHandler)
}

func main() {}

/* with the slingshot pdk, always call `callHandler`
	
	./slingshot run \
	--wasm=./redis.wasm \
	--handler=callHandler

*/