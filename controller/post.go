package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func postToRsu() {
	url := "http://localhost:8080/wiser/rsu/123/cars/stop"
	fmt.Println("URL:>", url)

	test := "1.0"
	var jsonStr = []byte(`{
		"apiVersion":"` + test + `",
		"typeOfVehicule":"car",
		"idVehicule":"1234",
		"idRsu":"5678",
		"date":"2017-11-06T16:34:41:000Z",
		"params":{
			"speed":"21"
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
