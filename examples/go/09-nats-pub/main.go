package main

import (
	"errors"

	"github.com/extism/go-pdk"
	"github.com/valyala/fastjson"
)

//export hostPrintln
func hostPrintln(offset uint64) uint64

func Println(text string) {
	memoryText := pdk.AllocateString(text)
	hostPrintln(memoryText.Offset())
}

//export hostGetEnv
func hostGetEnv(offset uint64) uint64

func GetEnv(name string) string {
	// copy the name of the environment variable to the shared memory
	variableName := pdk.AllocateString(name)
	// call the host function
	offset := hostGetEnv(variableName.Offset())

	// read the value of the result from the shared memory
	variableValue := pdk.FindMemory(offset)
	buffer := make([]byte, variableValue.Length())
	variableValue.Load(buffer)

	// cast the buffer to string and return the value
	envVarValue := string(buffer)
	return envVarValue
}

var parser = fastjson.Parser{}

//export hostInitNatsConnection
func hostInitNatsConnection(offset uint64) uint64

func InitNatsConnection(natsConnectionId string, natsUrl string) (string, error) {
	jsonStrArguments := `{"id":"` + natsConnectionId + `","url":"` + natsUrl + `"}`
	// Copy the string value to the shared memory
	arguments := pdk.AllocateString(jsonStrArguments)

	// Call the host function with Json string argument
	offset := hostInitNatsConnection(arguments.Offset())

	// Get result from the shared memory
	// The host function (hostInitNatsConnection) returns a JSON buffer:
	// {
	//   "success": "the NATS connexion id",
	//   "failure": "error message if error, else empty"
	// }
	memoryResult := pdk.FindMemory(offset)
	buffer := make([]byte, memoryResult.Length())
	memoryResult.Load(buffer)

	JSONData, err := parser.ParseBytes(buffer)
	if err != nil {
		return "", err
	}
	if len(JSONData.GetStringBytes("failure")) == 0 {
		return string(JSONData.GetStringBytes("success")), nil
	} else {
		return "", errors.New(string(JSONData.GetStringBytes("failure")))
	}

}


//export hostNatsPublish
func hostNatsPublish(offset uint64) uint64

func NatsPublish(natsServer string, subject string, data string) (string, error) {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the NATS client",
	//    "subject": "name",
	//    "data": "Bob Morane"
	// }
	//jsonStr := `{"url":"` + natsServer + `","subject":"` + subject + `","data":"` + data + `"}`
	jsonStr := `{"id":"` + natsServer + `","subject":"` + subject + `","data":"` + data + `"}`

	// Copy the string value to the shared memory
	arguments := pdk.AllocateString(jsonStr)

	// Call host function with the offset of the arguments
	offset := hostNatsPublish(arguments.Offset())

	// Get result from the shared memory
	// The host function (hostNatsPublish) returns a JSON buffer:
	// {
	//   "success": "the message",
	//   "failure": "error message if error, else empty"
	// }
	memoryResult := pdk.FindMemory(offset)
	buffResult := make([]byte, memoryResult.Length())
	memoryResult.Load(buffResult)

	JSONData, err := parser.ParseBytes(buffResult)

	if err != nil {
		return "", err
	}
	if len(JSONData.GetStringBytes("failure")) == 0 {
		return string(JSONData.GetStringBytes("success")), nil
	} else {
		return "", errors.New(string(JSONData.GetStringBytes("failure")))
	}
}

//export publish
func publish() uint64 {
	input := pdk.Input()

	natsURL := GetEnv("NATS_URL")
	Println("💜 NATS_URL: " + natsURL)
	idNatsConnection, errInit := InitNatsConnection("natsconn01", natsURL)
	if errInit != nil {
		Println("😡 " + errInit.Error())
	} else {
		Println("🙂 " + idNatsConnection)
	}

	res, err := NatsPublish("natsconn01", "news", string(input))

	if err != nil {
		Println("😡 " + err.Error())
	} else {
		Println("🙂 " + res)
	}
	return 0
}

func main() {}
