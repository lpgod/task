package auth_test

import (
	"testing"
	"time"

	"github.com/lpgod/task/auth"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := auth.HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestVerifyPassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := auth.HashPassword(password)
	assert.NoError(t, err)

	err = auth.VerifyPassword(hashedPassword, password)
	assert.NoError(t, err)

	invalidPassword := "wrongpassword"
	err = auth.VerifyPassword(hashedPassword, invalidPassword)
	assert.Error(t, err)
}

func TestGenerateAndParseToken(t *testing.T) {
	secretKey := "mysecretkey"
	expirationTime := time.Hour

	user := &auth.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
	}

	// Generate a new token
	tokenString, err := auth.GenerateToken(user, secretKey, expirationTime)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Parse the token
	claims, err := auth.ParseToken(tokenString, secretKey)
	assert.NoError(t, err)
	assert.NotNil(t, claims)

	// Check the claims
	assert.Equal(t, float64(user.ID), claims["id"])
	assert.Equal(t, user.Name, claims["name"])
	assert.Equal(t, user.Email, claims["email"])
}

func TestParseInvalidToken(t *testing.T) {
	// An invalid token
	invalidTokenString := "invalid.token.string"
	secretKey := "mysecretkey"

	// Parse the invalid token
	claims, err := auth.ParseToken(invalidTokenString, secretKey)
	assert.Error(t, err)
	assert.Nil(t, claims)
}
