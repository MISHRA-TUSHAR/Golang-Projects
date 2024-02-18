package main

import (
	"github.com/MISHRA-TUSHAR/go_with_mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/.id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/.id", uc.DeleteUser)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
