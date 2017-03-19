package controllers

import (
	"net/http"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"../services"
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
	retrievedUser := services.FindUserById(ac.session, id)
	if (models.User{}) == retrievedUser {
		http.Error(w, "User with given id does not exist", http.StatusBadRequest)
		return
	}

	payload, _ := json.Marshal(retrievedUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(payload))
}


func(ac AuthController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	newUser, err := services.CreateUser(ac.session, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} 

	payload, _ := json.Marshal(newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte(payload))
}

func(ac AuthController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")
	err := services.DeleteUserById(ac.session, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("User Deleted"))
}