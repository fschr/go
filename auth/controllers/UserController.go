package controllers

import (
	"net/http"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"../services/UserService"
	"../models"
)

type (
	UserController struct{
		session *mgo.Session
	}
)

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func(uc UserController) GetUser(w  http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")
	retrievedUser := UserService.FindUserById(uc.session, id)
	if (models.User{}) == retrievedUser {
		http.Error(w, "User with given id does not exist", http.StatusBadRequest)
		return
	}

	payload, _ := json.Marshal(retrievedUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(payload))
}


func(uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	newUser, err := UserService.CreateUser(uc.session, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} 

	payload, _ := json.Marshal(newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte(payload))
}

func(uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")
	err := UserService.DeleteUserById(uc.session, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("User Deleted"))
}