package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dataResponse := getUsuariosModel()

	enc := json.NewEncoder(w)
	enc.Encode(dataResponse)
}

func getUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dataResponse := getUsuarioModel(id)

	enc := json.NewEncoder(w)
	enc.Encode(dataResponse)
}

func postUsuario(w http.ResponseWriter, r *http.Request) {
	var usuarioRequest Usuario
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&usuarioRequest)
	showError(err)
	defer r.Body.Close()

	dataResponse := insertUsuario(usuarioRequest)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	enc.Encode(dataResponse)
}

func login(w http.ResponseWriter, r *http.Request) {
	var usuarioReq Usuario
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&usuarioReq)
	showError(err)
	defer r.Body.Close()

	var usuario Usuario

	for _, u := range usuarios {
		if u.Email == usuarioReq.Email && u.Senha == usuarioReq.Senha {
			fmt.Println(u)
			usuario = u
			break
		}
	}

	dataResponse := usuario

	w.Header().Set("Content-Type", "application/json")

	if dataResponse.Nome == "" {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	enc := json.NewEncoder(w)
	enc.Encode(dataResponse)
}
