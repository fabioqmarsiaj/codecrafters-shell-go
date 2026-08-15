package main

import (
	"os"
	"os/exec"
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

		// 1. Verifica se o arquivo existe e não é um diretório
		if err == nil && !info.IsDir() {
			// 2. Pega as permissões atuais do arquivo
			mode := info.Mode()

			// 3. Verifica se tem permissão de execução para o Dono, Grupo ou Outros (0111)
			if mode&0111 != 0 {
				return fullPath, true
			}
		}
	}
	return "", false
}

func createCommand(input string) (Command, string, string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return NotFound, "", ""
	}

	cmdName := parts[0]
	switch cmdName {
	case "type":
		return Type, cmdName, ""
	case "echo":
		return Echo, cmdName, ""
	case "exit":
		return Exit, cmdName, ""
	default:
		// Executa o LookPath apenas AQUI
		if fullPath, err := exec.LookPath(cmdName); err == nil {
			return External, cmdName, fullPath
		}
		return NotFound, cmdName, ""
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

func (c Command) String() string {
	return string(c)
}
