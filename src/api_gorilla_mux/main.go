package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	port := "3000"
	server := http.ListenAndServe(":"+port, router)

	log.Fatal(server)
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Rest Go com gorilla/mux")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login")
}
