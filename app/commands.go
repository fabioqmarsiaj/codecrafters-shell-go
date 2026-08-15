package main

import "strings"

type Command string

const (
	Type     Command = "type"
	Echo     Command = "echo"
	Exit     Command = "exit"
	NotFound Command = "not found"
)

func createCommand(input string) Command {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return NotFound
	}
	switch parts[0] {
	case "type":
		return Type
	case "echo":
		return Echo
	case "exit":
		return Exit
	default:
		return NotFound
	}
}

func (c Command) isValid() bool {
	switch c {
	case Type, Echo, Exit:
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
