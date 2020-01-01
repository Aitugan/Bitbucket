package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	tpl  *template.Template
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
}

type Course struct {
	Price              float64
	Title              string
	Description        string //
	Author             string
	Rating             float64
	RatesAmount        int
	RegisteredStudents int
	Hours              float64
	Resources          int
	GivesCertificate   bool
	Discount           float64
	PriceWithDisc      float64 //
	Image              string
	CourseId           int
}

func main() {
	fsloc := http.FileServer(http.Dir("styles"))
	mux2.Handle("/styles/", http.StripPrefix("/styles/", fsloc))

	fsscript := http.FileServer(http.Dir("scripts"))
	mux2.Handle("/scripts/", http.StripPrefix("/scripts/", fsscript))

	mux2.HandleFunc("/", index)                //main page
	mux2.HandleFunc("/admin", Admin)           //admin page
	mux2.HandleFunc("/catalog", Catalog)       //catalog page
	mux2.HandleFunc("/product", Product)       //product
	mux2.HandleFunc("/delete", DeleteElem)     //delete
	mux2.HandleFunc("/edit", UpdatePage)       //update
	mux2.HandleFunc("/create", CreateData)     //update
	mux2.HandleFunc("/createData", CreateElem) //update
	mux2.HandleFunc("/editData", UpdateElem)   //update

	http.ListenAndServe(":8080", mux2)
}

func Product(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "HELLO")
	ids, ok := r.URL.Query()["id"]

	if !ok || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	id, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	courses, err := ReadTableCourse()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	tpl = template.Must(template.ParseGlob("productPage/index.html"))

	err = tpl.Execute(w, courses[id])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func CreateElem(w http.ResponseWriter, r *http.Request) {

	sqlSt := `
	INSERT INTO CoursesForDevelopment3(Price,Title,Description,Author,Rating,RatesAmount,RegisteredStudents,Hours,Resources,GivesCertificate,Discount,PriceWithDisc,Image)
	VALUES(
	$1 ,
	$2 ,
	$3 ,
	$4 ,
	$5 ,
	$6 ,
	$7 ,
	$8 ,
	$9 ,
	$10 ,
	$11 ,
	$12 ,
	$13);
	`
	_, err := db.Exec(
		sqlSt,
		PriceED,
		TitleED,
		DescriptionED,
		AuthorED,
		RatingED,
		RatesAmountED,
		RegisteredStudentsED,
		HoursED,
		ResourcesED,
		GivesCertificateED,
		DiscountED,
		PriceWithDiscED,
		ImageED)

	if err != nil {
		fmt.Println("NOT UPDATED")
		panic(err)
	}
	tpl = template.Must(template.ParseGlob("adminPage/index.html"))
	err = tpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func CreateData(w http.ResponseWriter, r *http.Request) {
	var err error

	PriceED, err = strconv.ParseFloat(r.FormValue("Price"), 64)
	if r.FormValue("Price") != "" && err != nil {
		panic(err)
	}
	TitleED = r.FormValue("Title")
	DescriptionED = r.FormValue("Description")
	AuthorED = r.FormValue("Author")
	RatingED, err = strconv.ParseFloat(r.FormValue("Rating"), 64)
	if r.FormValue("Rating") != "" && err != nil {
		panic(err)
	}

	RatesAmountED, err = strconv.Atoi(r.FormValue("RatesAmount"))
	if r.FormValue("RatesAmount") != "" && err != nil {
		panic(err)
	}
	RegisteredStudentsED, err = strconv.Atoi(r.FormValue("RegisteredStudents"))
	if r.FormValue("RegisteredStudents") != "" && err != nil {
		panic(err)
	}
	HoursED, err = strconv.ParseFloat(r.FormValue("Hours"), 64)
	if r.FormValue("Hours") != "" && err != nil {
		panic(err)
	}
	ResourcesED, err = strconv.Atoi(r.FormValue("Resources"))
	if r.FormValue("Resources") != "" && err != nil {
		panic(err)
	}
	GivesCertificateED, err = strconv.ParseBool(r.FormValue("GivesCertificate"))
	if r.FormValue("GivesCertificate") != "" && err != nil {
		panic(err)
	}
	DiscountED, err = strconv.ParseFloat(r.FormValue("Discount"), 64)
	if r.FormValue("Discount") != "" && err != nil {
		panic(err)
	}
	PriceWithDiscED, err = strconv.ParseFloat(r.FormValue("PriceWithDisc"), 64)
	if r.FormValue("PriceWithDisc") != "" && err != nil {
		panic(err)
	}
	ImageED = r.FormValue("Image")

	CourseIDED, err = strconv.Atoi(r.FormValue("CourseID"))
	if r.FormValue("CourseID") != "" && err != nil {
		panic(err)
	}

	fmt.Println(PriceED)
	fmt.Println(TitleED)
	fmt.Println(DescriptionED)
	fmt.Println(AuthorED)
	fmt.Println(RatingED)
	fmt.Println(RatesAmountED)
	fmt.Println(RegisteredStudentsED)
	fmt.Println(HoursED)
	fmt.Println(ResourcesED)
	fmt.Println(GivesCertificateED)
	fmt.Println(DiscountED)
	fmt.Println(PriceWithDiscED)
	fmt.Println(ImageED)
	fmt.Println(CourseIDED)

	if AuthorED != "" {
		// var str strings.Builder
		// str.WriteString("/editData?id=")
		// str.WriteString(strconv.Itoa(CourseIDED))
		fmt.Println("HERE")
		http.Redirect(w, r, "/createData", http.StatusSeeOther)
	}

	tpl = template.Must(template.ParseGlob("createPage/index.html"))

	err = tpl.Execute(w, nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func DeleteElem(w http.ResponseWriter, r *http.Request) {
	ids, ok := r.URL.Query()["id"]

	if !ok || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	id, err := strconv.Atoi(ids[0])
	if err != nil {
		panic(err)
	}
	sqlStatement := `
	DELETE FROM CoursesForDevelopment3
	WHERE CourseID = $1;`
	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

var (
	PriceED              float64
	TitleED              string
	DescriptionED        string
	AuthorED             string
	RatingED             float64
	RatesAmountED        int
	RegisteredStudentsED int
	HoursED              float64
	ResourcesED          int
	GivesCertificateED   bool
	DiscountED           float64
	PriceWithDiscED      float64
	ImageED              string
	CourseIDED           int
)

func UpdatePage(w http.ResponseWriter, r *http.Request) {
	var err error

	PriceED, err = strconv.ParseFloat(r.FormValue("Price"), 64)
	if r.FormValue("Price") != "" && err != nil {
		panic(err)
	}
	TitleED = r.FormValue("Title")
	DescriptionED = r.FormValue("Description")
	AuthorED = r.FormValue("Author")
	RatingED, err = strconv.ParseFloat(r.FormValue("Rating"), 64)
	if r.FormValue("Rating") != "" && err != nil {
		panic(err)
	}

	RatesAmountED, err = strconv.Atoi(r.FormValue("RatesAmount"))
	if r.FormValue("RatesAmount") != "" && err != nil {
		panic(err)
	}
	RegisteredStudentsED, err = strconv.Atoi(r.FormValue("RegisteredStudents"))
	if r.FormValue("RegisteredStudents") != "" && err != nil {
		panic(err)
	}
	HoursED, err = strconv.ParseFloat(r.FormValue("Hours"), 64)
	if r.FormValue("Hours") != "" && err != nil {
		panic(err)
	}
	ResourcesED, err = strconv.Atoi(r.FormValue("Resources"))
	if r.FormValue("Resources") != "" && err != nil {
		panic(err)
	}
	GivesCertificateED, err = strconv.ParseBool(r.FormValue("GivesCertificate"))
	if r.FormValue("GivesCertificate") != "" && err != nil {
		panic(err)
	}
	DiscountED, err = strconv.ParseFloat(r.FormValue("Discount"), 64)
	if r.FormValue("Discount") != "" && err != nil {
		panic(err)
	}
	PriceWithDiscED, err = strconv.ParseFloat(r.FormValue("PriceWithDisc"), 64)
	if r.FormValue("PriceWithDisc") != "" && err != nil {
		panic(err)
	}
	ImageED = r.FormValue("Image")

	CourseIDED, err = strconv.Atoi(r.FormValue("CourseID"))
	if r.FormValue("CourseID") != "" && err != nil {
		panic(err)
	}

	if CourseIDED != 0 {
		var str strings.Builder
		str.WriteString("/editData?id=")
		str.WriteString(strconv.Itoa(CourseIDED))

		http.Redirect(w, r, str.String(), http.StatusSeeOther)
	}
	tpl = template.Must(template.ParseGlob("editPage/index.html"))

	err = tpl.Execute(w, nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func UpdateElem(w http.ResponseWriter, r *http.Request) {

	// rows, err := db.Query("SELECT * FROM CoursesForDevelopment3;")
	// if err != nil {
	// 	http.Error(w, http.StatusText(500), 500)
	// 	return
	// }
	// defer rows.Close()

	// courses := make([]Course, 0)
	// for rows.Next() {
	// 	crs := Course{}
	// 	err := rows.Scan(&crs.Price, &crs.Title, &crs.Description, &crs.Author, &crs.Rating, &crs.RatesAmount, &crs.RegisteredStudents, &crs.Hours, &crs.Resources, &crs.GivesCertificate, &crs.Discount, &crs.PriceWithDisc, &crs.Image, &crs.CourseId)

	// 	if err != nil {
	// 		http.Error(w, http.StatusText(500), 500)
	// 		return
	// 	}
	// 	courses = append(courses, crs)
	// }
	// if err = rows.Err(); err != nil {
	// 	http.Error(w, http.StatusText(500), 500)
	// 	return
	// }
	// fmt.Println("here")
	courses, err := ReadTableCourse()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	sqlSt := `
	UPDATE CoursesForDevelopment3
	SET
	Price = $1 ,
	Title = $2 ,
	Description = $3 ,
	Author = $4 ,
	Rating = $5 ,
	RatesAmount = $6 ,
	RegisteredStudents = $7 ,
	Hours = $8 ,
	Resources = $9 ,
	GivesCertificate = $10 ,
	Discount = $11 ,
	PriceWithDisc = $12 ,
	Image = $13
	WHERE CourseID = $14;
	`
	_, err = db.Exec(
		sqlSt,
		PriceED,
		TitleED,
		DescriptionED,
		AuthorED,
		RatingED,
		RatesAmountED,
		RegisteredStudentsED,
		HoursED,
		ResourcesED,
		GivesCertificateED,
		DiscountED,
		PriceWithDiscED,
		ImageED,
		CourseIDED)

	if err != nil {
		fmt.Println("NOT UPDATED")
		panic("NOT UPDATED")
	}
	tpl = template.Must(template.ParseGlob("editPage/index.html"))
	err = tpl.Execute(w, courses)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "HELLO")
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// rows, err := db.Query("SELECT * FROM CoursesForDevelopment3;")
	// if err != nil {
	// 	http.Error(w, http.StatusText(500), 500)
	// 	return
	// }
	// defer rows.Close()

	// courses := make([]Course, 0)
	// for rows.Next() {
	// 	crs := Course{}
	// 	err := rows.Scan(&crs.Price, &crs.Title, &crs.Description, &crs.Author, &crs.Rating, &crs.RatesAmount, &crs.RegisteredStudents, &crs.Hours, &crs.Resources, &crs.GivesCertificate, &crs.Discount, &crs.PriceWithDisc, &crs.Image, &crs.CourseId)

	// 	if err != nil {
	// 		// http.Error(w, http.StatusText(500), 500)
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	courses = append(courses, crs)
	// }
	// if err = rows.Err(); err != nil {
	// 	http.Error(w, http.StatusText(500), 500)
	// 	return
	// }
	courses, err := ReadTableCourse()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	tpl = template.Must(template.ParseGlob("mainPage/index.html"))

	err = tpl.Execute(w, courses)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Admin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	courses, err := ReadTableCourse()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// rows, err := db.Query("SELECT * FROM CoursesForDevelopment3;")
	// if err != nil {
	// 	http.Error(w, http.StatusText(500), 500)
	// 	return
	// }
	// defer rows.Close()

	// courses := make([]Course, 0)
	// for rows.Next() {
	// 	crs := Course{}
	// 	err := rows.Scan(&crs.Price, &crs.Title, &crs.Description, &crs.Author, &crs.Rating, &crs.RatesAmount, &crs.RegisteredStudents, &crs.Hours, &crs.Resources, &crs.GivesCertificate, &crs.Discount, &crs.PriceWithDisc, &crs.Image, &crs.CourseId)

	// 	if err != nil {
	// 		http.Error(w, http.StatusText(500), 500)
	// 		return
	// 	}
	// 	courses = append(courses, crs)
	// }
	// if err = rows.Err(); err != nil {
	// 	http.Error(w, http.StatusText(500), 500)
	// 	return
	// }

	tpl = template.Must(template.ParseGlob("adminPage/index.html"))

	err = tpl.Execute(w, courses)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ReadTableCourse() ([]Course, error) {
	rows, err := db.Query("SELECT * FROM CoursesForDevelopment3;")
	if err != nil {
		// http.Error(w, http.StatusText(500), 500)
		// return
		return nil, err
	}
	defer rows.Close()

	courses := make([]Course, 0)
	for rows.Next() {
		crs := Course{}
		err := rows.Scan(&crs.Price, &crs.Title, &crs.Description, &crs.Author, &crs.Rating, &crs.RatesAmount, &crs.RegisteredStudents, &crs.Hours, &crs.Resources, &crs.GivesCertificate, &crs.Discount, &crs.PriceWithDisc, &crs.Image, &crs.CourseId)

		if err != nil {
			// http.Error(w, http.StatusText(500), 500)
			// return
			return nil, err

		}
		courses = append(courses, crs)
	}
	if err = rows.Err(); err != nil {
		// http.Error(w, http.StatusText(500), 500)
		// return
		return nil, err
	}
	return courses, nil
}

func Catalog(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// rows, err := db.Query("SELECT * FROM CoursesForDevelopment3;")
	// if err != nil {
	// 	http.Error(w, http.StatusText(500), 500)
	// 	return
	// }
	// defer rows.Close()

	// courses := make([]Course, 0)
	// for rows.Next() {
	// 	crs := Course{}
	// 	err := rows.Scan(&crs.Price, &crs.Title, &crs.Description, &crs.Author, &crs.Rating, &crs.RatesAmount, &crs.RegisteredStudents, &crs.Hours, &crs.Resources, &crs.GivesCertificate, &crs.Discount, &crs.PriceWithDisc, &crs.Image, &crs.CourseId, &crs.UserID)

	// 	if err != nil {
	// 		http.Error(w, http.StatusText(500), 500)
	// 		return
	// 	}
	// 	courses = append(courses, crs)
	// }
	// if err = rows.Err(); err != nil {
	// 	http.Error(w, http.StatusText(500), 500)
	// 	return
	// }
	courses, err := ReadTableCourse()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	tpl = template.Must(template.ParseFiles("catalogPage/index.html"))

	err = tpl.Execute(w, courses)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
