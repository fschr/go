package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"
	"../models"
	"../services/AuthService"
	"../services/UserService"
)

type (
	AuthController struct{
		session *mgo.Session
	}
)

func NewAuthController(s *mgo.Session) *AuthController {
	return &AuthController{s}
}

func(ac AuthController) Login(w  http.ResponseWriter, r *http.Request, p httprouter.Params){
	requestBody := models.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	requestUser := UserService.FindUserByEmail(ac.session, requestBody.Username)
	if (models.User{}) == requestUser {
		http.Error(w, "User with given email does not exist", http.StatusBadRequest)
		return
	}

	token, err := AuthService.IssueToken(&requestUser, requestBody.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println([]byte(token))

	w.WriteHeader(200)
	w.Write([]byte(token))
}