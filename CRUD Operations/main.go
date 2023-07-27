package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Crud Operations on Server");
	r := mux.NewRouter();
	r.HandleFunc("/",serverHome).Methods("GET")

	// running a server
	// to log the errors coming in, we use a log module
	log.Fatal(http.ListenAndServe(":4000",r));
}


func serverHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome in CRUD app</h1>"));
}