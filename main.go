package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const COLOR = "green"

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("%v - Received request on %v", time.Now(), r.URL.Path)
	fmt.Fprintf(w, "Hello, %v!\n", vars["name"])
}

func colorHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v - Recieved request on %v", time.Now(), r.URL.Path)
	fmt.Fprintf(w, "Color is %v\n", COLOR)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v - Recieved request on %v", time.Now(), r.URL.Path)
	fmt.Fprintf(w, "version: %v\n", Version)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/say-hello/{name}", sayHelloHandler).Methods("GET")
	router.HandleFunc("/color", colorHandler).Methods("GET")
	router.HandleFunc("/version", versionHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
