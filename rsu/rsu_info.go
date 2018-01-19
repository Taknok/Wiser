//curl -H "Content-Type: application/json" --data @bodytest.json http://localhost:8082/wiser/rsu/123/cars

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/gorilla/mux"
)

type params struct {
	Speed string
	Coolant_temp string
	Rpm string
	FuelPressure string	
}

type action struct {
	Stop string
}

type carsInfo struct {
	ApiVersion     string
	TypeOfVehicule string
	IdVehicule     string
	IdRsu          string
	Date           string
	Params         params
	Action		   action
}

func forward2Controller(rw http.ResponseWriter, req *http.Request) {

	//  DECODAGE
	decoder := json.NewDecoder(req.Body)
	var t carsInfo
	err := decoder.Decode(&t)
	
	if err != nil {
		panic(err)
	}
	log.Println("JSON = ",t)


	// SECURITE


	// MODIFIE JSON
	MyIdRsu := "MyIdRsu"
	var jsonStr = []byte(`{
		"apiVersion":"` + t.ApiVersion + `",
		"typeOfVehicule":"car",
		"idVehicule":"` + t.IdVehicule + `",
		"idRsu":"` + MyIdRsu + `",
		"date":"` + t.Date + `",
		"params":{
				"speed":"` + t.Params.Speed + `",
				"coolant_temp": "` + t.Params.Coolant_temp + `",
				"rpm":"` + t.Params.Rpm + `",
				"fuelPressure":"` + t.Params.FuelPressure + `"
				 },
		"actions":{
				  "stop":"false"
				 }
		}`)


	url := "http://192.168.43.35:8082/wiser/controller/" +MyIdRsu+"/cars"
	fmt.Println("URL:>", url)

	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func main() {
	log.Println(" -- MAIN --")	
	router := mux.NewRouter()
	router.HandleFunc("/wiser/cars", forward2Controller)
	log.Fatal(http.ListenAndServe(":8083", router))
}
