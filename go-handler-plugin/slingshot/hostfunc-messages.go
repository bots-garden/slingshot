package slingshot

//export hostGetMessage
func hostGetMessage(offset uint64) uint64

func GetMessage(key string) string {
	// Copy the string key into the shared memory
	// It's the argument to call the host function
	memoryArgument := CopyStringToMemory(key)

	// call the host function
	// memoryArgument.Offset() is the position and the length of memoryKey into the memory 
	// (2 values into only one value)
	offset := hostGetMessage(memoryArgument.Offset())

	// read the result value into the memory
	resultValue := ReadStringFromMemory(offset)

	return resultValue
}