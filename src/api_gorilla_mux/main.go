package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", obterIndex).Methods("GET")
	router.HandleFunc("/usuarios", obterUsuarios).Methods("GET")
	router.HandleFunc("/usuario/{id}", obterUsuario).Methods("GET")
	router.HandleFunc("/login", login).Methods("POST")

	port := "3000"
	server := http.ListenAndServe(":"+port, router)

	log.Fatal(server)
}

func obterIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Rest Go com gorilla/mux")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login")
}
