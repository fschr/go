package main

import (
	"net/http"
	"os"
	"./controllers"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/user/", jwtMiddleware.Handler(controllers.GetUser)).Methods("GET")
	r.Handle("/user", controllers.CreateUser).Methods("POST")
	r.Handle("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	r.Handle("/login", controllers.Login).Methods("POST")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
  ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
    return []byte("MasterOfNone"), nil
  },
  SigningMethod: jwt.SigningMethodHS256,
})