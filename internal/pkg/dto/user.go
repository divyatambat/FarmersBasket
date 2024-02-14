package dto

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          int64  `json:"ID"`
	Name        string `json:"name"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"-"`
	PhoneNumber int64  `json:"phone_number"`
	UserType    string `json:"user_type"`
}
