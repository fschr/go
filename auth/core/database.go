package core

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"../models"
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

func(ds *DataBase) FindUserById(id string) models.User {
	session := ds.Session
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

func(ds *DataBase) FindUserByEmail(username string) models.User {
	session := ds.Session
	result := models.User{}
	err := session.DB("AuthService").C("users").Find(bson.M{"username":username}).One(&result)
	if err != nil {
		return models.User{}
	}
	return result
}

func(ds *DataBase) InsertUser(newUser *models.User) {
	ds.Session.DB("AuthService").C("users").Insert(newUser) 
}

func(ds *DataBase) DeleteUser(id string) {
	oid := bson.ObjectIdHex(id)
	ds.Session.DB("AuthService").C("users").RemoveId(oid)
}