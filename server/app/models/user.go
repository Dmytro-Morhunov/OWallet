package models

import "time"

// User - Data Model To Store user Data in
type User struct {
	ID         uint `gorm:"primarykey"`
	Expired_At time.Time
	RegistrationBody
}

type LoginBody struct {
	Email    string `gorm:"unique" json:"email" example:"JohnDoe@mail.com"`
	Password string `json:"password"`
}

type RegistrationBody struct {
	First_name string `json:"first_name" example:"John"`
	Last_name  string `json:"last_name" example:"Doe"`
	Email      string `gorm:"unique" json:"email" example:"JohnDoe@mail.com"`
	Age        string `json:"age" example:"21"`
	Password   string `json:"-"`
}
