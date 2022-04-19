package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,string,omitempty"`
}

type PersonWithDate struct {
	Person Person
	Date   string
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	// set header type to json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var p Person

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// add date and time to data
	currentTime := time.Now()
	dt := currentTime.Format("2006-01-02 15:04:05")
	pd := PersonWithDate{Person: p, Date: dt}

	// send json back
	json.NewEncoder(w).Encode(pd)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	// write hello back
	fmt.Fprintf(w, "hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
