package internal

type Tokenizer struct{}

func NewTokenizer() Tokenizer {
	return Tokenizer{}
}

func (t *Tokenizer) advance() bool {
	panic("not implemented")
}

func (t Tokenizer)tokenType() TokenType{
	panic("not implemented")
}

func (t Tokenizer)keyword() KeywordType{
	panic("not implemented")
}

func (t Tokenizer) symbol() string {
	panic("not implemented")
}

func (t Tokenizer) identifier() string {
	panic("not implemented")
}

func (t Tokenizer) intVal() int {
	panic("not implemented")
}

func (t Tokenizer) stringVal() string {
	panic("not implemented")
}
