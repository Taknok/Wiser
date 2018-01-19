package main

var MAXSPEED int = 30

func intelligence(speedInt int, idVehicule string, idRsu string, date string, coolant_temp string, fuelPressure string, rpm string) {
	//intelligence
	if speedInt > MAXSPEED {
		go postToRsu(idVehicule, idRsu, rsuip, date)
		//go postToWeb(idVehicule, idRsu, date)

		//insert DB
		stop := "true"
		insertData(speedInt, idVehicule, idRsu, date, stop, coolant_temp, fuelPressure, rpm)
	} else {

		//insert DB
		stop := "false"
		insertData(speedInt, idVehicule, idRsu, date, stop, coolant_temp, fuelPressure, rpm)
	}

}

func postFromWeb(idVehicule string, date string) {
	//intelligence
	go postToRsu(idVehicule, rsuid, rsuip, date)

	speedInt := 0
	coolant_temp := "0"
	fuelPressure := "0"
	rpm := "0"
	stop := "true"
	insertData(speedInt, idVehicule, rsuid, date, stop, coolant_temp, fuelPressure, rpm)

}
