package controller

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"time"

	"OWallet.com/app/helpers"
	"OWallet.com/app/models"
	"OWallet.com/app/service"
	"github.com/golang-jwt/jwt"
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
	g.POST("/register", Register)

	// User
	g.GET("/users", GetUsers)
	g.GET("/user/:id", GetUser, helpers.AuthorizationMiddleware)
	g.DELETE("/user/:id", DeleteUser)
	g.PUT("/user", UpdateUser)
	g.POST("/user", CreateUer)
}

// Login
// @Title Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param creds body models.User.Email true "Login"
// @Success 200 {object} map[string]interface{}
// @Router /api/login [post]
func Login(c echo.Context) error {
	var err, json_map = ParseRequestBodyTo(c)
	email := GetKeyByValue(json_map, "email")
	password := GetKeyByValue(json_map, "password")
	fmt.Println("Email: ", email)
	fmt.Println("Password: ", password)

	// Find user
	user, dbError := service.GetUserByEmail(email)
	fmt.Println(user, "dbUser")
	if dbError != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	fmt.Println(err, "PASSWORD ERROR")
	if err != nil {
		fmt.Println("Invalid Password:", err)
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid credentials",
		})
	}

	secretKey := make([]byte, 64)
	_, generateError := rand.Read(secretKey)
	if generateError != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
	}

	// Generate JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	})
	token, err := claims.SignedString([]byte(secretKey))
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
// @Param Authorization header string true "Authorization" default(Bearer <Add access token here>)
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
// @Accept json
// @Param			id	path		int	true	"User ID"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user/{id} [delete]
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("id: ", id)

	service.DeleteUser(id)
	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// Create user godoc
// @Title Create user
// @Tags Users
// @Accept json
// @Param user body models.User true "Create user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user [post]
func CreateUer(c echo.Context) error {
	var bytes []byte
	var user models.User
	var err, json_map = ParseRequestBodyTo(c)
	user.First_name = GetKeyByValue(json_map, "first_name")
	user.Last_name = GetKeyByValue(json_map, "last_name")
	user.Age = GetKeyByValue(json_map, "age")
	user.Email = GetKeyByValue(json_map, "email")
	var passwordStr = GetKeyByValue(json_map, "password")
	bytes, err = bcrypt.GenerateFromPassword([]byte(passwordStr), 14)
	user.Password = string(bytes)
	user.Expired_At = time.Now().UTC()
	fmt.Println("json_map ", json_map)
	fmt.Println("err ", err)
	service.CreateUser(user)
	return c.JSON(http.StatusOK, map[string]interface{}{})
}

func Register(c echo.Context) error {
	var user models.User
	var err, json_map = ParseRequestBodyTo(c)
	user.First_name = GetKeyByValue(json_map, "first_name")
	user.Last_name = GetKeyByValue(json_map, "last_name")
	user.Age = GetKeyByValue(json_map, "age")
	user.Email = GetKeyByValue(json_map, "email")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(GetKeyByValue(json_map, "password")), bcrypt.DefaultCost)

	user.Password = string(hashedPassword)

	fmt.Println("json_map ", json_map)
	fmt.Println("err ", err)
	service.CreateUser(user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User registered successfully",
	})
}
