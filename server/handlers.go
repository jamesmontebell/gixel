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
	_, err = dbConnection.Exec()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
