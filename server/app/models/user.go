package models

import "time"

// User - Data Model To Store user Data in
type User struct {
	ID         uint   `gorm:"primarykey"`
	First_name string `json:"first_name" example:"John"`
	Last_name  string `json:"last_name" example:"Doe"`
	Email      string `gorm:"unique" json:"email" example:"JohnDoe@mail.com"`
	Age        string `json:"age" example:"21"`
	Password   string `json:"-"`
	Expired_At time.Time
}
