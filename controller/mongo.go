package main

import (
  "fmt"
  "github.com/ant0ine/go-json-rest/rest"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "net/http"
)

// mongodb://mongo:27017/wiser
const (
	addr      = "mongo:27017"
	database   = "wiser"
	username   = "admin"
	password   = "youPassword"
	collection = "people"
)

type Person struct {
  UserName string
  FullName string
  Phone string
  Age int
  Gender string
}

type Err struct {
  ErrCode int
  ErrMsg string
}

func FindUser(user string) (Person,error) {
  session, err := mgo.Dial(addr)
  p := Person{}
  if err != nil {
    return p, err
  }
  defer session.Close()
  c := session.DB(database).C(collection)
  if err = c.Find(bson.M{"username":user}).One(&p); err != nil {
    return p, err
  }
  return p, nil
}

func AddUser(p Person) (err_msg Err) {
  session, err := mgo.Dial(addr)
  if err != nil {
    err_msg = Err{1, "connection failed"}
  }
  defer session.Close()
  c := session.DB(database).C(collection)
  tp := Person{}  // temp person :P
  // if not found add new user
  // else return error user already exist
  if err := c.Find(bson.M{"username":&p.UserName}).One(&tp);err != nil {
    err_msg = Err{0, "done"}
    if err := c.Insert(&p); err != nil {
      err_msg = Err{1, "insert failed"}
    }
  } else {
    err_msg = Err{1, "user already exist"}
  }
  return
}

func GetUser(w *rest.ResponseWriter, req *rest.Request) {
  name, err := FindUser(req.PathParam("name"))
  if err != nil {
    err_msg := Err{1,"Not found"}
    w.WriteJson(&err_msg)
  } else {
    w.WriteJson(&name)
  }
}

func NewUser(w *rest.ResponseWriter, req *rest.Request) {
  p := Person{}
  err_msg := Err{}
  if err := req.DecodeJsonPayload(&p); err != nil {
    rest.Error(w, err.Error(), http.StatusInternalServerError)
    return
  } else {
    err_msg = AddUser(p)
  }
  w.WriteJson(&err_msg)
}

