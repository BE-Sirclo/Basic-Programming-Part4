package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, 8, Pow(2, 3), "Hasil harus sama")
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, 49, Pow(7, 2), "Hasil harus sama")
	})
	t.Run("Case 3", func(t *testing.T) {
		assert.Equal(t, 100000, Pow(10, 5), "Hasil harus sama")
	})
	t.Run("Case 4", func(t *testing.T) {
		assert.Equal(t, 24137569, Pow(17, 6), "Hasil harus sama")
	})
	t.Run("Case 5", func(t *testing.T) {
		assert.Equal(t, 125, Pow(5, 3), "Hasil harus sama")
	})
}
