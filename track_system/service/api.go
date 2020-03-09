package service

import (
	"../db"
	"../models"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

var Conf = &Config{}

func RunRest() {
	r := mux.NewRouter()
	r.HandleFunc("/api/diploms", getDiploms).Methods("GET")
	r.HandleFunc("/api/diploms/{id:[0-9]+}", getDiplom).Methods("GET")

	r.HandleFunc("/api/chairmans", getChairmans).Methods("GET")
	r.HandleFunc("/api/commissions", getCommissions).Methods("GET")
	r.HandleFunc("/api/diplomorders", getDiplomorders).Methods("GET")
	r.HandleFunc("/api/normcontrollers", getNormcontrollers).Methods("GET")
	r.HandleFunc("/api/pms", getPms).Methods("GET")
	r.HandleFunc("/api/reviewers", getReviewers).Methods("GET")
	r.HandleFunc("/api/specialtys", getSpecialtys).Methods("GET")

	r.HandleFunc("/api/diplom", createDiplom).Methods("POST")

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

func getDiploms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	diplom, err := db.GetAllDiploms()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(diplom)
}

func getDiplom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr, ok := mux.Vars(r)["id"]
	if !ok || len(idStr) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	diplom, err := db.GetDiplom(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(diplom)
}

func getChairmans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	chairman, err := db.GetChairman()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(chairman)
}

func getCommissions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	commission, err := db.GetCommissionn()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(commission)
}

func getDiplomorders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	diplomorders, err := db.GetDiplomOrder()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(diplomorders)
}

func getNormcontrollers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	normcontrollers, err := db.GetNormcontroller()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(normcontrollers)
}

func getPms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pms, err := db.GetPm()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pms)
}

func getReviewers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reviewers, err := db.GetReviewer()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reviewers)
}

func getSpecialtys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	specialtys, err := db.GetSpecialty()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(specialtys)
}

func createDiplom(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	body := json.NewDecoder(r.Body)
	var diplomModel models.Diplom
	err = body.Decode(&diplomModel)
	if err != nil {
		return
	}
	diplom, err := db.InsertDiplom(diplomModel)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(diplom)
	if err != nil {
		return
	}
}