package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/handlers"
	"os"
	"gopkg.in/mgo.v2"
	"./controllers"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getDBSession())
	ac := controllers.NewAuthController(getDBSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	r.POST("/login", ac.Login)

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

func getDBSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}

