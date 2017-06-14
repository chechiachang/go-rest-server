package main

import (
	"github.com/gocql/gocql"
	"log"
	"fmt"
)

func testCassandra(){
	cluster := gocql.NewCluster("172.17.0.3")
	cluster.Keyspace = "example"
	// create keyspace example with replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
	// create table example.tweet ( timeline text, id timeuuid, text text, primary key(timeline, id));
	cluster.Consistency = gocql.One

	session, _ := cluster.CreateSession()
	defer session.Close()

	if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
		"me", gocql.TimeUUID(), "hello world").Exec();err != nil{
		log.Fatal(err)
	}

	var id gocql.UUID
	var text string

	if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
		"me").Consistency(gocql.One).Scan(&id, &text); err != nil{
		log.Fatal(err)
	}

	fmt.Println("Tweet:", id, text)

	iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
	for iter.Scan(&id, &text){
		fmt.Println("Tweet:", id, text)
	}

	if err := iter.Close(); err != nil{
		log.Fatal(err)
	}
}

