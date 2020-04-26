package internal

import (
	"strings"
)

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
	// TODO: define map[string]CommandType
	if strings.HasPrefix(cmd, "push") {
		return C_PUSH
	}
	if strings.HasPrefix(cmd, "pop") {
		return C_POP
	}
	// TODO: implement
	panic("not implemented")
}

const (
	addSymbol = "add"
	subSymbol = "sub"
	negSymbol = "neg"
	eqSymbol  = "eq"
	gtSymbol  = "gt"
	ltSymbol  = "lt"
	andSymbol = "and"
	orSymbol  = "or"
	notSymbol = "not"
)

var arithmeticCommands = []string{
	addSymbol,
	subSymbol,
	negSymbol,
	eqSymbol,
	gtSymbol,
	ltSymbol,
	andSymbol,
	orSymbol,
	notSymbol,
}

var unaryFunctions = []string{
	negSymbol,
	notSymbol,
}

func symbol(cmd string) string {
	return strings.Split(cmd, " ")[0]
}

func isArithmeticCommand(cmd string) bool {
	for _, arithmeticCommand := range arithmeticCommands {
		if strings.HasPrefix(cmd, arithmeticCommand) {
			return true
		}
	}
	return false
}

func isUnaryFunction(cmd string) bool {
	symbol := strings.Split(cmd, " ")[0]
	for _, unaryFunction := range unaryFunctions {
		if symbol == unaryFunction {
			return true
		}
	}
	return false
}

// XXX: 3変数関数が存在しない前提
func isBinaryFunction(cmd string) bool {
	return !isUnaryFunction(cmd)
}

var symbolOperatorMap = map[string]string{
	addSymbol: "+",
	subSymbol: "-",
	negSymbol: "-",
	eqSymbol:  "JEQ",
	gtSymbol:  "JGT",
	ltSymbol:  "JLT",
	andSymbol: "&",
	orSymbol:  "|",
	notSymbol: "!",
}

func getOperator(cmd string) string {
	symbol := strings.Split(cmd, " ")[0]
	operator, _ := symbolOperatorMap[symbol]
	return operator
}

var segmentSymbolMap = map[string]string{
	"local":    "LCL",
	"argument": "ARG",
	"this":     "THIS",
	"that":     "THAT",
	"temp":     "5",
	"pointer":  "3",
}

func getSegmentSymbol(segment string) string {
	return segmentSymbolMap[segment]
}
