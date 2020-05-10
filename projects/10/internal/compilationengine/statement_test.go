package compilationengine

import (
	"testing"

	"github.com/sirupsen/logrus"

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
		logrus.Debug(testTokenizer)
		compilationEngine := NewCompilationEngine(&testTokenizer)
		compilationEngineImpl := compilationEngine.(*CompilationEngineImpl)
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
		compilationEngineImpl := compilationEngine.(*CompilationEngineImpl)
		gotXML := compilationEngineImpl.compileDoStatement()
		assert.Equal(t, c.wantXML, gotXML)
	}
}

func Test_compileWhileStatement(t *testing.T) {
	cases := []struct {
		tokens  []tokenizer.Token
		wantXML string
	}{
		{
			[]tokenizer.Token{
				{"while", types.KEYWORD},
				{"(", types.SYMBOL},
				{"key", types.IDENTIFIER},
				{")", types.SYMBOL},
				{"{", types.SYMBOL},
				{"let", types.KEYWORD},
				{"key", types.IDENTIFIER},
				{"=", types.SYMBOL},
				{"key", types.IDENTIFIER},
				{";", types.SYMBOL},
				{"do", types.KEYWORD},
				{"moveSquare", types.IDENTIFIER},
				{"(", types.SYMBOL},
				{")", types.SYMBOL},
				{";", types.SYMBOL},
				{"}", types.SYMBOL},
			},
			`<whileStatement>
<keyword> while </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> key </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</whileStatement>
`,
		},
	}

	for _, c := range cases {
		testTokenizer := NewTestTokenizer(c.tokens)
		compilationEngine := NewCompilationEngine(&testTokenizer)
		compilationEngineImpl := compilationEngine.(*CompilationEngineImpl)
		gotXML := compilationEngineImpl.compileWhileStatement()
		assert.Equal(t, c.wantXML, gotXML)
	}
}

func Test_compileStatements(t *testing.T) {
	cases := []struct {
		tokens  []tokenizer.Token
		wantXML string
	}{
		{
			[]tokenizer.Token{
				{"let", types.KEYWORD},
				{"key", types.IDENTIFIER},
				{"=", types.SYMBOL},
				{"key", types.IDENTIFIER},
				{";", types.SYMBOL},
				{"do", types.KEYWORD},
				{"moveSquare", types.IDENTIFIER},
				{"(", types.SYMBOL},
				{")", types.SYMBOL},
				{";", types.SYMBOL},
			},
			`<statements>
<letStatement>
<keyword> let </keyword>
<identifier> key </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
`,
		},
		{
			[]tokenizer.Token{
				{"if", types.KEYWORD},
				{"(", types.SYMBOL},
				{"i", types.IDENTIFIER},
				{")", types.SYMBOL},
				{"{", types.SYMBOL},
				{"let", types.KEYWORD},
				{"s", types.IDENTIFIER},
				{"=", types.SYMBOL},
				{"i", types.IDENTIFIER},
				{";", types.SYMBOL},
				{"let", types.KEYWORD},
				{"s", types.IDENTIFIER},
				{"=", types.SYMBOL},
				{"j", types.IDENTIFIER},
				{";", types.SYMBOL},
				{"let", types.KEYWORD},
				{"a", types.IDENTIFIER},
				{"[", types.SYMBOL},
				{"i", types.IDENTIFIER},
				{"]", types.SYMBOL},
				{"=", types.SYMBOL},
				{"j", types.IDENTIFIER},
				{";", types.SYMBOL},
				{"}", types.SYMBOL},
				{"else", types.KEYWORD},
				{"{", types.SYMBOL},
				{"let", types.KEYWORD},
				{"i", types.IDENTIFIER},
				{"=", types.SYMBOL},
				{"i", types.IDENTIFIER},
				{";", types.SYMBOL},
				{"let", types.KEYWORD},
				{"j", types.IDENTIFIER},
				{"=", types.SYMBOL},
				{"j", types.IDENTIFIER},
				{";", types.SYMBOL},
				{"let", types.KEYWORD},
				{"i", types.IDENTIFIER},
				{"=", types.SYMBOL},
				{"i", types.IDENTIFIER},
				{"|", types.SYMBOL},
				{"j", types.IDENTIFIER},
				{";", types.SYMBOL},
				{"}", types.SYMBOL},
				{"return", types.KEYWORD},
				{";", types.SYMBOL},
			},
			`<statements>
<ifStatement>
<keyword> if </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> i </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> s </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> i </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> s </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> j </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> a </identifier>
<symbol> [ </symbol>
<expression>
<term>
<identifier> i </identifier>
</term>
</expression>
<symbol> ] </symbol>
<symbol> = </symbol>
<expression>
<term>
<identifier> j </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
<keyword> else </keyword>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> i </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> i </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> j </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> j </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<letStatement>
<keyword> let </keyword>
<identifier> i </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> i </identifier>
</term>
<symbol> | </symbol>
<term>
<identifier> j </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
</statements>
<symbol> } </symbol>
</ifStatement>
<returnStatement>
<keyword> return </keyword>
<symbol> ; </symbol>
</returnStatement>
</statements>
`,
		},
		{
			[]tokenizer.Token{
				{"while", types.KEYWORD},
				{"(", types.SYMBOL},
				{"key", types.IDENTIFIER},
				{")", types.SYMBOL},
				{"{", types.SYMBOL},
				{"let", types.KEYWORD},
				{"key", types.IDENTIFIER},
				{"=", types.SYMBOL},
				{"key", types.IDENTIFIER},
				{";", types.SYMBOL},
				{"do", types.KEYWORD},
				{"moveSquare", types.IDENTIFIER},
				{"(", types.SYMBOL},
				{")", types.SYMBOL},
				{";", types.SYMBOL},
				{"}", types.SYMBOL},
			},
			`<statements>
<whileStatement>
<keyword> while </keyword>
<symbol> ( </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ) </symbol>
<symbol> { </symbol>
<statements>
<letStatement>
<keyword> let </keyword>
<identifier> key </identifier>
<symbol> = </symbol>
<expression>
<term>
<identifier> key </identifier>
</term>
</expression>
<symbol> ; </symbol>
</letStatement>
<doStatement>
<keyword> do </keyword>
<identifier> moveSquare </identifier>
<symbol> ( </symbol>
<expressionList>
</expressionList>
<symbol> ) </symbol>
<symbol> ; </symbol>
</doStatement>
</statements>
<symbol> } </symbol>
</whileStatement>
</statements>
`,
		},
	}

	for _, c := range cases {
		testTokenizer := NewTestTokenizer(c.tokens)
		compilationEngine := NewCompilationEngine(&testTokenizer)
		compilationEngineImpl := compilationEngine.(*CompilationEngineImpl)
		gotXML := compilationEngineImpl.compileStatements()
		assert.Equal(t, c.wantXML, gotXML)
	}
}
