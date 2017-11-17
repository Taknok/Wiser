//curl -H "Content-Type: application/json" --data @bodytest.json http://localhost:8082/wiser/rsu/123/cars

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

	//intelligence
	if speedInt > 10 {
		go postToRsu(t.IdVehicule, t.IdRsu)
		go postToWeb(t.IdVehicule, t.IdRsu)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/wiser/rsu/{idrsu}/cars", handleRsu)

	log.Fatal(http.ListenAndServe(":8082", router))
}
