package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceToMap(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		o []string
		m map[string]string
	}{
		{
			[]string{},
			map[string]string{},
		}, {
			[]string{"a=", "b==b", " c = e "},
			map[string]string{"a": "", "b": "=b", "c": "e"},
		},
	}
	for _, t := range tests {
		assert.EqualValues(t.m, sliceToMap(t.o))
	}

}
