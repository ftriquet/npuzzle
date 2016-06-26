package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsing(t *testing.T) {
	validMaps := []string{
		"maps/valid_aligned.txt",
		"maps/valid_basic.txt",
		"maps/valid_comments.txt",
		"maps/valid_wtf.txt",
	}
	invalidMaps := []string{
		"maps/invalid_no_zero.txt",
		"maps/invalid_line_sizes.txt",
		"maps/invalid_wrong_size.txt",
		"maps/invalid_non_number_size.txt",
	}
	for _, m := range validMaps {
		f, err := os.Open(m)
		require.True(t, err == nil)
		if err != nil {
			_, err := parseBoard(f)
			assert.True(t, err == nil)
		}
	}
	for _, m := range invalidMaps {
		f, err := os.Open(m)
		require.True(t, err == nil)
		if err != nil {
			_, err := parseBoard(f)
			assert.True(t, err != nil)
		}
	}
}
