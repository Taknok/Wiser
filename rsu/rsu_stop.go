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

func handleStop(rw http.ResponseWriter, req *http.Request) {

	//  DECODAGE
	decoder := json.NewDecoder(req.Body)
	var t carsInfo
	err := decoder.Decode(&t)
	
	if err != nil {
		panic(err)
	}
	log.Println("JSON = ",t)


	// SECURITE
	// if t.Action.Stop != "true" {
	// 	fmt.Println("action = true --> OK")
	// }

	var jsonStr = []byte(`{
		"apiVersion":"` + t.ApiVersion + `",
		"typeOfVehicule":"car",
		"idVehicule":"` + t.IdVehicule + `",
		"idRsu":"` + t.IdRsu + `",
		"date":"` + t.Date + `",
		"params":{
				 },
		"actions":{
				  "stop":"` + t.Action.Stop + `"
				 }
		}`)


	url := "http://localhost:8083/wiser/cars/stop"
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
	router.HandleFunc("/wiser/rsu/{idrsu}/cars/stop", handleStop)
	log.Fatal(http.ListenAndServe(":8082", router))
}
