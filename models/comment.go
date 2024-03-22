package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID         uint      `gorm:"primary key;auto_increment" json:"id"`
	UserID     uint      `json:"user_id"`
	PhotoID    uint      `json:"photo_id"`
	Message    string    `json:"message"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	User       *User     `json:"User,omitempty"`
	Photo      *Photo    `json:"Photo,omitempty"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	errorMsg := ""
	if c.Message == "" {
		errorMsg += "Message Can't be Empty, "
	}

	if errorMsg != "" {
		err = errors.New(strings.TrimSuffix(errorMsg, ", "))
	}
	return
}
