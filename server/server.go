package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "modernc.org/sqlite"
)

var dbConnection *sql.DB

func main() {

	var err error
	dbConnection, err = Connect("./database.db")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /newcommit", newCommit)
	mux.HandleFunc("GET /characters/{userEmail}", getCharacter)

	s := http.Server{
		Addr:    ":1234",
		Handler: mux,
	}

	fmt.Println("Server is listening!")
	panic(s.ListenAndServe())
}

// Connect to SQLite database helper function
func Connect(path string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite", ("file:" + path))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
