package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsing(t *testing.T) {
	b := board{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	f, e := os.Open("board.txt")
	require.True(t, e == nil, "Open Error, can't test")
	bo, e := parseBoard(f)
	require.True(t, e == nil)
	require.True(t, b.equals(bo))
	f, e = os.Open("invalid_board.txt")
	require.True(t, e == nil, "Open Error, can't test")
	bo, e = parseBoard(f)
	assert.True(t, e != nil)
	f, e = os.Open("invalid_board_2.txt")
	require.True(t, e == nil, "Open Error, can't test")
	bo, e = parseBoard(f)
	assert.True(t, e != nil)
	f, e = os.Open("invalid_board_3.txt")
	require.True(t, e == nil, "Open Error, can't test")
	bo, e = parseBoard(f)
	assert.True(t, e != nil)
}
