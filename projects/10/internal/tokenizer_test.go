package internal

import (
	"testing"

	"github.com/kaito2/nand2tetris/internal/testdata"

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
