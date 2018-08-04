package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/rota1", rota1)
	router.HandleFunc("/usuarios", get_usuarios)
	router.HandleFunc("/usuario/{id}", get_usuario)

	port := "3000"
	server := http.ListenAndServe(":"+port, router)

	log.Fatal(server)
}

func index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "API Rest Go com gorilla/mux")
}

func rota1(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Rota 1")
}

func get_usuarios(w http.ResponseWriter, r *http.Request) {
	usuarios := Usuarios{
		Usuario{1, "Maria da Silva", "maria@mailtest.com", "123456"},
		Usuario{2, "João da Silva", "joao@mailtest.com", "456789"},
		Usuario{3, "Paulo da Silva", "paulo@mailtest.com", "789123"},
	}

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.Encode(usuarios)
}

func get_usuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, error := strconv.Atoi(params["id"])
	showError(error)

	id -= 1

	usuarios := Usuarios{
		Usuario{1, "João da Silva", "joao@mailtest.com", "456789"},
		Usuario{2, "Paulo da Silva", "paulo@mailtest.com", "789123"},
		Usuario{3, "Paulo da Silva", "paulo@mailtest.com", "789123"},
	}

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.Encode(usuarios[id])
}
