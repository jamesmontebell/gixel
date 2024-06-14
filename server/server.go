package main

import (
	"database/sql"
	"net/http"

	_ "modernc.org/sqlite"
)

var dbConnection *sql.DB

func SetupServer(db *sql.DB) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /newcommit", newCommit)
	mux.HandleFunc("GET /characters/{userEmail}", getCharacter)

	s := http.Server{
		Addr:    ":1234",
		Handler: mux,
	}

	dbConnection = db
	return &s
}

func main() {

	var err error
	dbConnection, err = Connect("./database.db")
	if err != nil {
		panic(err)
	}
	panic(SetupServer(dbConnection).ListenAndServe())
}

// Connect to SQLite database helper function
func Connect(path string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite", ("file:" + path))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
