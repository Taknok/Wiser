package main

var MAXSPEED int = 20

func intelligence(speedInt int, idVehicule string, idRsu string, date string) {
	//intelligence
	if speedInt > MAXSPEED {
		go postToRsu(idVehicule, idRsu, date)
		go postToWeb(idVehicule, idRsu, date)
	}
}
