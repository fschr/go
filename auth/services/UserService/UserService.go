package UserService

import (
	"net/http"
	"errors"
	"encoding/json"

	"../../models"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

func FindUserById(session *mgo.Session, id string) models.User {
	retrievedUser := models.User{}

	if !bson.IsObjectIdHex(id) {
		return models.User{}
	}

	oid := bson.ObjectIdHex(id)
	if err := session.DB("AuthService").C("users").FindId(oid).One(&retrievedUser); err != nil{
		return models.User{}
	}
	return retrievedUser
}

func FindUserByEmail(session *mgo.Session, username string) models.User {
	result := models.User{}
	err := session.DB("AuthService").C("users").Find(bson.M{"username":username}).One(&result)
	if err != nil {
		return models.User{}
	}
	return result
}

func CreateUser(session *mgo.Session, r *http.Request) (models.User, error) {
	newUser := models.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	//Check if username is already in use
	existing := FindUserByEmail(session, newUser.Username)
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
	session.DB("AuthService").C("users").Insert(newUser) 
	return newUser, nil
}

func DeleteUserById(session *mgo.Session, id string) error{
	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid User ID")
	}

	oid := bson.ObjectIdHex(id)
	if err := session.DB("AuthService").C("users").RemoveId(oid); err != nil {
		return errors.New("Invalid User ID")
	}
	return nil
}