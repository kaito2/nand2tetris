package compilationengine

import (
	"testing"

	"github.com/kaito2/nand2tetris/internal/tokenizer"
	"github.com/stretchr/testify/assert"
)

func TestCompilationEngineImpl_Compile(t *testing.T) {
	tk, err := tokenizer.NewTokenizer("../../sample/Square/Main.jack")
	assert.NoError(t, err)
	cp := NewCompilationEngine(tk)
	cp.Compile()
}
