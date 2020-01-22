package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func createConnectin() *sql.DB {
	connStr := "user=postgres password=123qwe123 dbname=test1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func GetData() {
	db := createConnectin()
	rows, err := db.Query("select * from Products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []product{}

	for rows.Next(){
		p := product{}
		err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil{
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products{
		fmt.Println(p.id, p.model, p.company, p.price)
	}
}

type product struct{
	id int
	model string
	company string
	price int
}