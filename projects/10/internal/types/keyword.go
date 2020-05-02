package types

type Keyword string

var (
	CLASS       Keyword = "class"
	METHOD      Keyword = "method"
	FUNCTION    Keyword = "function"
	CONSTRUCTOR Keyword = "constructor"
	INT         Keyword = "int"
	BOOLEAN     Keyword = "boolean"
	CHAR        Keyword = "char"
	VOID        Keyword = "void"
	VAR         Keyword = "var"
	STATIC      Keyword = "static"
	FIELD       Keyword = "field"
	LET         Keyword = "let"
	DO          Keyword = "do"
	IF          Keyword = "if"
	ELSE        Keyword = "else"
	WHILE       Keyword = "while"
	RETURN      Keyword = "return"
	TRUE        Keyword = "true"
	FALSE       Keyword = "false"
	NULL        Keyword = "null"
	THIS        Keyword = "this"
)

var (
	allKeywords = []Keyword{
		CLASS,
		METHOD,
		FUNCTION,
		CONSTRUCTOR,
		INT,
		BOOLEAN,
		CHAR,
		VOID,
		VAR,
		STATIC,
		FIELD,
		LET,
		DO,
		IF,
		ELSE,
		WHILE,
		RETURN,
		TRUE,
		FALSE,
		NULL,
		THIS,
	}
)

func isKeyword(s string) bool {
	for _, keyword := range allKeywords {
		if Keyword(s) == keyword {
			return true
		}
	}
	return false
}

func GetKeyword(s string) Keyword {
	for _, keyword := range allKeywords {
		if Keyword(s) == keyword {
			return keyword
		}
	}
	// TODO: error handling
	panic("unknown keyword")
}
