package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolParser_ScanLCommands(t *testing.T) {
	filepath := "testdata/Max.asm"
	symbolParser := NewSymbolParser(filepath)
	err := symbolParser.ScanLCommands()
	assert.NoError(t, err)

	want := map[string]uint16{
		"OUTPUT_FIRST":  uint16(10),
		"OUTPUT_D":      uint16(12),
		"INFINITE_LOOP": uint16(14),
	}

	assert.Equal(t, want, symbolParser.symbolTable)
}

func Test_addSymbol(t *testing.T) {
	cases := []struct {
		input []string
		want  map[string]uint16
	}{
		{[]string{"LOOP", "Xxx", "Yyy"}, map[string]uint16{"LOOP": 0, "Xxx": 1, "Yyy": 2}},
	}

	filepath := "sample"
	for _, c := range cases {
		symbolParser := NewSymbolParser(filepath)
		for _, symbol := range c.input {
			err := symbolParser.addSymbol(symbol, symbolParser.romCounter)
			assert.NoError(t, err)
			symbolParser.romCounter++
		}
		assert.Equal(t, c.want, symbolParser.symbolTable)
	}
}
