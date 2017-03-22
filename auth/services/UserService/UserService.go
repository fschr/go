package UserService

import (
	"net/http"
	"errors"
	"encoding/json"

	"../../models"
	"../../core"
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
)

var DB = core.InitDataBase()

func CreateUser(r *http.Request) (models.User, error) {
	newUser := models.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	//Check if username is already in use
	_, err = DB.FindUserByEmail(newUser.Username)
	if err == nil {
		return newUser, errors.New("User with given email already exists")
	} 

	newUser.Id = bson.NewObjectId()

	//Validate the POST request
	invalidUser := newUser.Validate()
	if invalidUser != nil {
		return newUser, invalidUser
	}

	//Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    newUser.Password = string(hashedPassword)

	//Insert new user into DB
	err = DB.InsertUser(&newUser)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func DeleteUserById(id string) error{
	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid User ID")
	}

	err := DB.DeleteUser(id)
	return err
}