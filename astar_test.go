package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPermutations(t *testing.T) {
	b := board{
		{1, 4, 2},
		{0, 3, 8},
		{5, 7, 6},
	}
	expectedPermutations := []board{
		board{
			{0, 4, 2},
			{1, 3, 8},
			{5, 7, 6},
		},
		board{
			{1, 4, 2},
			{5, 3, 8},
			{0, 7, 6},
		},
		board{
			{1, 4, 2},
			{3, 0, 8},
			{5, 7, 6},
		},
	}
	states := b.getNextStates()
	require.Equal(t, 3, len(states), "Should be 3 available states")
	for i := range expectedPermutations {
		assert.True(t, expectedPermutations[i].equals(states[i]))
	}

}
