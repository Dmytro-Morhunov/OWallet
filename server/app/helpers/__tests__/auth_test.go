package helpers

import (
	"testing"

	"OWallet.com/app/helpers"
	"github.com/stretchr/testify/assert"
)

func TestCreteToken(t *testing.T) {
	t.Run("SUCCESS create token", func(t *testing.T) {
		email := "test@test.com"
		token, err := helpers.CreateToken(email)

		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("ERROR create token with empty email", func(t *testing.T) {
		email := ""
		token, err := helpers.CreateToken(email)

		assert.Error(t, err)
		assert.Empty(t, token)
	})
}

func TestLogin(t *testing.T) {
	t.Run("SUCCESS Login", func(t *testing.T) {
		email := "test@test.com"
		token, err := helpers.CreateToken(email)

		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		claims, err := helpers.VerifyToken(token)
		isValidError := claims.Valid()

		assert.NoError(t, err)
		assert.NoError(t, isValidError)
	})

	t.Run("FAILURE Login with wrong token", func(t *testing.T) {
		_, err := helpers.VerifyToken("test")
		assert.Error(t, err)
	})

	t.Run("FAILURE Login with empty token", func(t *testing.T) {
		_, err := helpers.VerifyToken("")
		assert.Error(t, err)
	})
}
