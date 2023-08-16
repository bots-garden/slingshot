package slingshot


//export hostPrint
func hostPrint(offset uint64) uint64

func Print(text string) {
	memoryText := CopyStringToMemory(text)
	hostPrint(memoryText.Offset())
}

//export hostLog
func hostLog(offset uint64) uint64

func Log(text string) {
	memoryText := CopyStringToMemory(text)
	hostLog(memoryText.Offset())
}
