package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// API endpoint that reads a new commit created by users, calculates the characters
// new level/experience and updates the database
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
	fmt.Println(exp, level, e.Exp)
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

// API endpoint that returns a users character from the database
func getCharacter(w http.ResponseWriter, r *http.Request) {
	userEmail := r.PathValue("userEmail")
	fmt.Println(userEmail)
	var character Character

	err := dbConnection.QueryRow("SELECT user_id, user_email, name, level, experience FROM Character WHERE user_email = ?", userEmail).Scan(&character.User_id, &character.User_email, &character.Name, &character.Level, &character.Experience)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResult, err := json.Marshal(character)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResult)
}
