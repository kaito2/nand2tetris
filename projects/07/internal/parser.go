package internal

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

type Parser struct {
	filename       string
	file           *os.File
	scanner        *bufio.Scanner
	currentCommand string
}

func NewParser(filepath string) (Parser, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return Parser{}, fmt.Errorf("failed to os.Open: %w", err)
	}
	s := bufio.NewScanner(f)
	return Parser{
		filename: filepath,
		file:     f,
		scanner:  s,
	}, nil
}

// e.g. sample-data/MemoryAccess/StaticTest1/StaticTest.vm
//      => "StaticTest"
func (p Parser) getFileBaseName() string {
	return strings.Split(path.Base(p.filename), ".")[0]
}

func (p Parser) Parse(outputFilename string) error {
	codeWriter, err := NewCodeWriter(outputFilename)
	if err != nil {
		return fmt.Errorf("failed to get NewCodeWriter: %w", err)
	}
	codeWriter.setFilename(p.getFileBaseName())
	for p.advance() {
		switch commandType(p.currentCommand) {
		case C_ARITHMETIC:
			codeWriter.writeArithmetic(p.currentCommand)
		case C_PUSH, C_POP:
			arg2Num, err := arg2(p.currentCommand)
			if err != nil {
				return fmt.Errorf("failed to get arg2: %w", err)
			}
			codeWriter.writePushPop(commandType(p.currentCommand), arg1(p.currentCommand), arg2Num)
		case C_LABEL:
			codeWriter.writeLabel(arg1(p.currentCommand))
		case C_GOTO:
			codeWriter.writeGoto(arg1(p.currentCommand))
		case C_IF:
			codeWriter.writeIf(arg1(p.currentCommand))
		case C_FUNCTION:
			arg2Num, err := arg2(p.currentCommand)
			if err != nil {
				return fmt.Errorf("failed to get arg2: %w", err)
			}
			codeWriter.writeFunction(arg1(p.currentCommand), arg2Num)
		case C_CALL:
			arg2Num, err := arg2(p.currentCommand)
			if err != nil {
				return fmt.Errorf("failed to get arg2: %w", err)
			}
			codeWriter.writeCall(arg1(p.currentCommand), arg2Num)
		case C_RETURN:
			codeWriter.writeReturn()
		default:
			// TODO: implement
			panic("not implemented !")
		}
	}
	return nil
}

func (p *Parser) advance() bool {
	for {
		if !p.scanner.Scan() {
			return false
		}
		text := removeComment(p.scanner.Text())
		if !isCommand(text) {
			continue
		}
		p.currentCommand = text
		return true
	}
}

func arg1(cmd string) string {
	operands := strings.Split(cmd, " ")
	// XXX: if len(operands) < 0 => panic
	// NOTE: if cmd == "add" or "sub" etc. => return "add" or "sub"
	if len(operands) == 1 {
		return operands[0]
	}
	return operands[1]
}

// call this function commandType is in (C_PUSH, C_POP, C_FUNCTION, C_CALL)
func arg2(cmd string) (uint16, error) {
	arg2String := strings.Split(cmd, " ")[2]
	arg2, err := strconv.ParseInt(removeWhiteSpace(arg2String), 10, 16)
	if err != nil {
		return 0, fmt.Errorf("unknown arg2 characters in cmd: %s", cmd)
	}
	return uint16(arg2), nil
}

func isCommand(text string) bool {
	if isEmptyLine(text) {
		return false
	}
	if isComment(text) {
		return false
	}
	return true
}

func isEmptyLine(text string) bool {
	whitespaceRemoved := removeWhiteSpace(text)
	return len(whitespaceRemoved) == 0
}

func isComment(text string) bool {
	whitespaceRemoved := removeWhiteSpace(text)
	return whitespaceRemoved[0:2] == "//"
}

func removeWhiteSpace(text string) string {
	halfspaceRemovedText := strings.ReplaceAll(text, " ", "")
	return strings.ReplaceAll(halfspaceRemovedText, "\t", "")
}

func removeComment(text string) string {
	return strings.Split(text, "//")[0]
}

func (p *Parser) close() {
	p.file.Close()
}
