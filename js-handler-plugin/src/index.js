import { callHandler } from "./core/receiver"

// change this to `main`
function handle() {
	
	console.log("HELLO")

	callHandler(param => {
		let output = "param: " + param
		let err = null

		return [output, err]
	})
}

module.exports = {handle}
