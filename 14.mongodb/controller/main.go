package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thinhbuihong/goweb/14.mongodb/controller/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
