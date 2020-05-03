package internal

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/kaito2/nand2tetris/internal/testdata"
	"github.com/sergi/go-diff/diffmatchpatch"

	"github.com/stretchr/testify/assert"
)

func Test_tokenizeLine(t *testing.T) {
	cases := []struct {
		input string
		want  []string
	}{
		{"Class Bar {", []string{"Class", "Bar", "{"}},
		{"method Fraction foo(int y) {", []string{"method", "Fraction", "foo", "(", "int", "y", ")", "{"}},
		{"var int temp;", []string{"var", "int", "temp", ";"}},
		{"let temp = (xxx+12)*-63; // this is comment.", []string{"let", "temp", "=", "(", "xxx", "+", "12", ")", "*", "-", "63", ";"}},
		{"// comment line.", nil},
	}

	for _, c := range cases {
		got := tokenizeLine(c.input)
		assert.Equal(t, c.want, got)
	}
}

func TestNew_advance(t *testing.T) {
	tokenizer, err := NewTokenizer("testdata/sample.jack")
	assert.NoError(t, err)
	for _, sampleToken := range testdata.SampleTokens {
		got := tokenizer.advance()
		assert.True(t, got)
		assert.Equal(t, sampleToken, tokenizer.currentToken)
	}
	got := tokenizer.advance()
	assert.False(t, got)
}

func TestTokenizer_GenerateTokenFile(t *testing.T) {
	inputFilename := "../sample/Square/Main.jack"
	outputFilename := "testoutput/output.xml"
	wantFilename := "../sample/Square/MainT.xml"

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
