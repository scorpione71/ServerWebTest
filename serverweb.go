package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Gestore principale strumento HTTP
	http.HandleFunc("/", luisHandler)
	log.Fatal(http.ListenAndServe("localhost:9999", nil))
	//Senza log di errori//
	//http.ListenAndServe("localhost:9999", nil)//
}

func luisHandler(response http.ResponseWriter, request *http.Request) {
	//URL.Path[1:] serve a specificare che verrà visualizzato il path a partire dal secondo carattere, quindi visualizzerà tutto ciò che verra digitato dopo "/"
	fmt.Fprintf(response, "ciao da %s\n", request.URL.Path[1:])
}
