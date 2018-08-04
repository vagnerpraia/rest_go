package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var usuarios = Usuarios{
	Usuario{1, "Maria da Silva", "maria@mailtest.com", "123456"},
	Usuario{2, "João da Silva", "joao@mailtest.com", "456789"},
	Usuario{3, "Paulo da Silva", "paulo@mailtest.com", "789123"},
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", obterIndex)
	router.HandleFunc("/usuarios", obterUsuarios)
	router.HandleFunc("/usuario/{id}", obterUsuario)

	port := "3000"
	server := http.ListenAndServe(":"+port, router)

	log.Fatal(server)
}

func obterIndex(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "API Rest Go com gorilla/mux")
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

	enc := json.NewEncoder(w)
	enc.Encode(usuarios[id])
}
