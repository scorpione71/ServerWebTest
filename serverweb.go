package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	server      http.Server
	titoloHome  string = "HOME"
	bodyHome    string = "Body della Home"
	titoloPrima string = "PRIMA PAGINA"
	bodyPrima   string = "Body della Prima pagina"
)

func main() {
	w := app.New()
	a := w.NewWindow("ServerWeb")
	dialogo := dialog.NewConfirm("ESCI", "Vuoi disattivare il Server ?", func(b bool) {
		if b {
			w.Quit()
		}
	}, a)
	a.SetContent(container.NewGridWithRows(2, widget.NewButton("AVVIA SERVER 1", func() {
		creaFile("prova.txt")
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

func (p *Page) InitPage(title string, body string) {
	p.Title = title
	p.Body = []byte(body)
}

func (p *Page) GetPage() *Page {
	return p
}

func existFile(namefile string) bool {
	var ver bool = false
	//os.Stat serve a verificare che il file esiste//
	if _, err := os.Stat(namefile); err == nil {
		ver = true
	}
	return ver
}

func creaFile(namefile string) {
	ver := existFile(namefile)
	if !ver {
		file, err := os.Create(namefile)
		fmt.Printf("File : %v - Creato !", namefile)
		if err != nil {
			panic(err)
		}
		file.Close()
	} else {
		fmt.Printf("File : %v - Già esistente !", namefile)
	}

}

func writeFile(namefile string, p *Page) {
	final := []byte(p.Title + string(p.Body))
	err := ioutil.WriteFile("prova.txt", final, 0644)
	if err != nil {
		panic(err)
	}
}

func Multiplexer1() {
	//http Handler che è il Multiplexer
	r := http.NewServeMux()
	r.HandleFunc("/", principale)
	//r.HandleFunc("/test/", test)
	//r.HandleFunc("/home", home)
	//r.HandleFunc("/prima", primaPagina)
	//r.HandleFunc("/mod", modPage)
	//r.HandleFunc("/close", close)

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

func principale(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/mod":
		p, _ := template.ParseFiles("./risorse_HTML/modifiche.html")
		p.Execute(w, "")
		modPage(r)

	case "/home":
		var p Page
		p.InitPage(titoloHome, bodyHome)
		h := p.GetPage()
		writeFile("prova.txt", h)
		t, _ := template.ParseFiles("./risorse_HTML/Home.html")
		t.Execute(w, h)
	case "/prima":
		var p Page
		p.InitPage(titoloPrima, bodyPrima)
		h := p.GetPage()
		writeFile("prova.txt", h)
		t, _ := template.ParseFiles("./risorse_HTML/prima.html")
		t.Execute(w, h)
	case "/close":
		fmt.Printf("Chiusura - (%v) ", r.URL.Path)
		server.Close()
	}

}

/*
	func close(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Chiusura - (%v)", r.URL.Path)
		server.Close()

}
*/
func modPage(r *http.Request) {
	/*
		p, _ := template.ParseFiles("./risorse_HTML/modifiche.html")
		p.Execute(w, "")*/
	if r.Method == "POST" {
		/*
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "r.ParseForm() error =  %v", err)
				return
			}*/
		switch r.FormValue("pagina") {
		case "home":
			titoloHome = r.FormValue("titolo")
			bodyHome = r.FormValue("body")

		case "prima":
			titoloPrima = r.FormValue("titolo")
			bodyPrima = r.FormValue("body")

		}

	}

}

/*
func test(response http.ResponseWriter, request *http.Request) {
	//URL.Path[1:] serve a specificare che verrà visualizzato il path a partire dal secondo carattere, quindi visualizzerà tutto ciò che verra digitato dopo "/"
	fmt.Fprintf(response, "HOME PAGE - LUIGI ORLANDO\n PATH DIGITATO [ %s ]", request.URL.Path[len("/test/"):])
}


func home(w http.ResponseWriter, r *http.Request) {
	var p Page
	p.InitPage(titoloHome, bodyHome)
	h := p.GetPage()
	//p := InitPage("HOME PAGE", "Contenuto del body della pagina Home. Test !")
	t, _ := template.ParseFiles("./risorse_HTML/Home.html")
	t.Execute(w, h)

}

func primaPagina(w http.ResponseWriter, r *http.Request) {
	var p Page
	p.InitPage(titoloPrima, bodyPrima)
	t, _ := template.ParseFiles("./risorse_HTML/prima.html")
	t.Execute(w, p)
}*/
