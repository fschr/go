package controllers

import (
	"net/http"
	"encoding/json"
	
	"../models"
	"../services/AuthService"
	"../core"
)

var Login = http.HandlerFunc(func(w  http.ResponseWriter, r *http.Request){
	requestBody := models.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	requestUser := core.InitDataBase().FindUserByEmail(requestBody.Username)
	if (models.User{}) == requestUser {
		http.Error(w, "User with given email does not exist", http.StatusBadRequest)
		return
	}

	token, err := AuthService.IssueToken(&requestUser, requestBody.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(token))
})