package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	db, err := sql.Open("postgres", "postgres://bond:password@postgres:5432/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
