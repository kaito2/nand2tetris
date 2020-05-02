package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidCharacter(t *testing.T) {
	validCharacters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	for _, c := range validCharacters {
		assert.True(t, isValidCharacter(c), fmt.Sprintf("character: %s", string(c)))
	}

	invalidCharacters := "<>*+=-"
	for _, c := range invalidCharacters {
		assert.False(t, isValidCharacter(c))
	}
}

func TestIsIdentifier(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"abc", true},
		{"a_b_c", true},
		{"ABC", true},
		{"A_B_C", true},
		{"a1", true},
		{"Abc", true},
		{"1ab", false},
		{"a*b", false},
	}

	for _, c := range cases {
		assert.Equal(t, IsIdentifier(c.input), c.want)
	}
}
