package service

import (
	"fmt"

	"OWallet.com/app/models"
	"OWallet.com/database"
)

func GetUsers() []models.User {
	var users []models.User
	var error = database.DB.Table("users").Take(&users)
	fmt.Println(users, "Users total: %q")
	fmt.Println("Error ---> ", error)
	return users

}

func GetUserByEmail(e string) (models.User, error) {
	var user models.User
	var error = database.DB.Model(&models.User{}).Find(&user, "email = ?", e)
	return user, error.Error
}

func GetUser(id string) (models.User, error) {
	var user models.User
	var error = database.DB.Model(&models.User{}).Find(&user, id)
	fmt.Println("Error", error.Error)
	fmt.Println(user, "User id: %q")
	return user, error.Error
}

func CreateUser(user models.User) {
	database.DB.Create(&user)
	fmt.Println(user.ID, "User created with id: %q")
}

func UpdateUser(user models.User) {
	error := database.DB.Model(&models.User{}).Save(user)
	fmt.Println("Error", error)
}

func DeleteUser(id string) {
	fmt.Println("id service: ", id)
	error := database.DB.Model(&models.User{}).Where("id = ?", id).Delete(id)
	fmt.Println("Error", error)
}
