package types

import "log"

type Token string

const (
	KEYWORD      Token = "keyword"
	SYMBOL       Token = "symbol"
	IDENTIFIER   Token = "identifier"
	INT_CONST    Token = "integerConstant"
	STRING_CONST Token = "stringConstant"
)

func CheckTokenType(token string) Token {
	if isKeyword(token) {
		return KEYWORD
	}
	if IsSymbol(token) {
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
	log.Fatalf("this token not supported: %s", token)
	return ""
}
