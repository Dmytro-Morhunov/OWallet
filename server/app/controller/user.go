package controller

import (
	"fmt"
	"net/http"
	"time"

	"OWallet.com/app/helpers"
	"OWallet.com/app/models"
	"OWallet.com/app/service"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	GetUsers func(e *echo.Echo) []models.User
	GetUser  func(e *echo.Echo, id uint) models.User
}

func InitUserController(g *echo.Group) {
	// Auth
	g.POST("/login", Login)
	g.POST("/registration", Registration)

	// User
	g.GET("/users", GetUsers, helpers.AuthorizationMiddleware)
	g.GET("/user/:id", GetUser, helpers.AuthorizationMiddleware)
	g.DELETE("/user/:id", DeleteUser, helpers.AuthorizationMiddleware)
	g.PUT("/user", UpdateUser, helpers.AuthorizationMiddleware)
	g.POST("/user", CreateUser, helpers.AuthorizationMiddleware)
}

// Login godoc
// @Title Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param Body body models.LoginBody true "Credentials"
// @Success 200 {object} map[string]interface{}
// @Router /api/login [post]
func Login(c echo.Context) error {
	var err, json_map = ParseRequestBodyTo(c)
	email := GetKeyByValue(json_map, "email")
	password := GetKeyByValue(json_map, "password")

	// Find user
	user, dbError := service.GetUserByEmail(email)
	if dbError != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Invalid Password:", err)
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid credentials",
		})
	}

	// Generate JWT token
	token, err := helpers.CreateToken(email)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// Get users godoc
// @Title Get users
// @Tags Users
// @Param Authorization header string true "Authorization" default(Add access token here)
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/users [get]
func GetUsers(c echo.Context) error {
	users := service.GetUsers()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": users,
	})
}

// Get user godoc
// @Title Get user
// @Tags Users
// @Param Authorization header string true "Authorization" default(Add access token here)
// @Param	id	path	int	true	"User ID"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user/{id} [get]
func GetUser(c echo.Context) error {
	id := c.Param("id")
	user, error := service.GetUser(id)
	if error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": user,
	})
}

// Update user godoc
// @Title Update user
// @Tags Users
// @Param Authorization header string true "Authorization" default(Add access token here)
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user/{id} [put]
func UpdateUser(c echo.Context) error {
	var form models.User
	form.First_name = c.FormValue("first_name")
	form.Last_name = c.FormValue("last_name")
	form.Age = c.FormValue("age")
	form.Email = c.FormValue("email")
	service.UpdateUser(form)
	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// Delete user godoc
// @Title Delete user
// @Tags Users
// @Param Authorization header string true "Authorization" default(Add access token here)
// @Accept json
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user/{id} [delete]
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	service.DeleteUser(id)
	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// Create user godoc
// @Title Create user
// @Tags Users
// @Param Authorization header string true "Authorization" default(Add access token here)
// @Accept json
// @Param user body models.User true "Create user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user [post]
func CreateUser(c echo.Context) error {
	var bytes []byte
	var user models.User
	var err, json_map = ParseRequestBodyTo(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error!",
		})
	}
	user.First_name = GetKeyByValue(json_map, "first_name")
	user.Last_name = GetKeyByValue(json_map, "last_name")
	user.Age = GetKeyByValue(json_map, "age")
	user.Email = GetKeyByValue(json_map, "email")
	var passwordStr = GetKeyByValue(json_map, "password")
	bytes, err = bcrypt.GenerateFromPassword([]byte(passwordStr), 14)
	user.Password = string(bytes)
	user.Expired_At = time.Now().UTC()
	service.CreateUser(user)
	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// Registration godoc
// @Title Registration
// @Tags Auth
// @Param Authorization header string true "Authorization" default(Add access token here)
// @Accept json
// @Produce json
// @Param Body body models.RegistrationBody true "Credentials"
// @Success 200 {object} map[string]interface{}
// @Router /api/registration [post]
func Registration(c echo.Context) error {
	var user models.User
	var parseError, json_map = ParseRequestBodyTo(c)
	if parseError != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error!",
		})
	}

	user.First_name = GetKeyByValue(json_map, "first_name")
	user.Last_name = GetKeyByValue(json_map, "last_name")
	user.Age = GetKeyByValue(json_map, "age")
	user.Email = GetKeyByValue(json_map, "email")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(GetKeyByValue(json_map, "password")), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Incorrect data",
		})
	}

	service.CreateUser(user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User registered successfully",
	})
}
