//curl -H "Content-Type: application/json" --data @bodytest.json http://localhost:8082/wiser/rsu/123/cars
// MongoDB installation: https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

	//Convert string to int
	speedInt, err := strconv.Atoi(t.Params.Speed)
	if err != nil {
		panic(err)
	}

	//Intelligence algo
	go intelligence(speedInt, t.IdVehicule, t.IdRsu)
}

func main() {
	// Database MongoDB need to install it and launch first
	mongoDatabase()

	//Router
	router := mux.NewRouter()
	router.HandleFunc("/wiser/rsu/{idrsu}/cars", handleRsu)

	log.Fatal(http.ListenAndServe(":8082", router))
}
