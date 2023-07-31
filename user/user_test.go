package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lpgod/task/user"
)

func TestHashPassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := user.HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestVerifyPassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := user.HashPassword(password)
	assert.NoError(t, err)

	err = user.VerifyPassword(hashedPassword, password)
	assert.NoError(t, err)

	invalidPassword := "wrongpassword"
	err = user.VerifyPassword(hashedPassword, invalidPassword)
	assert.Error(t, err)
}

func TestUserRepository(t *testing.T) {
	// Here you can write unit tests for your UserRepository implementation
	// Mock the data storage and retrieval operations to test your methods
	// Make sure to cover scenarios like CreateUser, GetUserByEmail, UpdateUser, and DeleteUser
}
