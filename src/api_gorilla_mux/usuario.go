package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Usuario struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type Usuarios []Usuario

var usuarios = Usuarios{
	Usuario{1, "Maria da Silva", "maria@mailtest.com", "123456"},
	Usuario{2, "João da Silva", "joao@mailtest.com", "456789"},
	Usuario{3, "Paulo da Silva", "paulo@mailtest.com", "789123"},
}

func obterUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.Encode(usuarios)
}

func obterUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, error := strconv.Atoi(params["id"])
	showError(error)

	id -= 1

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	enc.Encode(usuarios[id])
}
