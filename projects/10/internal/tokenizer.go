package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/kaito2/nand2tetris/internal/types"
)

type Tokenizer struct {
	inputFilename string
	inputFile     *os.File
	reader        *bufio.Reader
	currentToken  string
}

func NewTokenizer(inputFilename string) (Tokenizer, error) {
	file, err := os.Open(inputFilename)
	if err != nil {
		return Tokenizer{}, fmt.Errorf("failed to os.Open: %w", err)
	}
	reader := bufio.NewReader(file)
	return Tokenizer{
		inputFilename: inputFilename,
		inputFile:     file,
		reader:        reader,
	}, nil
}

func (t *Tokenizer) advance() bool {
	// 初期化
	for {
		r, _, err := t.reader.ReadRune()
		if err != nil {
			// return false only when err is EOF
			return false
		}
		nextString := string(r)
		if nextString == " " {
			continue
		}
		if nextString == "\n" {
			continue
		}
		t.currentToken = nextString
		break
	}

	if types.IsSymbol(t.currentToken) {
		return true
	}

	for {
		nextRune, _, err := t.reader.ReadRune()
		if err != nil {
			// return false only when err is EOF
			return false
		}
		nextString := string(nextRune)
		if nextString == "\n" {
			continue
		}
		if nextString == " " {
			return true
		}
		if types.IsSymbol(nextString) {
			err := t.reader.UnreadRune()
			if err != nil {
				// TODO: error handling
				log.Fatalf("failed: %s", nextString)
			}
			return true
		}
		t.currentToken = t.currentToken + nextString
	}
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
