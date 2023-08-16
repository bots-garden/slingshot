function hello() {

	// read function argument from the memory
	let input = Host.inputString()

	let output = "👋 Hello " + input
	// copy output to host memory
	Host.outputString(output)
}

module.exports = {hello}
