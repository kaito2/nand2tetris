package types

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIsIntegerConstant(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"-1", false},
		{"0", true},
		{"1", true},
		{"123", true},
		{"32767", true},
		{"32768", false},

		{"string", false},
		{"*", false},
	}

	for _, c := range cases {
		assert.Equal(t, IsIntegerConstant(c.input), c.want)
	}
}
