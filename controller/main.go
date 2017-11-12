package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Car struct {
	typeOfVehicule string  `json:"typeOfVehicule,omitempty"`
	idVehicule     string  `json:"id,omitempty"`
	date           string  `json:"date,omitempty"`
	Params         *Params `json:"params,omitempty"`
}
type Params struct {
	speed string `json:"speed,omitempty"`
}

var cars []Car

// our main function
func main() {

	cars = append(cars, Car{typeOfVehicule: "car", idVehicule: "1", date: time.Now().Format(time.RFC3339), Params: &Params{speed: "20"}})
	cars = append(cars, Car{typeOfVehicule: "car", idVehicule: "2", date: time.Now().Format(time.RFC3339), Params: &Params{speed: "20"}})

	router := mux.NewRouter()
	router.HandleFunc("/wiser/rsu/{idrsu}/cars", GetRsu).Methods("GET")
	router.HandleFunc("/wiser/rsu/{idrsu}/cars/stop", PostAction).Methods("POST")
	router.HandleFunc("/wiser/web/cars/stop", PostInformation).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetRsu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range cars {
		if item.idVehicule == params["idrsu"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Car{})
}

func PostAction(w http.ResponseWriter, r *http.Request)      {}
func PostInformation(w http.ResponseWriter, r *http.Request) {}
