package compilationengine

import (
	"testing"

	"github.com/kaito2/nand2tetris/internal/tokenizer"
	"github.com/kaito2/nand2tetris/internal/types"

	"github.com/stretchr/testify/assert"
)

func Test_compileReturn(t *testing.T) {
	cases := []struct {
		tokens  []tokenizer.Token
		wantXML string
	}{
		{
			[]tokenizer.Token{
				{"return", types.KEYWORD},
				{";", types.SYMBOL},
			},
			`<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
`,
		},
		{
			[]tokenizer.Token{
				{"return", types.KEYWORD},
				{"x", types.IDENTIFIER},
				{";", types.SYMBOL},
			},
			`<returnStatement>
<keyword> return </keyword>
<expression>
<term>
<identifier> x </identifier>
</term>
</expression>
<symbol> ; </symbol>
</returnStatement>
`,
		},
	}

	for _, c := range cases {
		testTokenizer := NewTestTokenizer(c.tokens)
		compilationEngine := NewCompilationEngine(&testTokenizer)
		compilationEngineImpl := compilationEngine.(CompilationEngineImpl)
		gotXML := compilationEngineImpl.compileReturn()
		assert.Equal(t, c.wantXML, gotXML)
	}
}

func Test_compileDoStatement(t *testing.T) {
	cases := []struct {
		tokens  []tokenizer.Token
		wantXML string
	}{
		{
			[]tokenizer.Token{
				{"do", types.KEYWORD},
				{"draw", types.IDENTIFIER},
				{"(", types.SYMBOL},
				{")", types.SYMBOL},
				{";", types.SYMBOL},
			},
			`<doStatement>
<keyword> do </keyword>
<identifier> draw </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
`,
		},
		{
			[]tokenizer.Token{
				{"do", types.KEYWORD},
				{"Memory", types.IDENTIFIER},
				{".", types.SYMBOL},
				{"deAlloc", types.IDENTIFIER},
				{"(", types.SYMBOL},
				{"this", types.KEYWORD},
				{")", types.SYMBOL},
				{";", types.SYMBOL},
			},
			`<doStatement>
<keyword> do </keyword>
<identifier> Memory </identifier>
<symbol> . </symbol>
<identifier> deAlloc </identifier>
<symbol> ( </symbol>
<expressionList>
<expression>
<term>
<keyword> this </keyword>
</term>
</expression>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
`,
		},
	}

	for _, c := range cases {
		testTokenizer := NewTestTokenizer(c.tokens)
		compilationEngine := NewCompilationEngine(&testTokenizer)
		compilationEngineImpl := compilationEngine.(CompilationEngineImpl)
		gotXML := compilationEngineImpl.compileDoStatement()
		assert.Equal(t, c.wantXML, gotXML)
	}
}
