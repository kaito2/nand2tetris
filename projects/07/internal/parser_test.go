package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParser(t *testing.T) {
	testFilename := "testdata/SimpleAdd.vm"
	_, err := NewParser(testFilename)
	assert.NoError(t, err)
}

func Test_advance(t *testing.T) {
	cases := []struct {
		filename     string
		wantCommands []string
	}{
		{"testdata/SimpleAdd.vm", []string{"push constant 7", "push constant 8", "add"}},
	}

	for _, c := range cases {
		parser, err := NewParser(c.filename)
		assert.NoError(t, err)
		for _, wantCmd := range c.wantCommands {
			hasCommand := parser.advance()
			assert.True(t, hasCommand)
			assert.Equal(t, wantCmd, parser.currentCommand)
		}
		assert.False(t, parser.advance())
	}
}
