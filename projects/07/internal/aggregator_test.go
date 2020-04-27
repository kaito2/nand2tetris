package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAggregator(t *testing.T) {
	testDir := "testdata/test-vm-dir"
	_, err := NewAggregator(testDir)
	assert.NoError(t, err)
}

func TestParser_getVMFilenames(t *testing.T) {
	cases := []struct {
		input      string
		want       []string
		wantsError bool
	}{
		{"testdata/test-vm-dir", []string{"testdata/test-vm-dir/Main.vm", "testdata/test-vm-dir/Sys.vm"}, false},
		{"testdata/test-vm-dir/Main.vm", []string{"testdata/test-vm-dir/Main.vm"}, false},
		{"testdata/test-invalid-dir", []string{}, false},
		{"testdata/not-exists-dir", nil, true},
		{"testdata/test-vm-dir/not-exists-file.vm", nil, true},
	}

	for _, c := range cases {
		got, err := getVMFilenames(c.input)
		if c.wantsError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.ElementsMatch(t, c.want, got)
	}
}
