package db

import (
	"../models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sort"
)

func createConnectin() *sql.DB {
	connStr := "user=postgres password=123qwe123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func InsertDiplom(diplom models.Diplom) (models.Diplom, error) {
	db := createConnectin()
	defer db.Close()
	err := db.QueryRow("insert into diplom (fio, topic, completion, score, deadline, queuenumber, pmid, normcontrollerid, reviewerid, chairmanid, diplomorderid, specialtyid, commissionid, execution, type) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) returning id",
		diplom.Fio, diplom.Topic, diplom.Completion, diplom.Score, diplom.Deadline, diplom.Queuenumber, diplom.PmId, diplom.NormcontrollerId, diplom.ReviewerId, diplom.ChairmanId, diplom.DiplomorderId, diplom.SpecialtyId, diplom.CommissionId, diplom.Execution, diplom.Type).Scan(&diplom.Id)
	if err != nil{
		panic(err)
	}
	return diplom, err
}

func DeleteDiplom(id int) error {
	db := createConnectin()
	defer db.Close()
	_, err := db.Exec("delete from diplom values where id = $1", id)
	if err != nil{
		panic(err)
	}
	return err
}

func UpdateDiplom(diplom models.Diplom) (models.Diplom, error) {
	db := createConnectin()
	defer db.Close()
	if _, err := db.Exec("update diplom set Fio = $2, Topic = $3, Completion = $4, Score = $5, Deadline = $6, Queuenumber = $7, PmId = $8, NormcontrollerId = $9, ReviewerId = $10, ChairmanId = $11, DiplomorderId = $12, SpecialtyId = $13, CommissionId = $14, execution = $15, type = $16 where id = $1",
		diplom.Id, diplom.Fio,
		diplom.Topic, diplom.Completion,
		diplom.Score, diplom.Deadline,
		diplom.Queuenumber, diplom.PmId,
		diplom.NormcontrollerId, diplom.ReviewerId,
		diplom.ChairmanId, diplom.DiplomorderId,
		diplom.SpecialtyId, diplom.CommissionId,
		diplom.Execution, diplom.Type);
	err != nil {
		return diplom, err
	}
	return diplom, nil
}

func GetAllDiploms() ([]models.Diplom, error){
	db := createConnectin()
	rows, err := db.Query("select * from diplom")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	diplomAll := []models.Diplom{}

	for rows.Next(){
		p := models.Diplom{}
		err := rows.Scan(&p.Id, &p.Fio, &p.Topic, &p.Completion, &p.Score, &p.Queuenumber, &p.Deadline, &p.PmId, &p.NormcontrollerId, &p.ReviewerId, &p.ChairmanId, &p.DiplomorderId, &p.SpecialtyId, &p.CommissionId, &p.Execution, &p.Type)
		if err != nil{
			fmt.Println(err)
			continue
		}
		diplomAll = append(diplomAll, p)
	}
	sort.Slice(diplomAll, func(i, j int) bool {
		return diplomAll[i].Id < diplomAll[j].Id
	})
	return diplomAll, err
}

func GetDiplom(id int) (models.Diplom, error){
	db := createConnectin()
	p := models.Diplom{}
	err := db.QueryRow("select * from diplom where id = $1 limit 1", id).Scan(&p.Id, &p.Fio, &p.Topic, &p.Completion, &p.Score, &p.Queuenumber, &p.Deadline, &p.PmId, &p.NormcontrollerId, &p.ReviewerId, &p.ChairmanId, &p.DiplomorderId, &p.SpecialtyId, &p.CommissionId, &p.Execution, &p.Type)
	return p, err
}

func InsertChairman(chairman models.Teacher) (models.Teacher, error) {
	db := createConnectin()
	defer db.Close()
	err := db.QueryRow("insert into chairman (fio) values ($1) returning id",
		chairman.Fio).Scan(&chairman.Id)
	if err != nil{
		panic(err)
	}
	return chairman, err
}

func GetChairman() ([]models.Teacher, error){
	db := createConnectin()
	rows, err := db.Query("select * from chairman")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	chairmanAll := []models.Teacher{}

	for rows.Next(){
		p := models.Teacher{}
		err := rows.Scan(&p.Id, &p.Fio)
		if err != nil{
			fmt.Println(err)
			continue
		}
		chairmanAll = append(chairmanAll, p)
	}
	sort.Slice(chairmanAll, func(i, j int) bool {
		return chairmanAll[i].Id < chairmanAll[j].Id
	})
	return chairmanAll, err
}

func InsertCommission(commission models.Teacher) (models.Teacher, error) {
	db := createConnectin()
	defer db.Close()
	err := db.QueryRow("insert into commission (fio) values ($1) returning id",
		commission.Fio).Scan(&commission.Id)
	if err != nil{
		panic(err)
	}
	return commission, err
}

func GetCommissionn() ([]models.Teacher, error){
	db := createConnectin()
	rows, err := db.Query("select * from commission")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	commissionAll := []models.Teacher{}

	for rows.Next(){
		p := models.Teacher{}
		err := rows.Scan(&p.Id, &p.Fio)
		if err != nil{
			fmt.Println(err)
			continue
		}
		commissionAll = append(commissionAll, p)
	}
	sort.Slice(commissionAll, func(i, j int) bool {
		return commissionAll[i].Id < commissionAll[j].Id
	})
	return commissionAll, err
}

func InsertDiplomOrder(diplomOrder models.DiplomOrder) (models.DiplomOrder, error) {
	db := createConnectin()
	defer db.Close()
	err := db.QueryRow("insert into diplomorder (name, dateorder) values ($1, $2) returning id",
		diplomOrder.Name, diplomOrder.Dateorder).Scan(&diplomOrder.Id)
	if err != nil{
		panic(err)
	}
	return diplomOrder, err
}

func GetDiplomOrder() ([]models.DiplomOrder, error){
	db := createConnectin()
	rows, err := db.Query("select * from diplomorder")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	diplomOrderAll := []models.DiplomOrder{}

	for rows.Next(){
		p := models.DiplomOrder{}
		err := rows.Scan(&p.Id, &p.Name, &p.Dateorder)
		if err != nil{
			fmt.Println(err)
			continue
		}
		diplomOrderAll = append(diplomOrderAll, p)
	}
	sort.Slice(diplomOrderAll, func(i, j int) bool {
		return diplomOrderAll[i].Id < diplomOrderAll[j].Id
	})
	return diplomOrderAll, err
}

func InsertNormcontroller(normcontroller models.Teacher) (models.Teacher, error) {
	db := createConnectin()
	defer db.Close()
	err := db.QueryRow("insert into normcontroller (fio) values ($1) returning id",
		normcontroller.Fio).Scan(&normcontroller.Id)
	if err != nil{
		panic(err)
	}
	return normcontroller, err
}

func GetNormcontroller() ([]models.Teacher, error){
	db := createConnectin()
	rows, err := db.Query("select * from normcontroller")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	normcontrollerAll := []models.Teacher{}

	for rows.Next(){
		p := models.Teacher{}
		err := rows.Scan(&p.Id, &p.Fio)
		if err != nil{
			fmt.Println(err)
			continue
		}
		normcontrollerAll = append(normcontrollerAll, p)
	}
	sort.Slice(normcontrollerAll, func(i, j int) bool {
		return normcontrollerAll[i].Id < normcontrollerAll[j].Id
	})
	return normcontrollerAll, err
}

func InsertPm(pm models.Teacher) (models.Teacher, error) {
	db := createConnectin()
	defer db.Close()
	err := db.QueryRow("insert into pm (fio) values ($1) returning id",
		pm.Fio).Scan(&pm.Id)
	if err != nil{
		panic(err)
	}
	return pm, err
}

func GetPm() ([]models.Teacher, error){
	db := createConnectin()
	rows, err := db.Query("select * from pm")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	pmAll := []models.Teacher{}

	for rows.Next(){
		p := models.Teacher{}
		err := rows.Scan(&p.Id, &p.Fio)
		if err != nil{
			fmt.Println(err)
			continue
		}
		pmAll = append(pmAll, p)
	}
	sort.Slice(pmAll, func(i, j int) bool {
		return pmAll[i].Id < pmAll[j].Id
	})
	return pmAll, err
}

func DeletePm(id int) error {
	db := createConnectin()
	defer db.Close()
	_, err := db.Exec("delete from pm values where id = $1", id)
	if err != nil{
		panic(err)
	}
	return err
}

func UpdatePm(pm models.Teacher) (models.Teacher, error) {
	db := createConnectin()
	defer db.Close()
	if _, err := db.Exec("update pm set Fio = $2 where id = $1",
		pm.Id, pm.Fio);
		err != nil {
		return pm, err
	}
	return pm, nil
}

func InsertReviewer(reviewer models.Teacher) (models.Teacher, error) {
	db := createConnectin()
	defer db.Close()
	err := db.QueryRow("insert into reviewer (fio) values ($1) returning id",
		reviewer.Fio).Scan(&reviewer.Id)
	if err != nil{
		panic(err)
	}
	return reviewer, err
}

func GetReviewer() ([]models.Teacher, error){
	db := createConnectin()
	rows, err := db.Query("select * from reviewer")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	reviewerAll := []models.Teacher{}

	for rows.Next(){
		p := models.Teacher{}
		err := rows.Scan(&p.Id, &p.Fio)
		if err != nil{
			fmt.Println(err)
			continue
		}
		reviewerAll = append(reviewerAll, p)
	}
	sort.Slice(reviewerAll, func(i, j int) bool {
		return reviewerAll[i].Id < reviewerAll[j].Id
	})
	return reviewerAll, err
}

func InsertSpecialty(specialty models.Specialyty) (models.Specialyty, error) {
	db := createConnectin()
	defer db.Close()
	err := db.QueryRow("insert into specialty (name) values ($1) returning id",
		specialty.Name).Scan(&specialty.Id)
	if err != nil{
		panic(err)
	}
	return specialty, err
}

func GetSpecialty() ([]models.Specialyty, error){
	db := createConnectin()
	rows, err := db.Query("select * from specialty")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	specialytyAll := []models.Specialyty{}

	for rows.Next(){
		p := models.Specialyty{}
		err := rows.Scan(&p.Id, &p.Name)
		if err != nil{
			fmt.Println(err)
			continue
		}
		specialytyAll = append(specialytyAll, p)
	}
	sort.Slice(specialytyAll, func(i, j int) bool {
		return specialytyAll[i].Id < specialytyAll[j].Id
	})
	return specialytyAll, err
}