package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairSum(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, []int{1, 3}, PairSum([]int{1, 2, 3, 4, 6}, 6), "Hasil harus sama")
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, []int{0, 2}, PairSum([]int{2, 5, 9, 11}, 11), "Hasil harus sama")
	})
	t.Run("Case 3", func(t *testing.T) {
		assert.Equal(t, []int{2, 3}, PairSum([]int{1, 3, 5, 7}, 12), "Hasil harus sama")
	})
	t.Run("Case 4", func(t *testing.T) {
		assert.Equal(t, []int{1, 2}, PairSum([]int{1, 4, 6, 8}, 10), "Hasil harus sama")
	})
	t.Run("Case 5", func(t *testing.T) {
		assert.Equal(t, []int{0, 1}, PairSum([]int{1, 5, 6, 7}, 6), "Hasil harus sama")
	})
}
