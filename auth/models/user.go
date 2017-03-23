package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
}

func (u *User) Validate() error {
	if !govalidator.IsEmail(u.Username) {
		return errors.New("Invalid email")
	}
	if len(u.Username) <= 0 || len(u.Username) > 255 {
		return errors.New("Username length invalid")
	}
	if !govalidator.IsAlphanumeric(u.Password) {
		return errors.New("Invalid characters in password")
	}
	if len(u.Password) <= 8 || len(u.Password) > 255 {
		return errors.New("Password length invalid")
	}
	return nil
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
