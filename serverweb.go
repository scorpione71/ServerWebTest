package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	//Gestore principale strumento HTTP
	http.HandleFunc("/test/", luisHandler)
	http.HandleFunc("/home", home_handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	//Senza log di errori//
	//http.ListenAndServe("localhost:9999", nil)//
}

func home_handler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./risorse_HTML/Home.html")
	page.Execute(w, "")
}

func luisHandler(response http.ResponseWriter, request *http.Request) {
	//URL.Path[1:] serve a specificare che verrà visualizzato il path a partire dal secondo carattere, quindi visualizzerà tutto ciò che verra digitato dopo "/"
	fmt.Fprintf(response, "HOME PAGE - LUIGI ORLANDO\n PATH DIGITATO [ %s ]", request.URL.Path[len("/test/"):])
}
