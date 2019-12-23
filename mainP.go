package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

// var db *sql.DB
// var tpl *template.Template
var (
	db *sql.DB
	tpl *template.Template
	mux2 = http.NewServeMux()
)
func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:ThreeBNet71@localhost/Users?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	// tpl = template.Must(template.ParseGlob("*.html"))
	// tpl = template.Must(template.ParseFiles("index.html"))

	// tpl = template.Must(template.ParseFiles("../productPage/index.html"))

}

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}
type Course struct {
	Price              float32
	Title              string
	Description        string //
	Author             string
	Rating             float32
	RatesAmount        int
	RegisteredStudents int
	Hours              float32
	Resources          int
	GivesCertificate   bool
	Discount           float32
	PriceWithDisc      float32 //
	Image			  string
	id                 int
}
func main() {
	// mux2 := http.NewServeMux()
	fsloc := http.FileServer(http.Dir("styles"))
	mux2.Handle("/styles/", http.StripPrefix("/styles/", fsloc))

	mux2.Handle("/userPage/", http.StripPrefix("/userPage/", http.FileServer(http.Dir("userPage"))))
	mux2.Handle("/productPage/assets/", http.StripPrefix("/productPage/assets/", http.FileServer(http.Dir("productPage/assets"))))
	mux2.Handle("/mainPage/", http.StripPrefix("/mainPage/", http.FileServer(http.Dir("mainPage"))))
	mux2.Handle("/loginPage/", http.StripPrefix("/loginPage/", http.FileServer(http.Dir("loginPage"))))
	mux2.Handle("/informationPage/", http.StripPrefix("/informationPage/", http.FileServer(http.Dir("informationPage"))))
	mux2.Handle("/catalogPage/", http.StripPrefix("/catalogPage/", http.FileServer(http.Dir("catalogPage"))))
	mux2.Handle("/adminPage/", http.StripPrefix("/adminPage/", http.FileServer(http.Dir("adminPage"))))


	// fs := http.FileServer(http.Dir("assets/styles"))
	// mux2.Handle("/assets/styles/", http.StripPrefix("/assets/styles/", fs))
	fsscript := http.FileServer(http.Dir("scripts"))
	mux2.Handle("/scripts/", http.StripPrefix("/scripts/", fsscript))
	// fsscript := http.FileServer(http.Dir("assets/scripts"))
	// mux2.Handle("/assets/scripts/", http.StripPrefix("/assets/scripts/", fsscript))

	mux2.HandleFunc("/", index)
	mux2.HandleFunc("/catalog", Catalog)

	http.ListenAndServe(":8080", mux2)
}

func index(w http.ResponseWriter, r *http.Request) {
if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM CoursesForDevelopment;")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	courses := make([]Course, 0)
	for rows.Next() {
		crs := Course{}
		err := rows.Scan(&crs.Price, &crs.Title, &crs.Description, &crs.Author, &crs.Rating, &crs.RatesAmount, &crs.RegisteredStudents, &crs.Hours, &crs.Resources, &crs.GivesCertificate, &crs.Discount, &crs.PriceWithDisc, &crs.Image, &crs.id)
	
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		courses = append(courses, crs)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	tpl = template.Must(template.ParseGlob("loginPage/index.html"))

	err = tpl.Execute(w, courses)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Catalog(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM CoursesForDevelopment;")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	courses := make([]Course, 0)
	for rows.Next() {
		crs := Course{}
		err := rows.Scan(&crs.Price, &crs.Title, &crs.Description, &crs.Author, &crs.Rating, &crs.RatesAmount, &crs.RegisteredStudents, &crs.Hours, &crs.Resources, &crs.GivesCertificate, &crs.Discount, &crs.PriceWithDisc, &crs.Image, &crs.id)
	
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		courses = append(courses, crs)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	tpl = template.Must(template.ParseFiles("catalogPage/index.html"))

	err = tpl.Execute(w, courses)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}


}
