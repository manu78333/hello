package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var tpl *template.Template
var err error

type DType struct {
	Title string
	Year  int
}

var data DType

func main() {

	data.Title = "My Tilte"
	data.Year = 2023

	tpl, err = compileTemplates(".")
	if err != nil {
		os.Exit(-1)
	}
	fmt.Println(tpl.DefinedTemplates())

	http.HandleFunc("/", HelloServer)

	fmt.Println("http://localhost:3000")

	http.ListenAndServe("localhost:3000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "page1", data)
	// fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
