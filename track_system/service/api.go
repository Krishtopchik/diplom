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
	r.HandleFunc("/api/diploms/{id:[0-9]+}", deleteDiplom).Methods("DELETE")

	r.HandleFunc("/api/chairmans", getChairmans).Methods("GET")
	r.HandleFunc("/api/chairmans", createChairmans).Methods("POST")
	r.HandleFunc("/api/chairmans/{id:[0-9]+}", deleteChairmans).Methods("DELETE")
	r.HandleFunc("/api/chairmans", updateChairmans).Methods("PUT")

	r.HandleFunc("/api/commissions", getCommissions).Methods("GET")
	r.HandleFunc("/api/commissions", createCommissions).Methods("POST")
	r.HandleFunc("/api/commissions/{id:[0-9]+}", deleteCommissions).Methods("DELETE")
	r.HandleFunc("/api/commissions", updateCommissions).Methods("PUT")

	r.HandleFunc("/api/diplomorders", getDiplomorders).Methods("GET")
	r.HandleFunc("/api/diplomorders", createDiplomorders).Methods("POST")
	r.HandleFunc("/api/diplomorders/{id:[0-9]+}", deleteDiplomorders).Methods("DELETE")
	r.HandleFunc("/api/diplomorders", updateDiplomorders).Methods("PUT")

	r.HandleFunc("/api/normcontrollers", getNormcontrollers).Methods("GET")
	r.HandleFunc("/api/normcontrollers", createNormcontrollers).Methods("POST")
	r.HandleFunc("/api/normcontrollers/{id:[0-9]+}", deleteNormcontrollers).Methods("DELETE")
	r.HandleFunc("/api/normcontrollers", updateNormcontrollers).Methods("PUT")

	r.HandleFunc("/api/pms", getPms).Methods("GET")
	r.HandleFunc("/api/pms", createPm).Methods("POST")
	r.HandleFunc("/api/pms/{id:[0-9]+}", deletePm).Methods("DELETE")
	r.HandleFunc("/api/pms", updatePm).Methods("PUT")

	r.HandleFunc("/api/reviewers", getReviewers).Methods("GET")
	r.HandleFunc("/api/reviewers", createReviewers).Methods("POST")
	r.HandleFunc("/api/reviewers/{id:[0-9]+}", deleteReviewers).Methods("DELETE")
	r.HandleFunc("/api/reviewers", updateReviewers).Methods("PUT")
	r.HandleFunc("/api/specialtys", getSpecialtys).Methods("GET")

	r.HandleFunc("/api/diploms", createDiplom).Methods("POST")
	r.HandleFunc("/api/diploms", updateDiplom).Methods("PUT")

	r.HandleFunc("/api/doc", createDoc).Methods("get")

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

func deleteDiplom(w http.ResponseWriter, r *http.Request) {
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
	err = db.DeleteDiplom(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(err)
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

func createChairmans(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	body := json.NewDecoder(r.Body)
	var teacherModel models.Teacher
	err = body.Decode(&teacherModel)
	if err != nil {
		return
	}
	teacher, err := db.InsertChairman(teacherModel)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(teacher)
	if err != nil {
		return
	}
}

func deleteChairmans(w http.ResponseWriter, r *http.Request) {
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
	err = db.DeleteChairman(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(err)
}

func updateChairmans(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	body := json.NewDecoder(r.Body)
	var chairmanModel models.Teacher
	err = body.Decode(&chairmanModel)
	if err != nil {
		return
	}
	teacher, err := db.UpdateChairman(chairmanModel)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(teacher)
	if err != nil {
		return
	}
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

func createCommissions(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	body := json.NewDecoder(r.Body)
	var teacherModel models.Teacher
	err = body.Decode(&teacherModel)
	if err != nil {
		return
	}
	teacher, err := db.InsertCommission(teacherModel)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(teacher)
	if err != nil {
		return
	}
}

func deleteCommissions(w http.ResponseWriter, r *http.Request) {
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
	err = db.DeleteCommissionn(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(err)
}

func updateCommissions(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	body := json.NewDecoder(r.Body)
	var teacherModel models.Teacher
	err = body.Decode(&teacherModel)
	if err != nil {
		return
	}
	teacher, err := db.UpdateCommissionn(teacherModel)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(teacher)
	if err != nil {
		return
	}
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

func createDiplomorders(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	body := json.NewDecoder(r.Body)
	var diplomOrderModel models.DiplomOrder
	err = body.Decode(&diplomOrderModel)
	if err != nil {
		return
	}
	diplomOrder, err := db.InsertDiplomOrder(diplomOrderModel)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(diplomOrder)
	if err != nil {
		return
	}
}

func deleteDiplomorders(w http.ResponseWriter, r *http.Request) {
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
	err = db.DeleteDiplomOrder(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(err)
}

func updateDiplomorders(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	body := json.NewDecoder(r.Body)
	var diplomOrderModel models.DiplomOrder
	err = body.Decode(&diplomOrderModel)
	if err != nil {
		return
	}
	diplomOrder, err := db.UpdateDiplomOrder(diplomOrderModel)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(diplomOrder)
	if err != nil {
		return
	}
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

func createNormcontrollers(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	body := json.NewDecoder(r.Body)
	var teacherModel models.Teacher
	err = body.Decode(&teacherModel)
	if err != nil {
		return
	}
	teacher, err := db.InsertNormcontroller(teacherModel)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(teacher)
	if err != nil {
		return
	}
}

func deleteNormcontrollers(w http.ResponseWriter, r *http.Request) {
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
	err = db.DeleteNormcontroller(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(err)
}

func updateNormcontrollers(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	body := json.NewDecoder(r.Body)
	var teacherModel models.Teacher
	err = body.Decode(&teacherModel)
	if err != nil {
		return
	}
	teacher, err := db.UpdateNormcontroller(teacherModel)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(teacher)
	if err != nil {
		return
	}
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

func createPm(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	body := json.NewDecoder(r.Body)
	var teacherModel models.Teacher
	err = body.Decode(&teacherModel)
	if err != nil {
		return
	}
	teacher, err := db.InsertPm(teacherModel)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(teacher)
	if err != nil {
		return
	}
}

func deletePm(w http.ResponseWriter, r *http.Request) {
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
	err = db.DeletePm(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(err)
}

func updatePm(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	body := json.NewDecoder(r.Body)
	var pmModel models.Teacher
	err = body.Decode(&pmModel)
	if err != nil {
		return
	}
	pm, err := db.UpdatePm(pmModel)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(pm)
	if err != nil {
		return
	}
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

func createReviewers(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	body := json.NewDecoder(r.Body)
	var teacherModel models.Teacher
	err = body.Decode(&teacherModel)
	if err != nil {
		return
	}
	teacher, err := db.InsertReviewer(teacherModel)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(teacher)
	if err != nil {
		return
	}
}

func deleteReviewers(w http.ResponseWriter, r *http.Request) {
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
	err = db.DeleteReviewer(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(err)
}

func updateReviewers(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	body := json.NewDecoder(r.Body)
	var teacherModel models.Teacher
	err = body.Decode(&teacherModel)
	if err != nil {
		return
	}
	teacher, err := db.UpdateReviewer(teacherModel)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(teacher)
	if err != nil {
		return
	}
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

func updateDiplom(w http.ResponseWriter, r *http.Request) {
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
	diplom, err := db.UpdateDiplom(diplomModel)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(diplom)
	if err != nil {
		return
	}
}

func createDoc(w http.ResponseWriter, r *http.Request)  {

}