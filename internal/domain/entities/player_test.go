package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer(t *testing.T) {
	t.Run("NewPlayer", func(t *testing.T) {
		player := NewPlayer("test")

		assert.Equal(t, "test", player.Nickname)
		assert.Equal(t, 0, player.ID)
		assert.NotNil(t, player)
	})

	t.Run("MakePlayer", func(t *testing.T) {
		player := MakePlayer(1, "test")

		assert.Equal(t, "test", player.Nickname)
		assert.Equal(t, 1, player.ID)
		assert.NotNil(t, player)
	})
}
