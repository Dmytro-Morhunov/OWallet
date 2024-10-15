package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"OWallet.com/app/helpers"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func MockVerifyToken(token string) (jwt.Claims, error) {
	if token == "valid_token" {
		return &jwt.StandardClaims{Subject: "test_user"}, nil
	}
	return nil, errors.New("invalid token")
}

func TestAuthorizationMiddleware(t *testing.T) {
	e := echo.New()
	token, _ := helpers.CreateToken("test@test.com")

	tests := []struct {
		name               string
		token              string
		expectedStatusCode int
		expectedMessage    string
	}{
		{
			name:               "No Token",
			token:              "",
			expectedStatusCode: http.StatusUnauthorized,
			expectedMessage:    "Invalid token!",
		},
		{
			name:               "Invalid Token",
			token:              "invalid_token",
			expectedStatusCode: http.StatusUnauthorized,
			expectedMessage:    "",
		},
		{
			name:               "Valid Token",
			token:              token,
			expectedStatusCode: http.StatusOK,
			expectedMessage:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", tt.token)
			}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Create a next handler that simply returns a success status
			nextHandler := func(c echo.Context) error {
				return c.String(http.StatusOK, "success")
			}

			middleware := helpers.AuthorizationMiddleware(nextHandler)
			middleware(c)

			assert.Equal(t, tt.expectedStatusCode, rec.Code)

			if tt.expectedMessage != "" {
				var response map[string]interface{}
				json.Unmarshal(rec.Body.Bytes(), &response)
				assert.Equal(t, tt.expectedMessage, response["message"])
			}
		})
	}
}
