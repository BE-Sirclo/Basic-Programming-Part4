package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeNumber(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, true, PrimeNumber(1000000007), "Hasil harus sama")
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, true, PrimeNumber(1500450271), "Hasil harus sama")
	})
	t.Run("Case 3", func(t *testing.T) {
		assert.Equal(t, false, PrimeNumber(1000000000), "Hasil harus sama")
	})
	t.Run("Case 4", func(t *testing.T) {
		assert.Equal(t, true, PrimeNumber(10000000019), "Hasil harus sama")
	})
	t.Run("Case 5", func(t *testing.T) {
		assert.Equal(t, true, PrimeNumber(10000000033), "Hasil harus sama")
	})
}
