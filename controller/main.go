//curl -H "Content-Type: application/json" --data @bodytest.json http://localhost:8082/wiser/rsu/123/cars

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type params struct {
	Speed string
}

type carsInfo struct {
	ApiVersion     string
	TypeOfVehicule string
	IdVehicule     string
	IdRsu          string
	Date           string
	Params         params
}

func handleRsu(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t carsInfo
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	log.Println(t)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/wiser/rsu/{idrsu}/cars", handleRsu)
	log.Fatal(http.ListenAndServe(":8082", router))

	handler := rest.ResourceHandler{}
	handler.SetRoutes(
	  rest.Route{"Get","/u/:name", GetUser},
	  rest.Route{"Post","/new", NewUser},
	)
	fmt.Println("Run on 127.0.0.1:8090")
	http.ListenAndServe(":8090", &handler)
}
