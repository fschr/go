package UserService

import (
	"net/http"
	"errors"
	"encoding/json"

	"../../models"
	"../../core"
	"gopkg.in/mgo.v2/bson"
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
	existing := DB.FindUserByEmail(newUser.Username)
	if (models.User{}) != existing {
		return newUser, errors.New("User with given email already exists")
	} 

	newUser.Id = bson.NewObjectId()

	//Validate the POST request
	invalidUser := newUser.Validate()
	if invalidUser != nil {
		return newUser, invalidUser
	}

	//Insert new user into DB
	DB.InsertUser(&newUser)
	return newUser, nil
}

func DeleteUserById(id string) error{
	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid User ID")
	}

	DB.DeleteUser(id)
	return nil
}