package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "jdoe@hotmail.com", "123$456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "jdoe@hotmail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "jdoe@hotmail.com", "123$456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123$456"))
	assert.False(t, user.ValidatePassword("123$4567"))
	assert.NotEqual(t, "123456", user.Password)
}
