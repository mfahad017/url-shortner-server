package models

import (
	"server/src/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// type UserModel struct{}

func (User) TableName() string {
	return "user"
}

func (u *User) GetUserByEmail() (*User, error) {

	return u, nil

}

func (u *User) Create() (*User, error) {

	r := database.DB.Create(u)

	if r.Error != nil {
		return nil, r.Error
	}

	return u, nil
}

func (u *User) GetByEmail(email string) error {

	r := database.DB.Where("email = ?", email).First(u)

	if r.Error != nil {
		return r.Error
	}

	return nil
}
