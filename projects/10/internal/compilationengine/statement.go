package compilationengine

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
