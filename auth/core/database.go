package core

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/fschr/go/auth/models"
)

type (
	DataBase struct{
		Session *mgo.Session
	}
)

var DataStore *DataBase = nil

func InitDataBase() *DataBase {
	if DataStore == nil {
		DataStore = &DataBase{getDBSession()}
	}
	return DataStore
}

func getDBSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}

func(ds *DataBase) FindUserById(id string) (retrievedUser models.User, err error) {
	session := ds.Session

	if !bson.IsObjectIdHex(id) {
		return retrievedUser, errors.New("Invalid user id")
	}

	oid := bson.ObjectIdHex(id)
	err = session.DB("AuthService").C("users").FindId(oid).One(&retrievedUser)
	return retrievedUser, err
}

func(ds *DataBase) FindUserByEmail(username string) (result models.User, err error) {
	session := ds.Session
	err = session.DB("AuthService").C("users").Find(bson.M{"username":username}).One(&result)
	return result, err
}

func(ds *DataBase) InsertUser(newUser *models.User) error {
	return ds.Session.DB("AuthService").C("users").Insert(newUser) 
}

func(ds *DataBase) DeleteUser(id string) error {
	oid := bson.ObjectIdHex(id)
	return ds.Session.DB("AuthService").C("users").RemoveId(oid)
}