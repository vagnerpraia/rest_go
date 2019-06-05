package main

import (
	"log"
	"net/http"

	"github.com/vagnerpraia/treinamento_go/src/api_gorilla_mux/router"
)

func main() {
	router := router.GetRouter()

	port := "3000"
	server := http.ListenAndServe(":"+port, router)

	log.Fatal(server)
}
