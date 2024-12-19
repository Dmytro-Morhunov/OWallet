package service

import (
	"fmt"

	"OWallet.com/app/database"
	"OWallet.com/app/models"
)

func GetWallets() []models.Wallet {
	var wallets []models.Wallet
	var error = database.DB.Table("wallets").Take(&wallets)
	fmt.Println(wallets, "Wallet total: %q")
	fmt.Println("Error ---> ", error)
	return wallets

}

func GetWallet(id string) (models.Wallet, error) {
	var wallet models.Wallet
	var error = database.DB.Model(&models.Wallet{}).Find(&wallet, id)
	fmt.Println("Error", error.Error)
	fmt.Println(wallet, "Wallet id: %q")
	return wallet, error.Error
}

func CreateWallet(wallet models.Wallet) {
	fmt.Println(wallet, "wallet")
	var error = database.DB.Table("wallets").Create(&wallet)
	fmt.Println(error, "error")
	fmt.Println(wallet.ID, "Wallet created with id: %q")
}

func UpdateWallet(wallet models.Wallet) {
	error := database.DB.Model(&models.Wallet{}).Save(wallet)
	fmt.Println("Error", error)
}

func DeleteWallet(id string) {
	fmt.Println("id service: ", id)
	error := database.DB.Model(&models.Wallet{}).Where("id = ?", id).Delete(id)
	fmt.Println("Error", error)
}
