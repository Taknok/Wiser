package main

var MAXSPEED int = 20

func intelligence(speedInt int, idVehicule string, idRsu string) {
	//intelligence
	if speedInt > MAXSPEED {
		go postToRsu(idVehicule, idRsu)
		go postToWeb(idVehicule, idRsu)
	}
}
