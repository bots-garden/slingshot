package main

import receiver "go-handler-plugin/core"

//export handle
func handle() {

	receiver.SetHandler(func(param []byte) ([]byte, error) {
		res := `{"message":"👋 Hello `+ string(param) + `", "number":42}`

		return []byte(res), nil
	})
}

func main() {

}
