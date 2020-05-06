package compilationengine

func (c *CompilationEngineImpl) compileStatements() (xml string) {
	defer func() {
		xml = assembleMultiLineXML("statements", xml)
	}()

	for statement := c.compileStatement(); statement != ""; statement = c.compileStatement() {
		xml += statement
	}
	return
}

func (c *CompilationEngineImpl) compileStatement() (xml string) {
	switch c.currentToken().String {
	case "let":
		return c.compileLetStatement()
	case "if":
		return c.compileIfStatement()
	case "while":
		return c.compileWhileStatement()
	case "do":
		return c.compileDoStatement()
	case "return":
		return c.compileReturn()
	default:
		// TODO: デフォルトの挙動を改善する?
		return ""
	}
}

// TODO: rename to compileReturnStatement
func (c *CompilationEngineImpl) compileReturn() (xml string) {
	defer func() {
		xml = assembleMultiLineXML("returnStatement", xml)
	}()

	// expect 'return'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect expression?
	xml += c.compileExpression()

	// expect ';'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	return
}

func (c *CompilationEngineImpl) compileDoStatement() (xml string) {
	defer func() {
		xml = assembleMultiLineXML("doStatement", xml)
	}()

	// expect 'do'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect subroutineCall
	xml += c.compileSubroutineCall()

	// expect ';'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	return
}

// TODO:
/*
e.g.

// expect '{'
xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
c.advance()

->

xml += c.compileTerm()
*/

func (c *CompilationEngineImpl) compileWhileStatement() (xml string) {
	defer func() {
		xml = assembleMultiLineXML("whileStatement", xml)
	}()

	// expect 'while'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect '('
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	xml += c.compileExpression()

	// expect ')'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect '{'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect statements
	xml += c.compileStatements()

	// expect '}'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	return
}

func (c *CompilationEngineImpl) compileLetStatement() (xml string) {
	defer func() {
		xml = assembleMultiLineXML("letStatement", xml)
	}()

	// expect 'let'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect varName (identifier)
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	if c.currentToken().String == "[" {
		// expect '['
		xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
		c.advance()

		// expect expression
		xml += c.compileExpression()

		// expect ']'
		xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
		c.advance()
	}

	// expect '='
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect expression
	xml += c.compileExpression()

	// expect ';'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	return
}

func (c *CompilationEngineImpl) compileIfStatement() (xml string) {
	defer func() {
		xml = assembleMultiLineXML("ifStatement", xml)
	}()

	// expect 'if'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect '('
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect expression
	xml += c.compileExpression()

	// expect ')'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect '{'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect statements
	xml += c.compileStatements()

	// expect '}'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	if c.currentToken().String != "else" {
		return
	}

	// expect 'else'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect '{'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	// expect statements
	xml += c.compileStatements()

	// expect '}'
	xml += assembleTermXML(c.currentToken().TypeString(), c.currentToken().String)
	c.advance()

	return
}
