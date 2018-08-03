package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/rota1", rota1)
	http.HandleFunc("/rota2", rota2)

	port := "3000"
	server := http.ListenAndServe(":"+port, nil)
	fmt.Println("Servidor rodando na porta " + port + ".")
	log.Fatal(server)

}

func home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "API Go com net/http")
}

func rota1(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Rota 1")
}

func rota2(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Rota 2")
}
