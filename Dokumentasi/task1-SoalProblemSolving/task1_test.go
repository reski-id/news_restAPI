package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Run("[]int{1, 2, 3} return 4", func(t *testing.T) {
		assert.Equal(t, 4, Solution([]int{1, 2, 3}))
	})

	t.Run("[]int{1, 2, 3, 4} return 5", func(t *testing.T) {
		assert.Equal(t, 5, Solution([]int{1, 2, 3, 4}))
	})

	t.Run("[]int{-1, -3} return 1", func(t *testing.T) {
		assert.Equal(t, 1, Solution([]int{-1, -3}))
	})

	t.Run("[]int{1, 3, 6, 4, 1, 2} return 5", func(t *testing.T) {
		assert.Equal(t, 5, Solution([]int{1, 3, 6, 4, 1, 2}))
	})

}
