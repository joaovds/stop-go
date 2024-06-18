package errs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	t.Run("should return a new error", func(t *testing.T) {
		message := "error message"
		error := NewError(message)

		assert.Equal(t, message, error.Message)
		assert.Equal(t, 500, error.Status)
	})

	t.Run("should return a new error with default status", func(t *testing.T) {
		assert.Equal(t, 500, NewError("error").Status)
	})
}

func TestError_Error(t *testing.T) {
	t.Run("should return the error message", func(t *testing.T) {
		message := "error message"
		error := NewError(message)
		assert.Equal(t, message, error.Error())
	})
}

func TestError_SetStatus(t *testing.T) {
	t.Run("should set the error status", func(t *testing.T) {
		error := NewError("error").SetStatus(400)
		assert.Equal(t, 400, error.Status)

		error.SetStatus(200)
		assert.Equal(t, 200, error.Status)
	})
}

func TestError_SetMessage(t *testing.T) {
	t.Run("should set the error message", func(t *testing.T) {
		error := NewError("error").SetMessage("new error")
		assert.Equal(t, "new error", error.Message)

		error.SetMessage("another error")
		assert.Equal(t, "another error", error.Message)
	})
}

func TestError_IsError(t *testing.T) {
	t.Run("should return true if error is not nil", func(t *testing.T) {
		assert.True(t, NewError("error").IsError())
	})

	t.Run("should return false if error is nil", func(t *testing.T) {
		var error *Error
		assert.False(t, error.IsError())
	})
}

func TestError_IsNil(t *testing.T) {
	t.Run("should return true if error is nil", func(t *testing.T) {
		var error *Error
		assert.True(t, error.IsNil())
	})

	t.Run("should return false if error is not nil", func(t *testing.T) {
		assert.False(t, NewError("error").IsNil())
	})
}
