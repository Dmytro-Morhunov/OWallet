package controller

import (
	"net/http"

	"OWallet.com/app/helpers"
	"github.com/labstack/echo/v4"
)

func InitWalletController(g *echo.Group) {
	walletGroup := g.Group("/wallet")
	walletGroup.GET("/", GetWallet, helpers.AuthorizationMiddleware)
}

func GetWallet(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]interface{}{})
}
