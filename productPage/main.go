// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"html/template"
// 	"net/http"

// 	_ "github.com/lib/pq"
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "ThreeBNet71"
// 	dbname   = "Users"
// )

// var (
// 	db *sql.DB
// 	// tpl *template.Template
// )

// type Course struct {
// 	Price              float32
// 	Title              string
// 	Description        string //
// 	Author             string
// 	Rating             float32
// 	RatesAmount        int
// 	RegisteredStudents int
// 	Hours              float32
// 	Resources          int
// 	GivesCertificate   bool
// 	Discount           float32
// 	PriceWithDisc      float32 //
// 	id                 int
// }

// func main() {

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
// 	mux := http.NewServeMux()
// 	fsloc := http.FileServer(http.Dir("assets"))
// 	fsglob := http.FileServer(http.Dir("../styles"))

// 	fmt.Println("Successfully connected!")

// 	// tpl = template.Must(template.ParseGlob("index.html"))

// 	// http.HandleFunc("/", index)

// 	// http.HandleFunc("/", CourseInfo)
// 	mux.Handle("/assets/", http.StripPrefix("/assets/", fsloc))
// 	mux.Handle("../styles/", http.StripPrefix("../styles/", fsglob))
// 	fmt.Println("HERE")

// 	mux.HandleFunc("/", index) //func(w http.ResponseWriter, r *http.Request) {
// 	// rows, err := db.Query("SELECT * FROM Courses;")
// 	// fmt.Println("HERE")

// 	// if err != nil {
// 	// 	fmt.Println("HERE")
// 	// 	http.Error(w, http.StatusText(500), 500)
// 	// 	return
// 	// }
// 	// defer rows.Close()
// 	// fmt.Println("HERE")

// 	// courses := make([]Course, 0)

// 	// for rows.Next() {
// 	// 	crs := Course{}
// 	// 	err := rows.Scan(&crs.Price, &crs.Title, &crs.Author, &crs.Rating, &crs.RatesAmount, &crs.RegisteredStudents, &crs.Hours, &crs.Resources, &crs.GivesCertificate, &crs.Discount, &crs.id)
// 	// 	if err != nil {
// 	// 		fmt.Println("HERE2")
// 	// 		fmt.Println(rows)

// 	// 		http.Error(w, http.StatusText(500), 500)
// 	// 		return
// 	// 	}

// 	// 	courses = append(courses, crs)
// 	// }

// 	// if err = rows.Err(); err != nil {
// 	// 	fmt.Println("HERE3")
// 	// 	http.Error(w, http.StatusText(500), 500)
// 	// 	return
// 	// }
// 	// fmt.Println("HERE3")

// 	// tpl.ExecuteTemplate(w, "index.html", courses)
// 	// })

// 	// mux := http.NewServeMux()
// 	// fs := http.FileServer(http.Dir("assets"))
// 	// mux.HandleFunc("/", index)

// 	http.ListenAndServe(":8080", nil)

// }

// var tpl = template.Must(template.ParseFiles("index.html"))

// func index(w http.ResponseWriter, r *http.Request) {
// 	rows, err := db.Query("SELECT * FROM Courses;")
// 	fmt.Println("HERE")

// 	if err != nil {
// 		fmt.Println("HERE")
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}
// 	defer rows.Close()
// 	fmt.Println("HERE")

// 	courses := make([]Course, 0)

// 	for rows.Next() {
// 		crs := Course{}
// 		err := rows.Scan(&crs.Price, &crs.Title, &crs.Author, &crs.Rating, &crs.RatesAmount, &crs.RegisteredStudents, &crs.Hours, &crs.Resources, &crs.GivesCertificate, &crs.Discount, &crs.id)
// 		if err != nil {
// 			fmt.Println("HERE2")
// 			fmt.Println(rows)

// 			http.Error(w, http.StatusText(500), 500)
// 			return
// 		}

// 		courses = append(courses, crs)
// 	}

// 	if err = rows.Err(); err != nil {
// 		fmt.Println("HERE3")
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}
// 	fmt.Println("HERE3")

// 	tpl.ExecuteTemplate(w, "index.html", courses)
// }

// // func index(w http.ResponseWriter, r *http.Request) {
// // 	// http.Redirect(w,r,"/books",http.StatusSeeOther)
// // 	// if r.Method != "GET" {
// // 	// 	http.Error(w, http.StatusText(500),500)
// // 	// 	return
// // 	// }

// // }

// // func CourseInfo(w http.ResponseWriter, r *http.Request) {
// // 	rows, err := db.Query("SELECT * FROM Courses;")
// // 	if err != nil {
// // 		http.Error(w, http.StatusText(500), 500)
// // 		return
// // 	}
// // 	defer rows.Close()

// // 	courses := make([]Course, 0)

// // 	for rows.Next() {
// // 		crs := Course{}
// // 		err := rows.Scan(&crs.Price)
// // 		if err != nil {
// // 			http.Error(w, http.StatusText(500), 500)
// // 			return
// // 		}
// // 		courses = append(courses, crs)
// // 	}

// // 	if err = rows.Err(); err != nil {
// // 		http.Error(w, http.StatusText(500), 500)
// // 		return
// // 	}

// // 	tpl.ExecuteTemplate(w, "", courses)
// // }

// // 	fmt.Println("Successfully connected!")

// // 	tpl = template.Must(template.ParseGlob("*.html"))

// // 	http.HandleFunc("/", index)
// // 	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
// // 		if r.Method != "GET" {
// // 			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
// // 			return
// // 		}

// // 		rows, err := db.Query("SELECT * FROM users;")
// // 		if err != nil {
// // 			http.Error(w, http.StatusText(500), 500)
// // 			return
// // 		}
// // 		defer rows.Close()

// // 		bks := make([]Book, 0)
// // 		for rows.Next() {
// // 			bk := Book{}
// // 			err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price) // order matters
// // 			if err != nil {
// // 				http.Error(w, http.StatusText(500), 500)
// // 				return
// // 			}
// // 			bks = append(bks, bk)
// // 		}
// // 		if err = rows.Err(); err != nil {
// // 			http.Error(w, http.StatusText(500), 500)
// // 			return
// // 		}

// // 		tpl.ExecuteTemplate(w, "books.html", bks)
// // 	})

// // 	// http.HandleFunc("/books", booksIndex)
// // 	// http.HandleFunc("/books/show", booksShow)
// // 	// http.HandleFunc("/books/create", booksCreateForm)
// // 	// http.HandleFunc("/books/create/process", booksCreateProcess)
// // 	// http.HandleFunc("/books/update", booksUpdateForm)
// // 	// http.HandleFunc("/books/update/process", booksUpdateProcess)
// // 	// http.HandleFunc("/books/delete/process", booksDeleteProcess)
// // 	http.ListenAndServe(":8080", nil)
// // }

// // func index(w http.ResponseWriter, r *http.Request) {
// // 	http.Redirect(w, r, "/books", http.StatusSeeOther)
// // }
