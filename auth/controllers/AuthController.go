package controllers

import (
	"net/http"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"../models"
)

type (
	AuthController struct{
		session *mgo.Session
	}
)

func NewAuthController(s *mgo.Session) *AuthController {
	return &AuthController{s}
}

func(ac AuthController) GetUser(w  http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User ID is invalid"))
		return
	}

	oid := bson.ObjectIdHex(id)
	retrievedUser := models.User{}

	if err := ac.session.DB("AuthService").C("users").FindId(oid).One(&retrievedUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User with given ID not found"))
		return
	}

	payload, _ := json.Marshal(retrievedUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(payload))
}


func(ac AuthController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	newUser := models.User{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newUser)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	//Check if username is already in use
	existing := findUserByUsername(ac.session, newUser.Username)
	if (models.User{}) != existing {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	} 

	newUser.Id = bson.NewObjectId()

	//Validate the POST request
	invalidUser := newUser.Validate()
	if invalidUser != nil {
		http.Error(w, invalidUser.Error(), http.StatusBadRequest)
		return
	}

	//Insert new user into DB
	ac.session.DB("AuthService").C("users").Insert(newUser) 

	payload, _ := json.Marshal(newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte(payload))
}

func(ac AuthController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User ID is invalid"))
		return
	}

	oid := bson.ObjectIdHex(id)
	if err := ac.session.DB("AuthService").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User with given ID not found"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("User Deleted"))
}

func findUserByUsername(session *mgo.Session, username string) models.User {
	result := models.User{}
	err := session.DB("AuthService").C("users").Find(bson.M{"username":username}).One(&result)
	if err != nil {
		return models.User{}
	}
	return result
}