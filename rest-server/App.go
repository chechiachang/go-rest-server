package main

import (
	"github.com/gorilla/mux"
	"database/sql"
	"fmt"
	"log"
)

type App struct {
	Router *mux.Router
	DB *sql.DB
}

func (a *App) Initialize(user, password, dbname string){
	connString := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string){}
