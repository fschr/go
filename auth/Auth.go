package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"./controllers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/user/:id", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
  ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
    return []byte("MasterOfNone"), nil
  },
  SigningMethod: jwt.SigningMethodHS256,
})