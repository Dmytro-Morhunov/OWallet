package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"OWallet.com/app/controller"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestParseRequestBodyTo(t *testing.T) {
	e := echo.New()
	requestBody := `{"name": "John", "age": "30"}`
	expectedMock := map[string]interface{}{
		"name": "John",
		"age":  "30",
	}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err, result := controller.ParseRequestBodyTo(c)

	assert.NoError(t, err)
	assert.Equal(t, expectedMock, result)
}

func TestGetKeyByValue(t *testing.T) {
	testMap := map[string]interface{}{
		"name":  "Alice",
		"email": "alice@example.com",
		"age":   "25",
	}

	name := controller.GetKeyByValue(testMap, "name")
	assert.Equal(t, "Alice", name)

	age := controller.GetKeyByValue(testMap, "age")
	assert.Equal(t, "25", age)

	nonExistent := controller.GetKeyByValue(testMap, "nonexistent")
	assert.Equal(t, "", nonExistent)
}
