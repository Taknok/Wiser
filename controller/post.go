package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func postToRsu(idVehicule string, idRsu string, ipRsu string, date string) {

	apiVersion := "1.0"
	//date := time.Now().Format(time.RFC3339)
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

	if err != nil {
		panic(err)
	}

	//handle multiple RSU
	//rsuIds, rsuIps := getRsusIDIP(idRsu)

	//rsuIdsstringsplit := strings.Split(rsuIds, "/")
	//rsuIpsstringsplit := strings.Split(rsuIps, "/")
	//fmt.Println(rsuIdsstringsplit)
	//fmt.Println(rsuIpsstringsplit)

	//handle multiple RSU
	//for index := 0; index < len(rsuIdsstringsplit); index++ {

	//url := "http://" + rsuIpsstringsplit[index] + ":8080/wiser/rsu/" + rsuIdsstringsplit[index] + "/cars/stop"
	url := "http://" + ipRsu + "/wiser/rsu/" + idRsu + "/cars/stop"
	//fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error try send to RSU\n")
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	//}
}

func postToWeb(idVehicule string, idRsu string, date string) {

	apiVersion := "1.0"
	//date := time.Now().Format(time.RFC3339)
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

	url := webip + "/wiser/web/cars/stop"
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error try to send to web server")
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
