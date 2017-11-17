package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func postToRsu(idVehicule string, idRsu string) {

	apiVersion := "1.0"
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

	idRsuInt, err := strconv.Atoi(idRsu)
	if err != nil {
		panic(err)
	}

	//handle multiple RSU
	for index := -1; index < 1; index++ {

		newIdRsuInt := idRsuInt + index
		newIdRsuString := strconv.Itoa(newIdRsuInt)

		url := "http://localhost:8080/wiser/rsu/" + newIdRsuString + "/cars/stop"
		fmt.Println("URL:>", url)

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
}

func postToWeb(idVehicule string, idRsu string) {

	apiVersion := "1.0"
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

	url := "http://localhost:8080/wiser/web/cars/stop"
	fmt.Println("URL:>", url)

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
