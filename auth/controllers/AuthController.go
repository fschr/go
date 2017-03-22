package controllers

import (
	"net/http"
	"encoding/json"
	
	"github.com/fschr/go/auth/models"
	"github.com/fschr/go/auth/services/AuthService"
)

var Login = http.HandlerFunc(func(w  http.ResponseWriter, r *http.Request){
	requestBody := models.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Error Parsing JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	requestUser, err := DataBase.FindUserByEmail(requestBody.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := AuthService.IssueToken(&requestUser, requestBody.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	_, err = w.Write([]byte(token))
	if err != nil {
		http.Error(w, "Error in logging in", http.StatusBadRequest)
	}
})