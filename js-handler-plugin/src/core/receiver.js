function setHandler(func) {
	// read function argument from the memory
	let input = Host.inputString()

	let res = func(input)

	// copy output to host memory
	Host.outputString(JSON.stringify({
		success: res[0], 
		failure: res[1]
	}))

	return 0
}

module.exports = {setHandler}

