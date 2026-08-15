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
outerLoop:
	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)

		cmd := createCommand(command)
		if !cmd.isValid() {
			fmt.Fprintln(os.Stderr, command+": command not found")
			continue
		}

		switch cmd {
		case Type:
			newCmd := createCommand(command[5:])
			if !newCmd.isValid() {
				fmt.Fprintln(os.Stderr, command[5:]+": "+newCmd.describe())
				continue
			}
			fmt.Println(newCmd.describe())
		case Echo:
			fmt.Println(command[5:])
		case Exit:
			break outerLoop
		}
	}
}
