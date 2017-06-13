package main

import (
	"net/http"
	"fmt"
	"html"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gocql/gocql"
)

func main(){
	//getProduct()

	getCassandra()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		r.ParseForm()
		fmt.Fprint(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}


type Product struct{
	gorm.Model
	Code string
	Price uint
}

func getProduct(){
	db, err := gorm.Open("mysql", "root:p$%57p@tcp(localhost:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "L1212", Price: 1000})

	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "L1212")

	db.Model(&product).Update("Price", 2000)

	db.Delete(&product)
}

func getCassandra(){
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
