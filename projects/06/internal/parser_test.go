package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewParser(t *testing.T) {
	want := Parser{}
	got := NewParser()

	assert.Equal(t, want, got)
}
