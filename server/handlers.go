package server

import (
	"encoding/json"
	"net/http"
)

func newCommit(w http.ResponseWriter, r *http.Request) {
	var e Experience
	err := json.NewDecoder(r.Body).Decode(&e)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if e.UserEmail == "" || e.Exp == 0 {
		http.Error(w, "missing field", http.StatusBadRequest)
		return
	}

	exp, level, err := calculateLevel(e)
	if err != nil {
		http.Error(w, "missing field", http.StatusInternalServerError)
	}

	_, err = dbConnection.Exec("UPDATE Character SET level = ?, experience = ? WHERE user_email = ?", level, exp, e.UserEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func getCharacter(w http.ResponseWriter, r *http.Request) {
	userEmail := r.PathValue("userEmail")
	var character Character

	err := dbConnection.QueryRow("SELECT * FROM Character WHERE user_email=?", userEmail).Scan(&character)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
