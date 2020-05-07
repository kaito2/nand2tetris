package tokenizer

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/kaito2/nand2tetris/internal/types"

	"github.com/sergi/go-diff/diffmatchpatch"

	"github.com/stretchr/testify/assert"
)

func Test_tokenizeLine(t *testing.T) {
	cases := []struct {
		input string
		want  []Token
	}{
		{
			"class Bar {",
			[]Token{
				{"class", types.KEYWORD},
				{"Bar", types.IDENTIFIER},
				{"{", types.SYMBOL},
			},
		},
		{
			"method Fraction foo(int y) {",
			[]Token{
				{"method", types.KEYWORD},
				{"Fraction", types.IDENTIFIER},
				{"foo", types.IDENTIFIER},
				{"(", types.SYMBOL},
				{"int", types.KEYWORD},
				{"y", types.IDENTIFIER},
				{")", types.SYMBOL},
				{"{", types.SYMBOL},
			},
		},
		{
			"var int temp;",
			[]Token{
				{"var", types.KEYWORD},
				{"int", types.KEYWORD},
				{"temp", types.IDENTIFIER},
				{";", types.SYMBOL},
			},
		},
		{
			"let temp = (xxx+12)*-63; // this is comment.",
			[]Token{
				{"let", types.KEYWORD},
				{"temp", types.IDENTIFIER},
				{"=", types.SYMBOL},
				{"(", types.SYMBOL},
				{"xxx", types.IDENTIFIER},
				{"+", types.SYMBOL},
				{"12", types.INT_CONST},
				{")", types.SYMBOL},
				{"*", types.SYMBOL},
				{"-", types.SYMBOL},
				{"63", types.INT_CONST},
				{";", types.SYMBOL},
			},
		},
		{
			"// comment line.",
			nil,
		},
	}

	for _, c := range cases {
		got := tokenizeLine(c.input)
		assert.Equal(t, c.want, got)
	}
}

func TestTokenizerImpl_Advance(t *testing.T) {
	sampleTokens := []Token{
		{"class", types.KEYWORD},
		{"Bar", types.IDENTIFIER},
		{"{", types.SYMBOL},

		{"method", types.KEYWORD},
		{"Fraction", types.IDENTIFIER},
		{"foo", types.IDENTIFIER},
		{"(", types.SYMBOL},
		{"int", types.KEYWORD},
		{"y", types.IDENTIFIER},
		{")", types.SYMBOL},
		{"{", types.SYMBOL},

		{"var", types.KEYWORD},
		{"int", types.KEYWORD},
		{"temp", types.IDENTIFIER},
		{";", types.SYMBOL},

		{"let", types.KEYWORD},
		{"temp", types.IDENTIFIER},
		{"=", types.SYMBOL},
		{"(", types.SYMBOL},
		{"xxx", types.IDENTIFIER},
		{"+", types.SYMBOL},
		{"12", types.INT_CONST},
		{")", types.SYMBOL},
		{"*", types.SYMBOL},
		{"-", types.SYMBOL},
		{"63", types.INT_CONST},
		{";", types.SYMBOL},

		{"}", types.SYMBOL},
		{"}", types.SYMBOL},
	}

	tokenizer, err := NewTokenizer("testdata/sample.jack")
	assert.NoError(t, err)
	for _, sampleToken := range sampleTokens {
		got := tokenizer.Advance()
		assert.True(t, got)
		assert.Equal(t, sampleToken, tokenizer.CurrentToken())
	}
	got := tokenizer.Advance()
	assert.False(t, got)
}

func TestTokenizer_GenerateTokenFile(t *testing.T) {
	// TODO: ファイル命名規則をメモ
	inputFilenames := []string{
		"../sample/Square/Main.jack",
	}

	for _, inputFilename := range inputFilenames {
		outputBase := strings.ReplaceAll(path.Base(inputFilename), ".jack", "T.xml")
		outputFilename := "testoutput/" + outputBase
		wantFilename := path.Dir(inputFilename) + "/" + outputBase

		testTokenizer_GenerateTokenFile(t, inputFilename, outputFilename, wantFilename)
	}
}

func testTokenizer_GenerateTokenFile(t *testing.T, inputFilename, outputFilename, wantFilename string) {
	tokenizer, err := NewTokenizer(inputFilename)
	assert.NoError(t, err)
	err = tokenizer.GenerateTokenFile(outputFilename)
	assert.NoError(t, err)

	outputFile := _readFileContent(outputFilename)
	// windows で編集したファイルなので改行コードが \r\n になっているため修正
	wantFile := strings.ReplaceAll(_readFileContent(wantFilename), "\r\n", "\n")

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(wantFile, outputFile, false)
	var validDiffs []diffmatchpatch.Diff
	for _, diff := range diffs {
		if diff.Type != diffmatchpatch.DiffEqual {
			validDiffs = append(validDiffs, diff)
		}
	}

	if len(validDiffs) != 0 {
		t.Errorf("assertion failed: %d diff", len(validDiffs))
		t.Log(dmp.DiffPrettyText(diffs))
	}
}

func _readFileContent(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to os.Open: %v", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("failed to ioutil.ReadAll: %v", err)
	}
	return string(b)
}
