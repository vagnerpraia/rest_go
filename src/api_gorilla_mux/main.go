package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/rota1", rota1)
	router.HandleFunc("/usuarios", usuarios)
	router.HandleFunc("/usuario/{id}", usuario)

	port := "3000"
	server := http.ListenAndServe(":"+port, router)
	fmt.Println("Servidor rodando na porta " + port + ".")
	log.Fatal(server)
}

func index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "API Go com gorilla/mux")
}

func rota1(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Rota 1")
}

func usuarios(response http.ResponseWriter, request *http.Request) {
	usuarios := Usuarios{
		Usuario{1, "Maria da Silva", "maria@mailtest.com", "123456"},
		Usuario{2, "João da Silva", "joao@mailtest.com", "456789"},
		Usuario{3, "Paulo da Silva", "paulo@mailtest.com", "789123"},
	}

	json.NewEncoder(response).Encode(usuarios)
}

func usuario(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	fmt.Fprintln(response, "Rota 2")
	fmt.Fprintln(response, "ID: "+id)
}
