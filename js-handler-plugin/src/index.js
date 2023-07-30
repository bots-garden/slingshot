import { setHandler } from "./core/receiver"

// change this to `main`
function handle() {
	
	setHandler(param => {
		let output = "param: " + param
		let err = null

		return [output, err]
	})
}

module.exports = {handle}
