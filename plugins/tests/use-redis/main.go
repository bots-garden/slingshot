package main

import (
	"github.com/extism/go-pdk"
	"github.com/valyala/fastjson"
)

//export hostGetEnv
func hostGetEnv(offset uint64) uint64

func GetEnv(name string) string {
	varNameMemory := pdk.AllocateString(name)
	offset := hostGetEnv(varNameMemory.Offset())
	varValueMemory := pdk.FindMemory(offset)
	buffer := make([]byte, varValueMemory.Length())
	varValueMemory.Load(buffer)

	envVarValue := string(buffer)

	return envVarValue

}

var parser = fastjson.Parser{}

//export hostInitRedisClient
func hostInitRedisClient(offset uint64) uint64

//export init_redis_cli
func init_redis_cli() int32 {

	redisUri := GetEnv("REDIS_URI")

	// Prepare json argument
	jsonStrArguments := `{"id":"redis-cli-wasm","uri":"` + redisUri + `"}`

	// Copy the string value to the shared memory
	memoryJsonStr := pdk.AllocateString(jsonStrArguments)

	// Call the host function with Json string argument
	offset := hostInitRedisClient(memoryJsonStr.Offset())

	// Read the result of the function in memory
	memoryResult := pdk.FindMemory(offset)
	buffJsonResult := make([]byte, memoryResult.Length())
	memoryResult.Load(buffJsonResult)

	//jsonStrResult := string(buffResult)
	JSONData, err := parser.ParseBytes(buffJsonResult)
	if err != nil {
		// Allocate space into the memory
		mem := pdk.AllocateString(err.Error())
		// copy output to host memory
		pdk.OutputMemory(mem)
	} else {
		// Allocate space into the memory
		mem := pdk.AllocateString(string(JSONData.GetStringBytes("success")))
		// copy output to host memory
		pdk.OutputMemory(mem)
	}

	/* Expected result
		"redis-cli-wasm"
	 */

	return 0
}



func main() {}
