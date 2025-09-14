package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/o-richard/unpoly-lsp/lsp"
)

func main() {
	if err := lsp.LoadAttributes(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(lsp.SplitFunc)
	state := lsp.NewState()

	for scanner.Scan() {
		data := scanner.Bytes()
		response, shouldExit := lsp.HandleRequestMessage(data, state)
		if response != "" {
			_, _ = os.Stdout.WriteString(response)
		}
		if shouldExit {
			break
		}
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err().Error())
		os.Exit(1)
	}
}
