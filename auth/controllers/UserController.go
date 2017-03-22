package controllers

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/gorilla/context"
	jwt "github.com/dgrijalva/jwt-go"
	"../services/UserService"
	"../services/AuthService"
	"../models"
	"../core"
)

var GetUser = http.HandlerFunc(func(w  http.ResponseWriter, r *http.Request){
	user := context.Get(r, "user")
	username := user.(*jwt.Token).Claims.(jwt.MapClaims)["username"]

	retrievedUser := core.InitDataBase().FindUserByEmail(username.(string))
	if (models.User{}) == retrievedUser {
		http.Error(w, "User with given id does not exist", http.StatusBadRequest)
		return
	}

	payload, _ := json.Marshal(retrievedUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(payload))
})


var CreateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	newUser, err := UserService.CreateUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} 

	token, err := AuthService.GenerateJWTToken(newUser.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(token))
})

var DeleteUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id,_ := vars["id"]
	err := UserService.DeleteUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("User Deleted"))
})