//go:build ignore
// +build ignore

// GABARITO — contacts_gabarito.go
//
// Este arquivo mostra a versão final do handler após migrar para sqlc.
// NÃO copie diretamente — tente implementar o Passo 3 você mesmo primeiro.
//
// Para usar o gabarito:
//   1. Renomeie contacts.go para contacts_map.go
//   2. Copie este arquivo para contacts.go (remova a linha //go:build ignore)

package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webii-ufrn/aula-0505/exercicio-a/internal/db"
)

type App struct {
	queries *db.Queries
}

func NewRouter() http.Handler {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("cannot connect to database: " + err.Error())
	}
	a := &App{queries: db.New(pool)}

	r := chi.NewRouter()
	r.Get("/contacts", a.listContacts)
	r.Post("/contacts", a.createContact)
	r.Get("/contacts/{id}", a.getContact)
	r.Delete("/contacts/{id}", a.deleteContact)
	return r
}

func (a *App) listContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := a.queries.ListContacts(r.Context())
	if err != nil {
		writeProblem(w, 500, "Internal Server Error", "failed to list contacts")
		return
	}
	if contacts == nil {
		contacts = []db.Contact{}
	}
	writeJSON(w, http.StatusOK, contacts)
}

func (a *App) createContact(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeProblem(w, 400, "Bad Request", "invalid json body")
		return
	}
	if body.Name == "" || body.Email == "" {
		writeProblem(w, 422, "Unprocessable Entity", "name and email are required")
		return
	}
	contact, err := a.queries.CreateContact(r.Context(), db.CreateContactParams{
		Name:  body.Name,
		Email: body.Email,
	})
	if err != nil {
		writeProblem(w, 500, "Internal Server Error", "failed to create contact")
		return
	}
	writeJSON(w, http.StatusCreated, contact)
}

func (a *App) getContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeProblem(w, 400, "Bad Request", "id must be an integer")
		return
	}
	contact, err := a.queries.GetContact(r.Context(), int32(id))
	if err != nil {
		writeProblem(w, 404, "Not Found", "contact not found")
		return
	}
	writeJSON(w, http.StatusOK, contact)
}

func (a *App) deleteContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeProblem(w, 400, "Bad Request", "id must be an integer")
		return
	}
	if err := a.queries.DeleteContact(r.Context(), int32(id)); err != nil {
		writeProblem(w, 404, "Not Found", "contact not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
