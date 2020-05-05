package compilationengine

import (
	"fmt"

	"github.com/kaito2/nand2tetris/internal/tokenizer"
)

type CompilationEngine interface {
	// TODO: add functions

}

type CompilationEngineImpl struct {
	tokenizer tokenizer.Tokenizer
	current   tokenizer.Token
}

func NewCompilationEngine(inputFilename string) (CompilationEngine, error) {
	t, err := tokenizer.NewTokenizer(inputFilename)
	if err != nil {
		return CompilationEngineImpl{}, fmt.Errorf("failed to get new Tokenizer: %w", err)
	}
	currentToken := t.CurrentToken()
	t.Advance()
	return CompilationEngineImpl{
		tokenizer: t,
		current:   currentToken,
	}, nil
}

func (c *CompilationEngineImpl) advance() {
	c.current = c.tokenizer.CurrentToken()
	c.tokenizer.Advance()
}

func (c CompilationEngineImpl) currentToken() tokenizer.Token {
	return c.currentToken()
}

func (c CompilationEngineImpl) nextToken() tokenizer.Token {
	return c.tokenizer.CurrentToken()
}

// TODO: implemented

func compileClass() {

}

func compileClassVarDec() {

}

func compileSubroutine() {

}

func compileParameterList() {

}

func compileVarDec() {

}

func compileStatements() {

}

func compileDo() {

}

func compileLet() {

}

func compileWhile() {

}

func compileReturn() {

}

func compileIf() {

}

func compileExpression() {

}
