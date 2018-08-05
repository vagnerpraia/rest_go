package main

import (
	"log"
	"net/http"
)

func main() {
	router := router()

	port := "3000"
	server := http.ListenAndServe(":"+port, router)

	log.Fatal(server)
}
