package internal

import (
	"testing"

	"github.com/kaito2/nand2tetris/internal/testdata"

	"github.com/stretchr/testify/assert"
)

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
