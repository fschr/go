package controllers

import (
	"net/http"
	"encoding/json"

	"github.com/fschr/go/auth/services/UserService"
	"github.com/fschr/go/auth/services/AuthService"
	"github.com/gorilla/mux"
	"github.com/gorilla/context"
	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/Sirupsen/logrus"
)

var GetUser = http.HandlerFunc(func(w  http.ResponseWriter, r *http.Request){
	user := context.Get(r, "user")
	username := user.(*jwt.Token).Claims.(jwt.MapClaims)["username"]

	retrievedUser, err := DataBase.FindUserByEmail(username.(string))
	if err != nil {
		log.Error(err)
		http.Error(w, "InvalidUserError", http.StatusBadRequest)
		return
	}

	payload, _ := json.Marshal(retrievedUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, err = w.Write([]byte(payload))
	if err != nil {
		log.Error(err)
		http.Error(w, "ResponseError", http.StatusInternalServerError)
	}
})


var CreateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	newUser, creationErr := UserService.CreateUser(r)
	if creationErr != nil {
		log.Error(creationErr)
		http.Error(w, creationErr.Error(), http.StatusBadRequest)
		return
	} 

	token, err := AuthService.GenerateJWTToken(newUser.Username)
	if err != nil {
		log.Error(err)
		http.Error(w, "TokenCreationError", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	_, err = w.Write([]byte(token))
	if err != nil {
		log.Error(err)
		http.Error(w, "ResponseError", http.StatusInternalServerError)
	}
})

var DeleteUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id,_ := vars["id"]
	err := UserService.DeleteUserById(id)
	if err != nil {
		log.Error(err)
		http.Error(w, "UserDeletionError", http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	_, err = w.Write([]byte("User Deleted"))
	if err != nil {
		log.Error(err)
		http.Error(w, "ResponseError", http.StatusInternalServerError)
	}
})