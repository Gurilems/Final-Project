package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID         uint      `gorm:"primary key;auto_increment" json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url"`
	UserID     uint      `json:"user_id"`
	Created_at time.Time `json:"created_at,omitempty"`
	Updated_at time.Time `json:"updated_at,omitempty"`
	User       *User     `json:"User,omitempty"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	errorMsg := ""
	if p.Title == "" {
		errorMsg += "Title Can't be Empty, "
	}

	if p.Photo_url == "" {
		errorMsg += "Photo_url Can't be Empty, "
	}

	if errorMsg != "" {
		err = errors.New(strings.TrimSuffix(errorMsg, ", "))
	}
	return
}
