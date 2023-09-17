package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Car struct {
	Make  string
	Model string
	Color string
	Power string
	Year  string
	id    int8
}

func main() {
	// connStr := "user=postgres password=pgadmpwd160923 dbname=go-postgres sslmode=disable"
	connStr := "postgres://postgres:pgadmpwd160923@localhost/go-postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Could not establish connection to database")
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Connection to database failed", err)
	}

	fmt.Println("Connection to database established")

	// var rows []Car
	rows, err := db.Query("select * from public.\"Cars\"")
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("Rows closed")

	var cars []Car

	fmt.Println("These are the cars")
	for rows.Next() {
		var car Car
		if err := rows.Scan(&car.Make, &car.Model, &car.Color, &car.Power, &car.Year, &car.id); err != nil { //&car
			log.Fatal(err)
		}
		cars = append(cars, car)
	}

	fmt.Println(cars)

	err = rows.Close()
	if err != nil {
		fmt.Println("Rows could not be closed", err)
	}

	err = db.Close()
	if err != nil {
		fmt.Println("Connection could not be closed", err)
	}
	fmt.Println("Connection closed")
}
