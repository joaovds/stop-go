package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoom(t *testing.T) {
	t.Run("NewRoom", func(t *testing.T) {
		room := NewRoom("test")

		assert.Equal(t, "test", room.Name)
		assert.Equal(t, 0, room.ID)
		assert.Equal(t, "", room.Code)
		assert.NotNil(t, room)
	})

	t.Run("MakeRoom", func(t *testing.T) {
		room := MakeRoom(1, "test", "code")

		assert.Equal(t, "test", room.Name)
		assert.Equal(t, 1, room.ID)
		assert.Equal(t, "code", room.Code)
		assert.NotNil(t, room)
	})
}
