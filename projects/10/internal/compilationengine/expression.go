package compilationengine

import (
	"fmt"
	"log"

	"github.com/kaito2/nand2tetris/internal/tokenizer"
	"github.com/kaito2/nand2tetris/internal/types"
)

// TODO: string の連結よりは []string を追記したほうがヨサソウ
// TODO: advance を内部で呼びまくる仕組みをどうにかできないか

// *** term ( op term )* ***
func (c *CompilationEngineImpl) compileExpression() (xml string) {
	defer func() {
		if len(xml) > 0 {
			xml = assembleMultiLineXML("expression", xml)
		}
	}()

	// expect term
	xml = c.compileTerm()
	for {
		if !isOpToken(c.currentToken().String) {
			break
		}
		// expect op
		xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
		c.advance()

		// expect term
		xml += c.compileTerm()
	}
	return xml
}

var opTokens = []string{
	"+",
	"-",
	"*",
	"/",
	"&",
	"|",
	"<",
	">",
	"=",
}

func isOpToken(tokenString string) bool {
	for _, opToken := range opTokens {
		if tokenString == opToken {
			return true
		}
	}
	return false
}

// *** ( expression ( ',' expression )* )? ***
func (c *CompilationEngineImpl) compileExpressionList() (xml string) {
	defer func() {
		xml = assembleMultiLineXML("expressionList", xml)
	}()

	// expect expression
	xml = c.compileExpression()
	for c.currentToken().String == "," {
		// expect ','
		xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
		c.advance()

		// expect term
		xml += c.compileExpression()
	}
	return xml
}

func (c *CompilationEngineImpl) compileTerm() (xml string) {
	defer func() {
		if len(xml) > 0 {
			xml = assembleMultiLineXML("term", xml)
		}
	}()

	if c.currentToken().Type == types.INT_CONST || c.currentToken().Type == types.STRING_CONST {
		xml = assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
		c.advance()
		return xml
	} else if isKeywordConstant(c.currentToken()) {
		xml = assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
		c.advance()
		return xml
	} else if c.currentToken().Type == types.SYMBOL {
		if c.currentToken().String == "(" {
			// expect "("
			xml = assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
			c.advance()

			// expect expression
			xml += c.compileExpression()

			// ")" is expected
			xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
			c.advance()
			return xml
		} else if c.currentToken().String == "~" || c.currentToken().String == "-" {
			// `unaryOp term` pattern
			// expect "~" or "-"
			xml = assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
			c.advance()

			// expect term
			xml += c.compileTerm()
			return xml
		}
	} else if c.currentToken().Type == types.IDENTIFIER {
		log.Println("type.IDENTIFIER is detected (next token: ", c.nextToken(), ")")

		// *** varName '['  ***
		if c.nextToken().String == "[" {
			xml = assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
			c.advance()
			// expect "["
			xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
			c.advance()
			xml += c.compileExpression()
			// expect "]"
			xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
			c.advance()
			return xml
		}
		// *** subroutineName '(' expression ')' ***
		if c.nextToken().String == "(" {
			xml = assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
			c.advance()
			// expect "("
			xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
			c.advance()
			xml += c.compileExpression()
			// expect ")"
			xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
			c.advance()
			return xml
		}
		// *** (className | varName) '.' subroutineName '(' expressionList ')'
		if c.nextToken().String == "." {
			log.Printf("pattern: `%s` is detected\n", "(className | varName) '.' subroutineName '(' expressionList ')'")

			xml = assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
			c.advance()

			// expect "."
			xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
			c.advance()

			// expect identifier (subroutineName)
			xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
			c.advance()

			// expect "("
			xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
			c.advance()

			// expect expressionList
			xml += c.compileExpressionList()

			// expect ")"
			xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
			c.advance()

			return xml
		}
		// *** varName ***
		xml = assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
		c.advance()
		return xml
	}
	// NOTE: expressionList が呼んだ際に、 expression がなければ空を返したい。
	return ""
}

func assembleTermXML(tag, content string) string {
	return fmt.Sprintf("<%s> %s </%s>\n", tag, content, tag)
}

func assembleMultiLineXML(tag, content string) string {
	return fmt.Sprintf("<%s>\n%s</%s>\n", tag, content, tag)
}

// TODO: tokenizer 側に移植?
var keywordConstants = []types.Keyword{
	types.TRUE,
	types.FALSE,
	types.NULL,
	types.THIS,
}

func isKeywordConstant(token tokenizer.Token) bool {
	if !(token.Type == types.KEYWORD) {
		return false
	}
	for _, keywordConstant := range keywordConstants {
		if token.String == string(keywordConstant) {
			return true
		}
	}
	return false
}
