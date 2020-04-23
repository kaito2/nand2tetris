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

func Test_commandType(t *testing.T) {
	cases := []struct {
		input string
		want  CommandType
	}{
		{"@2", ACommand},
		{"D=A", CCommand},
		{"@3", ACommand},
		{"D=D+A", CCommand},
		{"@0", ACommand},
		{"M=D", CCommand},
		{"(Loop)", LCommand},
		{"(Xxx)", LCommand},
		{"D;JGT ", CCommand},
		{"0;JMP", CCommand},
	}

	for _, c := range cases {
		got := commandType(c.input)
		assert.Equal(t, c.want, got)
	}
}

func Test_symbol(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"@2", "2"},
		{"@Xxx", "Xxx"},
		{"(Xxx)", "Xxx"},
	}

	for _, c := range cases {
		got := symbol(c.input)
		assert.Equal(t, c.want, got)
	}
}

func Test_dest(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"D=A", "D"},
		{"D=D+A", "D"},
		{"M=D", "M"},
		{"D;JGT ", DestNull},
		{"0;JMP", DestNull},
	}

	for _, c := range cases {
		got := dest(c.input)
		assert.Equal(t, c.want, got)
	}
}

func Test_comp(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"D=A", "A"},
		{"D=D+A", "D+A"},
		{"M=D", "D"},
		{"D;JGT ", "D"},
		{"0;JMP", "0"},
		{"M=!M", "!M"},
	}

	for _, c := range cases {
		got := comp(c.input)
		assert.Equal(t, c.want, got)
	}
}
