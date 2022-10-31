package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/luigi", luisHandler)
	log.Fatal(http.ListenAndServe("localhost:9999", nil))
	//Senza log di errori//
	//http.ListenAndServe("localhost:9999", nil)//
}

func luisHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "ciao da %s\n", request.URL.Path)
}
