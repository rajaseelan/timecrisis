package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Time Crisis Server - saving you, one request at a time ;)")
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	port, present := os.LookupEnv("TIMECRISIS_PORT")
	if !present {
		port = ":4000"
	} else {
		port = ":" + port
	}

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
