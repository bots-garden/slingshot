function handle() {

	// read function argument from the memory
	let http_request_data = Host.inputString()

	let JSONData = JSON.parse(http_request_data)

	let text = "💛 Hello " + JSONData.body

	let response = {
		headers: {
			"Content-Type": "application/json; charset=utf-8",
			"X-Slingshot-version": "0.0.0"
		},
		textBody: text,
		statusCode: 200
	}

	// copy output to host memory
	Host.outputString(JSON.stringify(response))
}

module.exports = {handle}
