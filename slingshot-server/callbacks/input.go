package callbacks

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/extism/extism"
)

func Input(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {
	offset := stack[0]
	bufferInput, err := plugin.ReadBytes(offset)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	prompt := string(bufferInput)
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")

	// return the value
	offset, errReturn := plugin.WriteString(input)
	if errReturn != nil {
		fmt.Println(errReturn.Error())
		panic(errReturn)
	}
	stack[0] = offset

}
