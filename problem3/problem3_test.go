package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinArrayRemoveDuplicate(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, []string{"apel", "anggur", "lemon", "leci", "nanas"}, JoinArrayRemoveDuplicate([]string{"apel", "anggur"}, []string{"lemon", "leci", "nanas"}), "Hasil harus sama")
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, []string{"samsung", "apple", "sony", "xiaomi"}, JoinArrayRemoveDuplicate([]string{"samsung", "apple"}, []string{"apple", "sony", "xiaomi"}), "Hasil harus sama")
	})
	t.Run("Case 3", func(t *testing.T) {
		assert.Equal(t, []string{"football", "basketball"}, JoinArrayRemoveDuplicate([]string{"football", "basketball"}, []string{"basketball", "football"}), "Hasil harus sama")
	})
}
