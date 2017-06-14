package main

import (
	"log"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct{
	Name, Phone string
}

func testMongo(){
	session, err := mgo.Dial("172.17.0.3")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")

	err = c.Insert(
		&Person{"Ale", "+55 42 1244 1231"},
		&Person{"Cla", "+34 24 1231 1253"},
	)

	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
