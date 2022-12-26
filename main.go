package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", postBook).Methods("POST")
	r.HandleFunc("/getbooks", getBooks).Methods("GET")
	r.HandleFunc("/", example).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

type Book struct {
	Name        string
	Author      string
	Publication string
}

func open() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprint("root:NikhilSharmaWe@(mysql:3306)/golibrary"))
	if err != nil {
		log.Fatalf("Error: opening connection to database %s\n", err.Error())
	}
	return db
}

func close(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("Error: closing connection %s\n", err.Error())
	}
}

func postBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	db := open()
	insertQuery, err := db.Prepare("insert into books values (?, ?, ?)")
	if err != nil {
		log.Fatalf(err.Error())
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf(err.Error())
	}

	_, err = tx.Stmt(insertQuery).Exec(book.Name, book.Author, book.Publication)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalf(err.Error())
	}
	close(db)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	db := open()
	rows, err := db.Query("select * from books")
	if err != nil {
		log.Fatalf(err.Error())
	}

	books := []Book{}
	for rows.Next() {
		var name, author, publication string
		err := rows.Scan(&name, &author, &publication)
		if err != nil {
			log.Fatalf(err.Error())
		}
		book := Book{
			Name:        name,
			Author:      author,
			Publication: publication,
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
	close(db)
}

func example(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
