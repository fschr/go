package main

import (
	"net/http"
	"os"

	"github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fschr/go/auth/config"
	"github.com/fschr/go/auth/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/user/", jwtMiddleware.Handler(controllers.GetUser)).Methods("GET")
	r.Handle("/user", controllers.CreateUser).Methods("POST")
	r.Handle("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	r.Handle("/login", controllers.Login).Methods("POST")

	mConfig := config.DevConfig
	http.ListenAndServe(mConfig.Env.Port, handlers.LoggingHandler(os.Stdout, r))
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("MasterOfNone"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
