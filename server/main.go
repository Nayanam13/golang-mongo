package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nayanam/golang-mongo/todo"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := todo.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.PostUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	if err := http.ListenAndServe("localhost:9000", r); err == nil {
		fmt.Println("Server Started")
		return
	}
	fmt.Println("Server Started")

}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
