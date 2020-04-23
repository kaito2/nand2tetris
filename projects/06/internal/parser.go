package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CommandType int

const (
	ACommand CommandType = iota
	CCommand
	LCommand
)

const (
	DestNull = "null"
)

type Parser struct {
	scanner        *bufio.Scanner
	currentCommand string
}

func NewParser(filename string) (Parser, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Parser{}, fmt.Errorf("failed to os.Open: %w", err)
	}
	// defer f.Close()
	buf := bufio.NewScanner(f)
	return Parser{scanner: buf, currentCommand: ""}, nil
}

// p.scanner.Scan() returns hasMoreCommand (line) or not
func (p Parser) hasMoreCommand() bool {
	panic("not implemented")
}

func (p *Parser) advance() bool {
	// doesn't have more command
	if !p.scanner.Scan() {
		return false
	}
	p.currentCommand = p.scanner.Text()
	return true
}

func commandType(cmd string) CommandType {
	if rune(cmd[0]) == '@' {
		return ACommand
	}
	if rune(cmd[0]) == '(' {
		return LCommand
	}
	return CCommand
}

func symbol(cmd string) string {
	if rune(cmd[0]) == '@' {
		return cmd[1:]
	}
	// remove '(' and ')'
	return cmd[1 : len(cmd)-1]
}

func dest(cmd string) string {
	if !strings.Contains(cmd, "=") {
		return DestNull
	}
	// return 1 operand
	return strings.Split(cmd, "=")[0]
}

func comp() string {
	panic("not implemented")
}

func jump() string {
	panic("not implemented")
}
