package main

import (
	"os"
	"fmt"
	"bufio"
	"pkg/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Color Calculator > ")

		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Println("Encountered an error while reading Stdin: ", err)
			continue
		}

		input := repl.Clean_Input(scanner.Text())

		if len(input) == 0 {
			fmt.Println("You must input anything")
			continue
		}

		cmd, ok := repl.Get_cmds()[input[0]]
		if !ok {
			fmt.Println("The command you've entered doesn't exist, type 'help' for list of commands and their usage")
			continue
		}

		if err := cmd.Callback(); err != nil {
			fmt.Println("Encountered an error when calling a command: ", err)
		}
	}
}
