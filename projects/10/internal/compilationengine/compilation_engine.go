package compilationengine

import (
	"log"

	"github.com/kaito2/nand2tetris/internal/tokenizer"
)

type CompilationEngine interface {
	// TODO: add functions
	Compile()
}

type CompilationEngineImpl struct {
	tokenizer tokenizer.Tokenizer
	current   tokenizer.Token
	hasNext   bool
}

func NewCompilationEngine(tokenizer tokenizer.Tokenizer) CompilationEngine {
	// TODO: tokenizer側で初期化時に currentToken をセットするように修正
	// tokenizer.Advance()

	currentToken := tokenizer.CurrentToken()
	hasNext := tokenizer.Advance()
	return &CompilationEngineImpl{
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

func (c *CompilationEngineImpl) Compile() {
	xml := c.compileClass()
	log.Println(xml)
}
