package compilationengine

import (
	"github.com/kaito2/nand2tetris/internal/tokenizer"
)

type CompilationEngine interface {
	// TODO: add functions

}

type CompilationEngineImpl struct {
	tokenizer tokenizer.Tokenizer
	current   tokenizer.Token
	hasNext   bool
}

func NewCompilationEngine(tokenizer tokenizer.Tokenizer) CompilationEngine {
	currentToken := tokenizer.CurrentToken()
	hasNext := tokenizer.Advance()
	return CompilationEngineImpl{
		tokenizer: tokenizer,
		current:   currentToken,
		hasNext:   hasNext,
	}
}

func (c *CompilationEngineImpl) advance() bool {
	if !c.hasNext {
		c.current = tokenizer.Token{}
		return false
	}
	c.current = c.tokenizer.CurrentToken()
	c.hasNext = c.tokenizer.Advance()
	return true
}

func (c CompilationEngineImpl) currentToken() tokenizer.Token {
	return c.current
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
