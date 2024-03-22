package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type SocialMedia struct {
	ID               uint      `gorm:"primary key;auto_increment" json:"id,omitempty" `
	Name             string    `json:"name,omitempty"`
	Social_Media_Url string    `json:"social_media_url,omitempty"`
	UserID           uint      `json:"user_id,omitempty"`
	Created_at       time.Time `json:"created_at,omitempty"`
	Updated_at       time.Time `json:"updated_at"`
	User             *User     `json:"User,omitempty"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	errorMsg := ""
	if s.Name == "" {
		errorMsg += "Name Can't be Empty, "
	}

	if s.Social_Media_Url == "" {
		errorMsg += "Social_media_url Can't be Empty, "
	}

	if errorMsg != "" {
		err = errors.New(strings.TrimSuffix(errorMsg, ", "))
	}
	return
}
