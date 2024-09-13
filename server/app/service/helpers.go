package service

import "gorm.io/gorm"

func HandleError(result *gorm.DB) string {
	if result.Error != nil {
		return "Error"
	}
	return ""
}
