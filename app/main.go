package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {

	reader := bufio.NewReader(os.Stdin)
	paths := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))

outerLoop:
	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)

		if command == "" {
			continue
		}

		cmdType, cmdName := createCommand(command, paths)
		if !cmdType.isValid() {
			fmt.Println(cmdName + ": command not found")
			continue
		}

		switch cmdType {
		case Type:
			// Pega o argumento do comando type (ex: "type grep" -> "grep")
			target := strings.TrimSpace(command[4:])

			// Avalia o que é o alvo do type
			targetType, _ := createCommand(target, paths)

			switch targetType {
			case Type, Echo, Exit:
				fmt.Printf("%s is a shell builtin\n", target)
			case External:
				fullPath, _ := findInPath(target, paths)
				fmt.Printf("%s is %s\n", target, fullPath)
			default:
				fmt.Printf("%s: not found\n", target)
			}
		case Echo:
			fmt.Println(command[5:])
		case Exit:
			break outerLoop
		}
	}
}
