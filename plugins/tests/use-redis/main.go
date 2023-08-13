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

//export hostRedisSet
func hostRedisSet(offset uint64) uint64

//export redis_set
func redis_set() int32 {
	// Prepare json argument
	jsonStrArguments := `{"id":"redis-cli-wasm","key":"001","value":"zero zero one"}`

	// Copy the string value to the shared memory
	memoryJsonStr := pdk.AllocateString(jsonStrArguments)

	// Call the host function with Json string argument
	offset := hostRedisSet(memoryJsonStr.Offset())

	// Read the result of the function in memory
	memoryResult := pdk.FindMemory(offset)
	buffJsonResult := make([]byte, memoryResult.Length())
	memoryResult.Load(buffJsonResult)

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
	"001"
	*/

	return 0
}

//export hostRedisGet
func hostRedisGet(offset uint64) uint64

//export redis_get
func redis_get() int32 {
	// Prepare json argument
	jsonStrArguments := `{"id":"redis-cli-wasm","key":"001"}`

	// Copy the string value to the shared memory
	memoryJsonStr := pdk.AllocateString(jsonStrArguments)

	// Call the host function with Json string argument
	offset := hostRedisGet(memoryJsonStr.Offset())

	// Read the result of the function in memory
	memoryResult := pdk.FindMemory(offset)
	buffJsonResult := make([]byte, memoryResult.Length())
	memoryResult.Load(buffJsonResult)

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
	"zero zero one"
	*/

	return 0
}

//export hostRedisDel
func hostRedisDel(offset uint64) uint64

//export redis_del
func redis_del() int32 {
	// Prepare json argument
	jsonStrArguments := `{"id":"redis-cli-wasm","key":"001"}`

	// Copy the string value to the shared memory
	memoryJsonStr := pdk.AllocateString(jsonStrArguments)

	// Call the host function with Json string argument
	offset := hostRedisDel(memoryJsonStr.Offset())

	// Read the result of the function in memory
	memoryResult := pdk.FindMemory(offset)
	buffJsonResult := make([]byte, memoryResult.Length())
	memoryResult.Load(buffJsonResult)

	JSONData, err := parser.ParseBytes(buffJsonResult)
	//_, err := parser.ParseBytes(buffJsonResult)
	if err != nil {
		// Allocate space into the memory
		mem := pdk.AllocateString(err.Error())
		//mem := pdk.AllocateString("😡")
		// copy output to host memory
		pdk.OutputMemory(mem)
	} else {
		// Allocate space into the memory
		mem := pdk.AllocateString(string(JSONData.GetStringBytes("success")))
		//mem := pdk.AllocateString("🙂")
		// copy output to host memory
		pdk.OutputMemory(mem)
	}

	/* Expected result
	"001"
	*/

	return 0
}




func main() {}
