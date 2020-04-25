package internal

import "strings"

type CommandType uint8

const (
	C_ARITHMETIC CommandType = iota
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
)

func commandType(cmd string) CommandType {
	if isArithmeticCommand(cmd) {
		return C_ARITHMETIC
	}
	if strings.HasPrefix(cmd, "push") {
		return C_PUSH
	}
	// TODO: implement
	panic("not implemented")
}

var arithmeticCommands = []string{
	"add",
	"sub",
	"neg",
	"eq",
	"gt",
	"lt",
	"and",
	"or",
	"not",
}

func isArithmeticCommand(cmd string) bool {
	for _, arithmeticCommand := range arithmeticCommands {
		if strings.HasPrefix(cmd, arithmeticCommand) {
			return true
		}
	}
	return false
}
