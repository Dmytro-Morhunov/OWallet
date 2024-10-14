package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"OWallet.com/app/controller"
	"OWallet.com/app/helpers/tests"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	e := echo.New()

	t.Run("SUCCESS Create user", func(t *testing.T) {
		requestBody := `{"first_name": "John", "age": "30", "last_name: "Dou", "email": "test@test.com", "password": "test"}`

		req := httptest.NewRequest(http.MethodPost, "/api/createUser", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		err := controller.CreateUser(c)

		assert.NoError(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	e := echo.New()
	error := tests.MockDB()
	if error != nil {
		fmt.Println("Error: ", error)
	}

	t.Run("SUCCESS Update user", func(t *testing.T) {
		requestBody := `{"first_name": "John", "age": "30", "last_name": "Dou1", "email": "test@test.com"}`

		req := httptest.NewRequest(http.MethodPost, "/api/updateUser", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		err := controller.UpdateUser(c)

		assert.NoError(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	e := echo.New()
	error := tests.MockDB()
	if error != nil {
		fmt.Println("Error: ", error)
	}

	t.Run("SUCCESS Delete user", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodDelete, "/api/user/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := controller.DeleteUser(c)

		assert.NoError(t, err)
	})
}

func TestRegisterUser(t *testing.T) {
	e := echo.New()
	error := tests.MockDB()
	if error != nil {
		fmt.Println("Error: ", error)
	}
	t.Run("SUCCESS Register user", func(t *testing.T) {
		requestBody := `{"first_name": "John", "age": "30", "last_name": "Dou1", "email": "test@test.com", "password": "testpassword"}`

		req := httptest.NewRequest(http.MethodDelete, "/api/registration", strings.NewReader(requestBody))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := controller.Registration(c)
		assert.Equal(t, 200, rec.Result().StatusCode)
		assert.NoError(t, err)
	})
}

func TestLogin(t *testing.T) {
	e := echo.New()
	error := tests.MockDB()
	if error != nil {
		fmt.Println("Error: ", error)
	}
	t.Run("SUCCESS login", func(t *testing.T) {
		requestBody := `{"email": "test@test.com", "password": "testpassword"}`

		req := httptest.NewRequest(http.MethodDelete, "/api/login", strings.NewReader(requestBody))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := controller.Login(c)
		assert.NoError(t, err)
	})
}
