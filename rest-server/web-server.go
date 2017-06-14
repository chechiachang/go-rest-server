package main

import (
	"net/http"
	"fmt"
	"html"
	"log"

	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Todo")
}




