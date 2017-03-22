package controllers

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/gorilla/context"
	jwt "github.com/dgrijalva/jwt-go"
	"../services/UserService"
	"../services/AuthService"
)

var GetUser = http.HandlerFunc(func(w  http.ResponseWriter, r *http.Request){
	user := context.Get(r, "user")
	username := user.(*jwt.Token).Claims.(jwt.MapClaims)["username"]

	retrievedUser, err := DataBase.FindUserByEmail(username.(string))
	if err != nil {
		http.Error(w, "User with given id does not exist", http.StatusBadRequest)
		return
	}

	payload, _ := json.Marshal(retrievedUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, err = w.Write([]byte(payload))
	if err != nil {
		http.Error(w, "Error in retrieving user", http.StatusBadRequest)
	}
})


var CreateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	newUser, err := UserService.CreateUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} 

	token, err := AuthService.GenerateJWTToken(newUser.Username)
	if err != nil {
		http.Error(w, "Error in creating token", http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	_, err = w.Write([]byte(token))
	if err != nil {
		http.Error(w, "Error in creating token", http.StatusBadRequest)
	}
})

var DeleteUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id,_ := vars["id"]
	err := UserService.DeleteUserById(id)
	if err != nil {
		http.Error(w, "Error in deleting user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	_, err = w.Write([]byte("User Deleted"))
	if err != nil {
		http.Error(w, "Error in deleting user", http.StatusBadRequest)
	}
})