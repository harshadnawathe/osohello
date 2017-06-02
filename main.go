package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const COLOR = "blue"

type info struct {
	Version        string `json:"version,omitempty"`
	BuildTimestamp string `json:"buildTimestamp,omitempty"`
	GitCommit      string `json:"gitCommit,omitempty"`
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(&info{
		Version:        Version,
		BuildTimestamp: BuildTimestampt,
	})

	w.Write(b)
}

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("%v - Received request on %v", time.Now(), r.URL.Path)
	fmt.Fprintf(w, "Hello, %v!\n", vars["name"])
}

func colorHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v - Recieved request on %v", time.Now(), r.URL.Path)
	fmt.Fprintf(w, "Color is %v\n", COLOR)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/say-hello/{name}", sayHelloHandler).Methods("GET")
	router.HandleFunc("/info", infoHandler).Methods("GET")
	router.HandleFunc("/color", colorHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
