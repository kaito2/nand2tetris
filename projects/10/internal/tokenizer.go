package internal

import "github.com/kaito2/nand2tetris/internal/types"

type Tokenizer struct{}

func NewTokenizer() Tokenizer {
	return Tokenizer{}
}

func (t *Tokenizer) advance() bool {
	panic("not implemented")
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

func stringVal() string {
	panic("not implemented")
}
