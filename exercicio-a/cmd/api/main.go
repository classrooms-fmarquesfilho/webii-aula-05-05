package main

import (
	"log"
	"net/http"

	"github.com/webii-ufrn/aula-0505/exercicio-a/handler"
)

func main() {
	router := handler.NewRouter()
	log.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
