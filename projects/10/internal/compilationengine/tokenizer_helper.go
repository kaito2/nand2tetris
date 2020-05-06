package compilationengine

import "github.com/kaito2/nand2tetris/internal/tokenizer"

type TestTokenizerImpl struct {
	leftTokens   []tokenizer.Token
	currentToken tokenizer.Token
}

func NewTestTokenizer(tokens []tokenizer.Token) TestTokenizerImpl {
	if len(tokens) > 0 {
		return TestTokenizerImpl{
			leftTokens:   tokens[1:],
			currentToken: tokens[0],
		}
	}
	return TestTokenizerImpl{}
}

func (t *TestTokenizerImpl) CurrentToken() tokenizer.Token {
	return t.currentToken
}

func (t *TestTokenizerImpl) Advance() bool {
	if len(t.leftTokens) == 0 {
		return false
	}
	t.currentToken, t.leftTokens = t.leftTokens[0], t.leftTokens[1:]
	return true
}

// TODO: remove this function from interface...
func (t TestTokenizerImpl) GenerateTokenFile(outputFilename string) error {
	return nil
}
