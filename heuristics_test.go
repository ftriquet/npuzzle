package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeuristics(t *testing.T) {
	b := getFinalBoard(3)
	c := getFinalBoard(3)
	c[0][1] = 1
	c[0][0] = 2
	d := board{
		{1, 4, 2},
		{0, 3, 8},
		{5, 7, 6},
	}
	assert.Equal(t, 2, int(Differences(b, c)))
	assert.Equal(t, 2, int(Distances(b, c)))
	assert.Equal(t, 22, int(Differences(b, d)))
	assert.Equal(t, 12, int(Distances(b, d)))
}
