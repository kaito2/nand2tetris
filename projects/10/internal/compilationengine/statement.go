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
