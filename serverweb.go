package main

import (
	"fmt"
	"html/template"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var server http.Server

func main() {
	w := app.New()
	a := w.NewWindow("ServerWeb")
	dialogo := dialog.NewConfirm("ESCI", "Vuoi disattivare il Server ?", func(b bool) {
		if b {
			w.Quit()
		}
	}, a)
	a.SetContent(container.NewGridWithRows(2, widget.NewButton("AVVIA SERVER 1", func() {
		Multiplexer1()
	}), widget.NewButton("ESCI", func() { dialogo.Show() })))
	a.Resize(fyne.NewSize(300, 200))
	a.CenterOnScreen()
	a.ShowAndRun()

}

type Page struct {
	Title string
	Body  []byte
}

func Init_Page(title string, body string) *Page {
	return &Page{Title: title, Body: []byte(body)}
}

func Multiplexer1() {
	//http Handler che è il Multiplexer
	r := http.NewServeMux()
	r.HandleFunc("/test/", test)
	r.HandleFunc("/luis", luis)
	r.HandleFunc("/home", home)
	r.HandleFunc("/prima", prima_Pagina)
	r.HandleFunc("/mod", mod)
	r.HandleFunc("/close", close)

	server = http.Server{Addr: ":8080", Handler: r}
	server.ListenAndServe()

	//log.Fatal(http.ListenAndServe(":8080", r))

	//Gestore principale strumento HTTP
	//Creazione Server Http Multiplexer

	/*Creazione server standard
	http.HandleFunc("/test/", luisHandler)
	http.HandleFunc("/home", home_handler)
	http.ListenAndServe("localhost:9999", nil)*/
}

func close(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Chiusura - (%v)", r.URL.Path)
	server.Close()

}
func mod(w http.ResponseWriter, r *http.Request) {
	p, _ := template.ParseFiles("./risorse_HTML/modifiche.html")
	p.Execute(w, "")
	if r.Method == "POST" {
		/*
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "r.ParseForm() error =  %v", err)
				return
			}*/
		pagina := r.FormValue("pagina")
		fmt.Fprintf(w, "Da modificare : %v", pagina)

	}

}

func luis(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./risorse_HTML/Home.html")
	page.Execute(w, "")
}

func test(response http.ResponseWriter, request *http.Request) {
	//URL.Path[1:] serve a specificare che verrà visualizzato il path a partire dal secondo carattere, quindi visualizzerà tutto ciò che verra digitato dopo "/"
	fmt.Fprintf(response, "HOME PAGE - LUIGI ORLANDO\n PATH DIGITATO [ %s ]", request.URL.Path[len("/test/"):])
}

func home(w http.ResponseWriter, r *http.Request) {
	p := Init_Page("HOME PAGE", "Contenuto del body della pagina Home. Test !")
	t, _ := template.ParseFiles("./risorse_HTML/Home.html")
	t.Execute(w, p)

}
func prima_Pagina(w http.ResponseWriter, r *http.Request) {
	p := Init_Page("PRIMA PAGINA", "Contenuto del body della Prima pagina. Test !")
	t, _ := template.ParseFiles("./risorse_HTML/prima.html")
	t.Execute(w, p)
}
