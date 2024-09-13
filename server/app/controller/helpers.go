package controller

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func ParseRequestBodyTo(c echo.Context) (error, map[string]interface{}) {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	return err, json_map
}

func GetKeyByValue(arr map[string]interface{}, key string) string {
	if name, ok := arr[key].(string); ok {
		return name
	}
	return ""
}
