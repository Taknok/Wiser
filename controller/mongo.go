// MongoDB installation: https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/

package main

import (
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	IsDrop  = true
	session *mgo.Session
	c       *mgo.Collection
	r       *mgo.Collection
	err     error
)

type Cars struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	TypeOfVehicule string
	IdVehicule     string
	IdRsu          string
	Date           string
	Speed          int
	Stop           string
}

type Rsus struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	RsuID          string
	RsuIP          string
	localisation   string
	RsuNeighbourId string
	RsuNeighbourIp string
}

func mongoDatabase() {
	session, err = mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	//defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	if IsDrop {
		err = session.DB("WISER").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	// Collection WISER/ carsinfo
	c = session.DB("WISER").C("carsinfo")

	// Index  Key in min !
	indexCars := mgo.Index{
		Key:        []string{"idvehicule", "idrsu", "date"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(indexCars)
	if err != nil {
		panic(err)
	}

	// Collection WISER/ rsusinfo
	r = session.DB("WISER").C("rsusinfo")

	// Index  Key in min !
	indexRsu := mgo.Index{
		Key:        []string{"id", "rsuid"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = r.EnsureIndex(indexRsu)
	if err != nil {
		panic(err)
	}

	//insert RSU info
	err = r.Insert(&Rsus{RsuID: "1", RsuIP: "192.168.1.0", localisation: "1", RsuNeighbourId: "2/3/", RsuNeighbourIp: "192.168.1.1/192.168.1.2/"})
	if err != nil {
		panic(err)
	}
	//printDataRsus()
	//getRsusIDIP("1234")

}

func insertData(speedInt int, idVehicule string, idRsu string, date string, stop string) {
	//err := c.Insert(&Cars{TypeOfVehicule: "car", IdVehicule: "123", IdRsu: "3455", Date: "21/10/2017", Speed: 21, Stop: "false"})
	err := c.Insert(&Cars{TypeOfVehicule: "car", IdVehicule: idVehicule, IdRsu: idRsu, Date: date, Speed: speedInt, Stop: stop})
	if err != nil {
		fmt.Printf("Error 2 times same data in DB\n")
		//panic(err)
	}
	printDataCars()
}

func printDataCars() {
	var results []Cars

	err := c.Find(nil).All(&results)
	if err != nil {
		// TODO: Do something about the error
	} else {
		fmt.Println("Results All: ", results)
	}
}

func printDataRsus() {
	var results []Rsus

	err := r.Find(nil).All(&results)
	if err != nil {
		// TODO: Do something about the error
	} else {
		fmt.Println("Results All: ", results)
	}

}

func getRsusIDIP(idRsu string) (string, string) {
	var results []Rsus
	r.Find(nil).All(&results)

	for index := 0; index < len(results); index++ {
		if results[index].RsuID == idRsu {
			RsuNeighbourIdString := results[index].RsuNeighbourId
			RsuNeighbourIpString := results[index].RsuNeighbourIp
			// add id and ip of current rsu
			RsuNeighbourIdString += results[index].RsuID
			RsuNeighbourIpString += results[index].RsuIP

			return RsuNeighbourIdString, RsuNeighbourIpString
		} else {
			fmt.Printf("Problem")
		}
	}
	return "", ""

}

/*

	// Insert Datas
	err = c.Insert(&Cars{TypeOfVehicule: "car", IdVehicule: "123", IdRsu: "3455", Date: "21/10/2017", Speed: "21", Stop: "false"},
		&Cars{TypeOfVehicule: "car", IdVehicule: "123", IdRsu: "34556", Date: "21/10/2018", Speed: "21", Stop: "false"})

	if err != nil {
		panic(err)
	}
*/
/*
	// Query One
	result := Cars{}
	err = c.Find(bson.M{"idvehicule": "123", "idrsu": "3455"}).Select(bson.M{"stop": "false"}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Vehicule", result)

	// Query All
	var results []Cars
	err = c.Find(bson.M{"idvehicule": "123"}).Sort("-date").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)
*/
/*
	// Update
	colQuerier := bson.M{"name": "Ale"}
	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	err = c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}

	// Query All
	err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)
*/
