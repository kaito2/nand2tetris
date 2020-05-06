package compilationengine

import (
	"testing"

	"github.com/kaito2/nand2tetris/internal/tokenizer"
	"github.com/kaito2/nand2tetris/internal/types"

	"github.com/stretchr/testify/assert"
)

func Test_compileTerm(t *testing.T) {
	tokens := []tokenizer.Token{
		{"SquareGame", types.IDENTIFIER},
		{".", types.SYMBOL},
		{"new", types.IDENTIFIER},
		{"(", types.SYMBOL},
		{")", types.SYMBOL},
	}

	wantXML := `<term>
<identifier> SquareGame </identifier>
<symbol> . </symbol>
<identifier> new </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
</term>
`

	testTokenizer := NewTestTokenizer(tokens)
	compilationEngine := NewCompilationEngine(&testTokenizer)
	compilationEngineImpl := compilationEngine.(*CompilationEngineImpl)
	gotXML := compilationEngineImpl.compileTerm()
	// t.Log("want: \n", wantXML)
	// t.Log("got: \n", gotXML)

	assert.Equal(t, wantXML, gotXML)
}

func Test_compileExpression(t *testing.T) {
	tokens := []tokenizer.Token{
		{"i", types.IDENTIFIER},
		{"*", types.SYMBOL},
		{"(", types.SYMBOL},
		{"-", types.SYMBOL},
		{"j", types.IDENTIFIER},
		{")", types.SYMBOL},
		{},
	}

	wantXML := `<expression>
<term>
<identifier> i </identifier>
</term>
<symbol> * </symbol>
<term>
<symbol> ( </symbol>
<expression>
<term>
<symbol> - </symbol>
<term>
<identifier> j </identifier>
</term>
</term>
</expression>
<symbol> ) </symbol>
</term>
</expression>
`

	testTokenizer := NewTestTokenizer(tokens)
	compilationEngine := NewCompilationEngine(&testTokenizer)
	compilationEngineImpl := compilationEngine.(*CompilationEngineImpl)
	gotXML := compilationEngineImpl.compileExpression()
	// t.Log("want: \n", wantXML)
	// t.Log("got: \n", gotXML)

	assert.Equal(t, wantXML, gotXML)
}

func Test_compileExpressionList(t *testing.T) {
	tokens := []tokenizer.Token{
		{"x", types.IDENTIFIER},
		{",", types.SYMBOL},
		{"(", types.SYMBOL},
		{"y", types.IDENTIFIER},
		{"+", types.SYMBOL},
		{"size", types.IDENTIFIER},
		{")", types.SYMBOL},
		{"-", types.SYMBOL},
		{"1", types.INT_CONST},
		{",", types.SYMBOL},
		{"x", types.IDENTIFIER},
		{"+", types.SYMBOL},
		{"size", types.IDENTIFIER},
		{",", types.SYMBOL},
		{"y", types.IDENTIFIER},
		{"+", types.SYMBOL},
		{"size", types.IDENTIFIER},
	}

	wantXML := `<expressionList>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<symbol> ( </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> ) </symbol>
</term>
<symbol> - </symbol>
<term>
<integerConstant> 1 </integerConstant>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> x </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
<symbol> , </symbol>
<expression>
<term>
<identifier> y </identifier>
</term>
<symbol> + </symbol>
<term>
<identifier> size </identifier>
</term>
</expression>
</expressionList>
`
	testTokenizer := NewTestTokenizer(tokens)
	compilationEngine := NewCompilationEngine(&testTokenizer)
	compilationEngineImpl := compilationEngine.(*CompilationEngineImpl)
	gotXML := compilationEngineImpl.compileExpressionList()
	// t.Log("want: \n", wantXML)
	// t.Log("got: \n", gotXML)

	assert.Equal(t, wantXML, gotXML)
}
