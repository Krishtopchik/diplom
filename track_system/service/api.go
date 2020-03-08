package service

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Conf = &Config{}
var books []Book

func RunRest() {
	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
	books = append(books, Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(Conf.StaticPath)))

	log.Printf("starting REST server on %s", Conf.ListenPort)

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	err := http.ListenAndServe(
		Conf.ListenPort,
		handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(r))
	if err != nil {
		log.Fatal(err)
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//chairman, err := db.InsertChairman(models.Chairman{Fio: "qwe3"})
//fmt.Println(chairman, err)
//commission, err := db.InsertCommission(models.Teacher{Fio: "commission"})
//fmt.Println(commission, err)
//diplomorder, err := db.InsertDiplomOrder(models.DiplomOrder{Name: "diplomorder1", Dateorder:"diplomorder1"})
//fmt.Println(diplomorder, err)
//normcontroller, err := db.InsertNormcontroller(models.Teacher{Fio: "normcontroller1"})
//fmt.Println(normcontroller, err)
//pm, err := db.InsertPm(models.Teacher{Fio: "pm1"})
//fmt.Println(pm, err)
//reviewer, err := db.InsertReviewer(models.Teacher{Fio: "reviewer1"})
//fmt.Println(reviewer, err)
//specialty, err := db.InsertSpecialty(models.Specialyty{Name: "specialty1"})
//fmt.Println(specialty, err)

//diplom, err := db.InsertDiplom(models.Diplom{
//	Fio:"qwe",
//	Topic:"qwe",
//	Completion:0,
//	Score:1,
//	Queuenumber:1,
//	Deadline:"123",
//	Pmid:1,
//	Normcontrollerid:1,
//	Reviewerid:1,
//	Chairmanid:1,
//	Diplomorderid:1,
//	Specialtyid:1,
//	Commissionid:1,
//})
//fmt.Println(diplom, err)

//diplom, err := db.GetAllDiploms()
//fmt.Println(diplom, err)

type Book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}