package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	//nc, err := nats.Connect("nats://0.0.0.0:4222")
	nc, err := nats.Connect(nats.DefaultURL)
	

	if err != nil {
		fmt.Println(err.Error())
	}
	defer nc.Close()

	err = nc.Publish("news", []byte("Hello World"))

	if err != nil {
		fmt.Println(err.Error())
	}

}
