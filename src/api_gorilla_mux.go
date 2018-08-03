package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/rota1", rota1)
	router.HandleFunc("/rota2", rota2)
	router.HandleFunc("/rota2/{id}", rota2_id)

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

func rota2(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Rota 2")
}

func rota2_id(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	fmt.Fprintln(response, "Rota 2")
	fmt.Fprintln(response, "ID: "+id)
}
