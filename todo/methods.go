package todo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	*mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)
	u := User{}
	if err := uc.Session.DB("golang-mongo").C("Users").FindId(oid).One(&u); err != nil {
		fmt.Println("Error: ID not found")
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error Marshalling")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) PostUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := User{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Println("Error Decoding")
	}

	u.Id = bson.NewObjectId()

	if err := uc.Session.DB("golang-mongo").C("Users").Insert(u); err != nil {
		fmt.Println("Error Inserting in DB")
	}

	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println("SomeERROR")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Entry Created %s \n", data)

}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.Session.DB("golang-mongo").C("Users").RemoveId(oid); err != nil {
		fmt.Println("Error: ID not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted Entry Successfully")

}
