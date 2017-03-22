package core

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"../models"
	"errors"
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

func(ds *DataBase) FindUserById(id string) (models.User, error) {
	session := ds.Session
	retrievedUser := models.User{}

	if !bson.IsObjectIdHex(id) {
		return models.User{}, errors.New("Invalid user id")
	}

	oid := bson.ObjectIdHex(id)
	if err := session.DB("AuthService").C("users").FindId(oid).One(&retrievedUser); err != nil{
		return models.User{}, err
	}
	return retrievedUser, nil
}

func(ds *DataBase) FindUserByEmail(username string) (models.User, error) {
	session := ds.Session
	result := models.User{}
	err := session.DB("AuthService").C("users").Find(bson.M{"username":username}).One(&result)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}

func(ds *DataBase) InsertUser(newUser *models.User) error {
	return ds.Session.DB("AuthService").C("users").Insert(newUser) 
}

func(ds *DataBase) DeleteUser(id string) error {
	oid := bson.ObjectIdHex(id)
	return ds.Session.DB("AuthService").C("users").RemoveId(oid)
}