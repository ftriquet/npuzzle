package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubBoards(t *testing.T) {
	b := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 0, 15, 6},
		{10, 9, 8, 7},
	}
	sub, _ := getSubBoard(b)
	require.True(t, len(sub) == 2)
	require.True(t, sub[0][0] == 13)
	require.True(t, sub[0][1] == 14)
	require.True(t, sub[1][0] == 0)
	require.True(t, sub[1][1] == 15)
	sub, finish := getSubBoard(sub)
	assert.True(t, finish)

	final := getFinalBoard(3)
	assert.True(t, final[0][0] == 1)
	assert.True(t, final[0][1] == 2)
	assert.True(t, final[0][2] == 3)
	assert.True(t, final[1][0] == 8)
	assert.True(t, final[1][1] == 0)
	assert.True(t, final[1][2] == 4)
	assert.True(t, final[2][0] == 7)
	assert.True(t, final[2][1] == 6)
	assert.True(t, final[2][2] == 5)
}
