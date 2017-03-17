package controllers

import (
	"net/http"
	"encoding/json"
	"log"

	"github.com/julienschmidt/httprouter"
	"../models"
)

type (
	UserController struct{}
)

func NewUserController() *UserController {
	return &UserController{}
}

func(uc UserController) GetUser(w  http.ResponseWriter, r *http.Request, p httprouter.Params){
	retrievedUser := models.User {
		Name: "Tommy Yu",
		Id: 0,
		Rank: 100,
	}

	payload, _ := json.Marshal(retrievedUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(payload))
}


func(uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	var newUser models.User

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newUser)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	log.Println(newUser)

	json.Marshal(newUser)
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
}

func(uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	w.WriteHeader(200)
}