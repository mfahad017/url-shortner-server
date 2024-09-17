package models

import (
	"server/src/database"

	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	ShortURL    string `json:"shortURL" gorm:"index"`
	OriginalURL string `json:"originalURL"`
	UserID      uint   `json:"userf"`
	User        User   `json:"usero"`
}

func (u URL) TableName() string {
	return "url"
}

func (u *URL) Create() error {
	r := database.DB.Create(u)

	if r.Error != nil {
		return r.Error
	}

	return nil
}

func (u *URL) GetByShortURL(url string) error {
	r := database.DB.Where("short_url = ?", url).First(u)
	if r.Error != nil {
		return r.Error
	}

	return nil
}
