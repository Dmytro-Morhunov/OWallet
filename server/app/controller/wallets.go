package controller

import (
	"net/http"
	"strconv"

	"OWallet.com/app/models"
	"OWallet.com/app/service"
	"github.com/labstack/echo/v4"
)

func InitWalletController(g *echo.Group) {
	g.GET("/wallets", GetWallets)
	g.GET("/wallets/:id", GetUserWallet)
	g.POST("/wallets", CreateWallet)
	g.PUT("/wallets", UpdateWallet)
	g.DELETE("/wallets/:id", DeleteWallet)
}

// Get wallets godoc
// @Title Get wallets
// @Tags Wallets
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/wallets [get]
func GetWallets(c echo.Context) error {
	wallets := service.GetWallets()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": wallets,
	})
}

func GetUserWallet(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// Create wallet godoc
// @Title Create wallet
// @Tags Wallets
// @Accept json
// @Param wallet body models.Wallet true "Create wallet"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/wallets [post]
func CreateWallet(c echo.Context) error {
	var wallet models.Wallet
	var err, json_map = ParseRequestBodyTo(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error!",
		})
	}
	wallet.Name = GetKeyByValue(json_map, "name")
	wallet.Currency_Type, _ = strconv.Atoi(GetKeyByValue(json_map, "currency_type"))

	wallet.Balance = 0
	service.CreateWallet(wallet)

	return c.JSON(http.StatusOK, nil)
}

func UpdateWallet(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func DeleteWallet(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
