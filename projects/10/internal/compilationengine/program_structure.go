package compilationengine

import (
	"log"
	"testing"

	"github.com/kaito2/nand2tetris/internal/types"

	"github.com/magiconair/properties/assert"
)

func (c *CompilationEngineImpl) compileClass() (xml string) {
	defer func() {
		xml = assembleMultiLineXML("class", xml)
	}()

	// expect 'class'
	log.Printf("expect: 'class', actual: '%s'\n", c.currentToken().String)
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect className (identifier)
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect '{'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect classVarDec*
	for classVarDec := c.compileClassVarDec(); classVarDec != ""; classVarDec = c.compileClassVarDec() {
		xml += classVarDec
	}

	// expect subroutineDec*
	for subroutineDec := c.compileSubroutineDec(); subroutineDec != ""; subroutineDec = c.compileSubroutineDec() {
		xml += subroutineDec
	}

	// expect '}'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	return
}

func (c *CompilationEngineImpl) compileClassVarDec() (xml string) {
	defer func() {
		if len(xml) > 0 {
			xml = assembleMultiLineXML("classVarDec", xml)
		}
	}()

	if c.currentToken().String == "static" || c.currentToken().String == "field" {
		// expect 'static' or 'field'
		log.Printf("want: %s, got: %s\n", "'static' or 'field'", c.currentToken().String)
		xml += c.compileTerminal()

		// expect type
		log.Printf("want: %s, got: %s\n", "type", c.currentToken().String)
		xml += c.compileTerminal()

		// expect varName (identifier)
		xml += c.compileTerminal()

		for c.currentToken().String == "," {
			// expect ','
			xml += c.compileTerminal()

			// expect type
			xml += c.compileTerminal()

			// expect varName
			xml += c.compileTerminal()
		}

		// expect ';'
		assert.Equal(&testing.T{}, c.currentToken().String, ";")
		xml += c.compileTerminal()
	}
	return
}

func (c *CompilationEngineImpl) compileSubroutineDec() (xml string) {
	defer func() {
		if len(xml) > 0 {
			xml = assembleMultiLineXML("subroutineDec", xml)
		}
	}()

	token := c.currentToken().String
	if token == "constructor" || token == "function" || token == "method" {
		// expect 'constructor' or 'function' or 'method'
		xml += c.compileTerminal()

		// REVIEW: type も終端として処理している
		// expect 'void' or type
		xml += c.compileTerminal()

		// expect subroutineName (identifier)
		xml += c.compileTerminal()

		// expect '('
		xml += c.compileTerminal()

		// expect parameterList
		xml += c.compileParameterList()

		// expect ')'
		xml += c.compileTerminal()

		// expect subroutineBody
		xml += c.compileSubroutineBody()
	}
	return
}

func (c *CompilationEngineImpl) compileSubroutineBody() (xml string) {
	defer func() {
		xml = assembleMultiLineXML("subroutineBody", xml)
	}()

	// expect '{'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect varDec*
	for varDec := c.compileVarDec(); varDec != ""; varDec = c.compileVarDec() {
		xml += varDec
	}

	xml += c.compileStatements()

	// expect '}'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	return
}

func (c *CompilationEngineImpl) compileParameterList() (xml string) {
	defer func() {
		xml = assembleMultiLineXML("parameterList", xml)
	}()

	// validate
	// 最初のトークンが type である => 'int' or 'char' or 'boolean' or className
	// TODO: リファクタ…
	if c.currentToken().Type != types.KEYWORD && c.currentToken().Type != types.IDENTIFIER {
		return
	}

	// expect type
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect varName (identifier)
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	for c.currentToken().String == "," {
		// expect ','
		xml += c.compileTerminal()

		// expect type
		xml += c.compileTerminal()

		// expect varName (identifier)
		xml += c.compileTerminal()
	}

	return
}

func (c *CompilationEngineImpl) compileVarDec() (xml string) {
	defer func() {
		if len(xml) > 0 {
			xml = assembleMultiLineXML("varDec", xml)
		}
	}()

	if c.currentToken().String == "var" {
		// expect 'var'
		xml += c.compileTerminal()

		// expect type
		xml += c.compileTerminal()

		// expect varName (identifier)
		xml += c.compileTerminal()

		// expect (',' varName)*
		for c.currentToken().String == "," {
			// expect ','
			xml += c.compileTerminal()

			// expect varName (identifier)
			xml += c.compileTerminal()
		}

		// expect ';'
		xml += c.compileTerminal()
	}
	return
}
