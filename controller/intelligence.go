package main

import "strconv"

var MAXSPEED int = 30
var i int = 0

func intelligence(speedInt int, idVehicule string, idRsu string, date string, coolant_temp string, fuelPressure string, rpm string) {
	//intelligence

	speedString := strconv.Itoa(speedInt)
	if err != nil {
		panic(err)
	}

	if speedInt > MAXSPEED {
		go postToRsu(idVehicule, idRsu, rsuip, date)
		//go postToWeb(idVehicule, idRsu, date)

		//insert DB
		stop := "true"
		if i == 0 {
			insertData(speedString, idVehicule, idRsu, date, stop, coolant_temp, fuelPressure, rpm)
		} else {
			updateData(speedString, idVehicule, idRsu, date, stop, coolant_temp, fuelPressure, rpm)
		}

	} else {

		//insert DB
		stop := "false"
		if i == 0 {
			insertData(speedString, idVehicule, idRsu, date, stop, coolant_temp, fuelPressure, rpm)
		} else {
			updateData(speedString, idVehicule, idRsu, date, stop, coolant_temp, fuelPressure, rpm)
		}
	}
	i++

}

func postFromWeb(idVehicule string, date string) {
	//intelligence
	go postToRsu(idVehicule, rsuid, rsuip, date)

	stop := "true"

	updateDataPostFromWeb(idVehicule, date, stop)

}
