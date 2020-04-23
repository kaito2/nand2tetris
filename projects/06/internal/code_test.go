package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_compBits(t *testing.T) {
	cases := []struct {
		input string
		want  uint16
	}{
		{"D+M", 0b1000010},
		{"D+A", 0b0000010},
	}

	for _, c := range cases {
		got := compBits(c.input)
		assert.Equal(t, c.want, got)
	}
}
