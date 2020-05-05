package tokenizer

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kaito2/nand2tetris/internal/types"
)

type Tokenizer interface {
	CurrentToken() Token
	Advance() bool

	// TODO: remove?
	GenerateTokenFile(outputFilename string) error
}

type TokenizerImpl struct {
	inputFilename string
	inputFile     *os.File
	scanner       *bufio.Scanner
	currentToken  Token
	lineTokens    []Token
}

func NewTokenizer(inputFilename string) (Tokenizer, error) {
	file, err := os.Open(inputFilename)
	if err != nil {
		return &TokenizerImpl{}, fmt.Errorf("failed to os.Open: %w", err)
	}
	scanner := bufio.NewScanner(file)
	return &TokenizerImpl{
		inputFilename: inputFilename,
		inputFile:     file,
		scanner:       scanner,
	}, nil
}

func (t TokenizerImpl) CurrentToken() Token {
	return t.currentToken
}

func (t *TokenizerImpl) GenerateTokenFile(outputFilename string) error {
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		return fmt.Errorf("failed to os.Create: %w", err)
	}
	defer outputFile.Close()

	outputFile.WriteString("<tokens>\n")
	defer outputFile.WriteString("</tokens>\n")

	for t.Advance() {
		token := t.currentToken
		// tokenType := types.CheckTokenType(token.TypeString())
		outputFile.WriteString(fmt.Sprintf("<%s> %s </%s>\n", token.Type, token.String, token.Type))
	}
	return nil
}

func (t *TokenizerImpl) Advance() bool {
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
		// TODO: レイヤーがおかしいので Tokenizer に含める方法を考える（普通に汚い）
		if strings.Contains(t.scanner.Text(), "/*") {
			for {
				if strings.Contains(t.scanner.Text(), "*/") {
					// 行のトークンがない場合は次の行を読みに行く
					if !t.scanner.Scan() {
						return false
					}
					break
				}
				// 行のトークンがない場合は次の行を読みに行く
				if !t.scanner.Scan() {
					return false
				}
			}
		}
		t.lineTokens = tokenizeLine(t.scanner.Text())
	}
}

func tokenizeLine(line string) (tokens []Token) {
	// skip empty line
	if len(line) == 0 {
		return nil
	}

	lineWithoutComment := removeComment(line)

	tmpString := ""
	for _, c := range lineWithoutComment {
		nextString := string(c)
		if nextString == "\n" {
			continue
		}
		if nextString == " " {
			if len(tmpString) == 0 {
				continue
			}

			// NOTE: string token の場合はスペースも追加する必要がある
			if tmpString[0] != '"' {
				tokens = append(tokens, NewToken(tmpString))
				tmpString = ""
				continue
			}
		}
		if types.IsSymbol(nextString) {
			if len(tmpString) != 0 {
				tokens = append(tokens, NewToken(tmpString))
				tmpString = ""
			}
			tokens = append(tokens, NewToken(nextString))
			continue
		}
		tmpString = tmpString + nextString
	}
	return tokens
}

func removeComment(line string) string {
	return strings.Split(line, "//")[0]
}

func stringVal(token string) string {
	return types.GetString(token)
}
