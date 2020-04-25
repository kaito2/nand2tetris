package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Parser struct {
	file    *os.File
	scanner *bufio.Scanner

	currentCommand string
}

func NewParser(filepath string) (Parser, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return Parser{}, fmt.Errorf("failed to os.Open: %w", err)
	}
	s := bufio.NewScanner(f)
	return Parser{
		file:    f,
		scanner: s,
	}, nil
}

func (p *Parser) advance() bool {
	for {
		if !p.scanner.Scan() {
			return false
		}
		text := p.scanner.Text()
		if !isCommand(text) {
			continue
		}
		p.currentCommand = text
		return true
	}
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
	fmt.Printf("whitespaceRemoved: %s(%d)\n", whitespaceRemoved, len(whitespaceRemoved))
	return len(whitespaceRemoved) == 0
}

func isComment(text string) bool {
	whitespaceRemoved := removeWhiteSpace(text)
	return whitespaceRemoved[0:2] == "//"
}

func removeWhiteSpace(text string) string {
	return strings.ReplaceAll(text, " ", "")
}

func (p *Parser) close() {
	p.file.Close()
}
