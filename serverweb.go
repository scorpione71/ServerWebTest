package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	w := app.New()
	a := w.NewWindow("ServerWeb")
	dialogo := dialog.NewConfirm("ESCI", "Vuoi disattivare il Server ?", func(b bool) {
		if b {
			w.Quit()
		}
	}, a)
	a.SetContent(container.NewGridWithRows(3, widget.NewButton("AVVIA SERVER 1", func() {
		Multiplexer1()
	}), widget.NewButton("AVVIA SERVER 2", func() {
		Multiplexer2()
	}), widget.NewButton("ESCI", func() { dialogo.Show() })))
	a.Resize(fyne.NewSize(300, 200))
	a.CenterOnScreen()
	a.ShowAndRun()

}

func Multiplexer1() {
	//http Handler che è il Multiplexer
	r := http.NewServeMux()
	r.HandleFunc("/test/", luisHandler)
	r.HandleFunc("/home", home_handler)

	log.Fatal(http.ListenAndServe(":8080", r))

	//Gestore principale strumento HTTP
	//Creazione Server Http Multiplexer

	/*Creazione server standard
	http.HandleFunc("/test/", luisHandler)
	http.HandleFunc("/home", home_handler)
	http.ListenAndServe("localhost:9999", nil)*/
}

func home_handler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./risorse_HTML/Home.html")
	page.Execute(w, "")
}

func luisHandler(response http.ResponseWriter, request *http.Request) {
	//URL.Path[1:] serve a specificare che verrà visualizzato il path a partire dal secondo carattere, quindi visualizzerà tutto ciò che verra digitato dopo "/"
	fmt.Fprintf(response, "HOME PAGE - LUIGI ORLANDO\n PATH DIGITATO [ %s ]", request.URL.Path[len("/test/"):])
}

type Page struct {
	Title string
	Body  []byte
}

func Init_Page(title string, body string) *Page {
	return &Page{Title: title, Body: []byte(body)}
}

func Multiplexer2() {
	r := http.NewServeMux()
	r.HandleFunc("/edit", edit)
	r.HandleFunc("/Home.html", home)
	log.Fatal(http.ListenAndServe(":8080", r))
}
func home(w http.ResponseWriter, r *http.Request) {
	p := Init_Page("Pagina Home Ricaricata", "Contenuto Pagina")
	t, _ := template.ParseFiles("./risorse_HTML/edit.html")
	t.Execute(w, p)

}
func edit(w http.ResponseWriter, r *http.Request) {
	p := Init_Page("Titolo Pagina Web", "Contenuto Pagina web")
	t, _ := template.ParseFiles("./risorse_HTML/edit.html")
	t.Execute(w, p)
}
