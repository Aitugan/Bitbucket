package main

import (
	"html/template"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html"))
var crs = Course{
	3.5,
	"JAVA",
	"smth",
	5,
	100,
	100,
	100,
	100,
	true,
	100,
	1,
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", crs)
}

type Course struct {
	Price float32
	Title string
	// Description        string //
	Author             string
	Rating             float32
	RatesAmount        int
	RegisteredStudents int
	Hours              float32
	Resources          int
	GivesCertificate   bool
	Discount           float32
	// PriceWithDisc      float32 //
	id int
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	mux := http.NewServeMux()

	// Add the following two lines
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
