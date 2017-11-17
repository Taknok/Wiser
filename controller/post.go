package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func postToRsu() {
	url := "http://localhost:8080/wiser/rsu/123/cars/stop"
	fmt.Println("URL:>", url)

	apiVersion := "1.0"
	idVehicule := "123"
	idRsu := "1234"
	date := time.Now().Format(time.RFC3339)
	stop := "true"

	var jsonStr = []byte(`{
		"apiVersion":"` + apiVersion + `",
		"typeOfVehicule":"car",
		"idVehicule":"` + idVehicule + `",
		"idRsu":"` + idRsu + `",
		"date":"` + date + `",
		"params":{
				 },
		"actions":{
				  "stop":"` + stop + `"
				 }
		}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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
