package types

type TokenType int

const (
	KEYWORD TokenType = iota
	SYMBOL
	IDENTIFIER
	INT_CONST
	STRING_CONST
)

func CheckTokenType(token string) TokenType {
	if isKeyword(token) {
		return KEYWORD
	}
	if isSymbol(token) {
		return SYMBOL
	}
	if IsIntegerConstant(token) {
		return INT_CONST
	}
	if IsIdentifier(token) {
		return IDENTIFIER
	}
	if isString(token) {
		return STRING_CONST
	}
	panic("this token not supported")
}
