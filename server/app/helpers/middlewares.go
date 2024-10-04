package helpers

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Middlewares() {}
func AuthorizationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var claims jwt.Claims
		var err error
		token := c.Request().Header.Get("Authorization")

		if len(token) <= 0 {
			return c.JSON(
				http.StatusUnauthorized,
				map[string]any{"message": "Invalid token!"},
			)
		}

		claims, err = VerifyToken(token)

		if err != nil {
			return c.JSON(
				http.StatusUnauthorized,
				map[string]any{"message": err},
			)
		}

		claims.Valid()

		return next(c)
	}
}
