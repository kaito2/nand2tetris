package types

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_isString(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{`"hoge"`, true},
		{`"hoge`, false},
		{`hoge"`, false},
		{`hoge`, false},
	}

	for _, c := range cases {
		assert.Equal(t, isString(c.input), c.want)
	}
}

func TestGetString(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{`"a"`, `a`},
		{`"abc"`, `abc`},
	}

	for _, c := range cases {
		assert.Equal(t, GetString(c.input), c.want)
	}
}
