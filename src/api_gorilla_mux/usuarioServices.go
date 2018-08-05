package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.Encode(usuarios)
}

func getUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, error := strconv.Atoi(params["id"])
	showError(error)

	id -= 1

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	enc.Encode(usuarios[id])
}
