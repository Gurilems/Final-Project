package models

import (
	"errors"
	"final-project/helpers"

	"net/mail"
	"strings"

	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         int       `gorm:"primary key;auto_increment" json:"id"`
	Username   string    `gorm:"unique" json:"username"`
	Email      string    `gorm:"unique" json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Age        int       `json:"age,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
	Updated_at time.Time `json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	errorMsg := ""
	if u.Email == "" {
		errorMsg += "Email Can't be Empty, "
	} else {
		_, errEmail := mail.ParseAddress(u.Email)
		if errEmail != nil {
			errorMsg += "Wrong Email Format, "
		}
	}

	if u.Username == "" {
		errorMsg += "Username Can't be Empty, "
	}

	if u.Password == "" {
		errorMsg += "Password Can't be Empty, "
	} else {
		if len(u.Password) < 6 {
			errorMsg += "Password Must Contains At least 6 Characters "
		}
	}

	if u.Age == 0 {
		errorMsg += "Age Can't be Empty or Zero, "
	} else {
		if u.Age <= 8 {
			errorMsg += "Age Must be At least 9, "
		}
	}

	if errorMsg != "" {
		err = errors.New(strings.TrimSuffix(errorMsg, ", "))
	}
	u.Password = helpers.PasswordHashing(u.Password)
	return
}
