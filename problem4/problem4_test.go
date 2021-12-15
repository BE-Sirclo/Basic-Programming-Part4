package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMunculSekali(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, []int{4}, MunculSekali("1234123"), "Hasil harus sama")
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, []int{6, 3}, MunculSekali("76523752"), "Hasil harus sama")
	})
	t.Run("Case 3", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5}, MunculSekali("12345"), "Hasil harus sama")
	})
	t.Run("Case 4", func(t *testing.T) {
		assert.Equal(t, []int{}, MunculSekali("1122334455"), "Hasil harus sama")
	})
	t.Run("Case 5", func(t *testing.T) {
		assert.Equal(t, []int{8, 7, 2, 5, 4}, MunculSekali("0872504"), "Hasil harus sama")
	})
}
