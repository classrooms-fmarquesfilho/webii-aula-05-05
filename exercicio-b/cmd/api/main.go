package main

import (
	"log"
	"net/http"

	"github.com/webii-ufrn/aula-0505/exercicio-b/handler"
	"github.com/webii-ufrn/aula-0505/exercicio-b/internal/repository"
)

func main() {
	// Em produção, trocaria MemoryRepository por PostgresRepository
	repo := repository.NewMemoryRepository()
	router := handler.NewRouter(repo)

	log.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
