package main

import (
	"database/sql"
	_ "fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InsertPerson(name string, age uint, email string, db *sql.DB) {
	// Prepare the statement for inserting data
	stmt, err := db.Prepare("INSERT INTO persons(name, age, email) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(name, age, email)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
}

func QueryAllPersons(db *sql.DB) []Person {
	rows, err := db.Query("SELECT  * from persons")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	persons := []Person{}

	for rows.Next() {
		var p Person
		var id uint
		err := rows.Scan(&id, &p.Name, &p.Age, &p.Email)
		if err != nil {
			log.Fatal(err)
		}

		persons = append(persons, p)
	}

	return persons
}

// func QueryPerson(name string, age uint, email string, db *sql.DB){
// 	rows, err := db.Query("SELECT ")
// }
