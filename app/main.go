package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

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

		// Agora capturamos o cmdPath retornado
		cmdType, cmdName, _ := createCommand(command)

		if !cmdType.isValid() {
			fmt.Println(cmdName + ": command not found")
			continue
		}

		switch cmdType {
		case Type:
			target := strings.TrimSpace(command[4:])

			// Avalia o alvo do type
			targetType, targetName, targetPath := createCommand(target)

			switch targetType {
			case Type, Echo, Exit:
				fmt.Printf("%s is a shell builtin\n", targetName)
			case External:
				// Usamos o targetPath que já veio pronto! Sem segundo LookPath.
				fmt.Printf("%s is %s\n", targetName, targetPath)
			default:
				fmt.Printf("%s: not found\n", targetName)
			}

		case Echo:
			fmt.Println(command[5:])

		case Exit:
			break outerLoop
		}
	}
}
