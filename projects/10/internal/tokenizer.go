package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kaito2/nand2tetris/internal/types"
)

type Tokenizer struct {
	inputFilename string
	inputFile     *os.File
	scanner       *bufio.Scanner
	currentToken  string
	lineTokens    []string
}

func NewTokenizer(inputFilename string) (Tokenizer, error) {
	file, err := os.Open(inputFilename)
	if err != nil {
		return Tokenizer{}, fmt.Errorf("failed to os.Open: %w", err)
	}
	scanner := bufio.NewScanner(file)
	return Tokenizer{
		inputFilename: inputFilename,
		inputFile:     file,
		scanner:       scanner,
	}, nil
}

func (t *Tokenizer) advance() bool {
	for {
		// 行のトークンが残っている場合
		if len(t.lineTokens) != 0 {
			// var next string
			next, left := t.lineTokens[0], t.lineTokens[1:]
			t.lineTokens = left
			t.currentToken = next
			return true
		}

		// 行のトークンがない場合は次の行を読みに行く
		if !t.scanner.Scan() {
			return false
		}
		line := t.scanner.Text()
		t.lineTokens = tokenizeLine(line)
	}
}

func tokenizeLine(line string) (tokens []string) {
	log.Println("line: ", line)
	// skip empty line
	if len(line) == 0 {
		return nil
	}

	lineWithoutComment := removeComment(line)

	tmpToken := ""
	for _, c := range lineWithoutComment {
		nextString := string(c)
		if nextString == "\n" {
			continue
		}
		if nextString == " " {
			if len(tmpToken) != 0 {
				tokens = append(tokens, tmpToken)
				tmpToken = ""
			}
			continue
		}
		if types.IsSymbol(nextString) {
			if len(tmpToken) != 0 {
				tokens = append(tokens, tmpToken)
				tmpToken = ""
			}
			tokens = append(tokens, nextString)
			continue
		}
		tmpToken = tmpToken + nextString
	}
	return tokens
}

func removeComment(line string) string {
	return strings.Split(line, "//")[0]
}

func keyword(token string) types.Keyword {
	return types.GetKeyword(token)
}

func symbol(token string) string {
	return token
}

func identifier(token string) string {
	return token
}

func intVal(token string) int32 {
	return types.GetIntegerConstant(token)
}

func stringVal(token string) string {
	return types.GetString(token)
}
