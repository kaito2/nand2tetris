package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParser(t *testing.T) {
	existsFile := "testdata/test.asm"
	_, err := NewParser(existsFile)
	assert.NoError(t, err)

	notExistsFile := "testdata/notExists.asm"
	_, err = NewParser(notExistsFile)
	assert.Error(t, err)
}

func Test_advance(t *testing.T) {
	want := []string{
		"@2",
		"D=A",
		"@3",
		"D=D+A",
		"@0",
		"M=D",
	}
	filename := "testdata/test.asm"

	p, err := NewParser(filename)
	assert.NoError(t, err)
	for _, wantCmd := range want {
		got := p.advance()
		assert.True(t, got)
		assert.Equal(t, wantCmd, p.currentCommand)
	}
	hasMoreCommand := p.advance()
	assert.False(t, hasMoreCommand)
}
