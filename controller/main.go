//curl -H "Content-Type: application/json" --data @bodytest.json http://localhost:8082/wiser/controller/123/cars

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type params struct {
	Speed        string
	Coolant_temp string
	Rpm          string
	FuelPressure string
}

type carsInfo struct {
	ApiVersion     string
	TypeOfVehicule string
	IdVehicule     string
	IdRsu          string
	Date           string
	Params         params
}

type stopInfo struct {
	ApiVersion     string
	TypeOfVehicule string
	IdVehicule     string
	Date           string
	Stop           actions
}

type actions struct {
	Strop string
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
	go intelligence(speedInt, t.IdVehicule, t.IdRsu, t.Date, t.Params.Coolant_temp, t.Params.FuelPressure, t.Params.Rpm)
}

func handleWebStop(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var s stopInfo
	err := decoder.Decode(&s)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	log.Println(s)

	//webstop
	go postFromWeb(s.IdVehicule, s.Date)
}

func main() {
	// Database MongoDB need to install it and launch first
	mongoDatabase()
	defer session.Close()

	//Router
	router := mux.NewRouter()
	router.HandleFunc("/wiser/controller/{idrsu}/cars", handleRsu)
	router.HandleFunc("/wiser/controller/cars/stop", handleWebStop)

	fmt.Println("listenning on port 8082")

	log.Fatal(http.ListenAndServe(":8082", router))
}
