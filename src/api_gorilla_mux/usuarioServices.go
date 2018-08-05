package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data := usuarios

	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func getUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	showError(err)

	id -= 1

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data := usuarios[id]

	enc := json.NewEncoder(w)
	enc.Encode(data)
}
