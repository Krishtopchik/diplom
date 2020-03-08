package db

import (
	"../models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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
	err := db.QueryRow("insert into diplom (fio, topic, completion, score, queuenumber, deadline, pmid, normcontrollerid, reviewerid, chairmanid, diplomorderid, specialtyid, commissionid) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) returning id",
		diplom.Fio, diplom.Topic, diplom.Completion, diplom.Score, diplom.Queuenumber, diplom.Deadline, diplom.Pmid, diplom.Normcontrollerid, diplom.Reviewerid, diplom.Chairmanid, diplom.Diplomorderid, diplom.Specialtyid, diplom.Commissionid).Scan(&diplom.Id)
	if err != nil{
		panic(err)
	}
	return diplom, err
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
		err := rows.Scan(&p.Id, &p.Fio, &p.Topic, &p.Completion, &p.Score, &p.Queuenumber, &p.Deadline, &p.Pmid, &p.Normcontrollerid, &p.Reviewerid, &p.Chairmanid, &p.Diplomorderid, &p.Specialtyid, &p.Commissionid)
		if err != nil{
			fmt.Println(err)
			continue
		}
		diplomAll = append(diplomAll, p)
	}
	return diplomAll, err
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
	return pmAll, err
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
	return specialytyAll, err
}