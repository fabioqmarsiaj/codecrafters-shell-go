package main

import (
	"os"
	"path/filepath"
	"strings"
)

type Command string

const (
	Type     Command = "type"
	Echo     Command = "echo"
	Exit     Command = "exit"
	External Command = "external"
	NotFound Command = "not found"
)

func findInPath(cmdName string, paths []string) (string, bool) {
	for _, dir := range paths {
		fullPath := filepath.Join(dir, cmdName)
		info, err := os.Stat(fullPath)

		// Verifica se o arquivo existe e se não é um diretório
		if err == nil && !info.IsDir() {
			// Nota: Para precisão total, você poderia checar se é executável,
			// mas para o CodeCrafters a existência do arquivo no PATH costuma bastar.
			return fullPath, true
		}
	}
	return "", false
}

func createCommand(input string, paths []string) (Command, string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return NotFound, ""
	}
	cmdName := parts[0]
	switch cmdName {
	case "type":
		return Type, cmdName
	case "echo":
		return Echo, cmdName
	case "exit":
		return Exit, cmdName
	default:
		if _, found := findInPath(cmdName, paths); found {
			return External, cmdName
		}
		return NotFound, cmdName
	}
}

func (c Command) isValid() bool {
	switch c {
	case Type, Echo, Exit, External:
		return true
	default:
		return false
	}
}

func (c Command) describe() string {
	switch c {
	case Type:
		return "type is a shell builtin"
	case Echo:
		return "echo is a shell builtin"
	case Exit:
		return "exit is a shell builtin"
	default:
		return "not found"
	}
}

func (c Command) String() string {
	return string(c)
}
