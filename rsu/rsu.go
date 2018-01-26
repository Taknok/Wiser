//curl -H "Content-Type: application/json" --data @bodytest.json http://localhost:8082/wiser/rsu/123/cars

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"

	"github.com/gorilla/mux"
)

var webip string = "127.0.0.1:8082"
var obuip string = "192.168.137.144:8081"
var MyIdRsu string = "1"

type params_info struct {
	Speed        string
	Coolant_temp string
	Rpm          string
	FuelPressure string
}

type params_stop struct {
	Speed string
}

type actions_stop struct {
	Stop string
}

type carsInfo struct {
	ApiVersion     string
	TypeOfVehicule string
	IdVehicule     string
	Date           string
	Params         params_info
}

type carsStop struct {
	ApiVersion     string
	TypeOfVehicule string
	IdVehicule     string
	Date           string
	Params         params_stop
	Actions        actions_stop
}

func forward2Controller(rw http.ResponseWriter, req *http.Request) {
	color.Green("Received Data from OBU, send it to controller")

	//  DECODAGE
	decoder := json.NewDecoder(req.Body)
	var t carsInfo
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}
	log.Println("JSON = ", t)

	// SECURITE

	// MODIFIE JSON
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

	url := "http://" + webip + "/wiser/controller/" + MyIdRsu + "/cars"
	//fmt.Println("URL:>", url)

	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func handleStop(rw http.ResponseWriter, req *http.Request) {
	color.Red("STOP received from controller")
	//  DECODAGE
	decoder := json.NewDecoder(req.Body)
	var t carsStop
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}
	log.Println("JSON = ", t)

	// SECURITE
	// if t.Action.Stop != "true" {
	// 	fmt.Println("action = true --> OK")
	// }

	var jsonStr = []byte(`{
		"apiVersion":"` + t.ApiVersion + `",
		"typeOfVehicule":"car",
		"idVehicule":"` + t.IdVehicule + `",
		"idRsu":"` + MyIdRsu + `",
		"date":"` + t.Date + `",
		"params":{
				 },
		"actions":{
				  "stop":"` + t.Actions.Stop + `"
				 }
		}`)

	url := "http://" + obuip + "/wiser/cars/stop"
	//fmt.Println("URL:>", url)

	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}

func main() {
	log.Println(" -- MAIN --")
	router := mux.NewRouter()
	router.HandleFunc("/wiser/rsu", forward2Controller)
	router.HandleFunc("/wiser/rsu/{idrsu}/cars/stop", handleStop)
	log.Fatal(http.ListenAndServe(":8083", router))
}
