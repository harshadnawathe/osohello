package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("%v - Recieved request on %v", time.Now(), r.URL.Path)
	fmt.Fprintf(w, "Hello, %v!", vars["name"])
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/say-hello/{name}", sayHelloHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
