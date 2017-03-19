package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/asaskevich/govalidator"
	"errors"
)

type (
	User struct {
		Id 	 bson.ObjectId	`json:"id" bson:"_id"` 
		Username string		`json:"username" bson:"username"`
		Password string		`json:"password" bson:"password"`
	}
)

func (u *User) Validate() error {
	if !govalidator.IsAlphanumeric(u.Username) {
		return errors.New("Invalid characters in username")
	}
	if len(u.Password) <= 0 || len(u.Password) > 255{
		return errors.New("Username length invalid")
	}
	if !govalidator.IsAlphanumeric(u.Password) {
		return errors.New("Invalid characters in password")
	}
	if len(u.Password) <= 8 || len(u.Password) > 255{
		return errors.New("Password length invalid")
	}
	return nil
}