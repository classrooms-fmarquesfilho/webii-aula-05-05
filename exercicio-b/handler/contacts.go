package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/webii-ufrn/aula-0505/exercicio-b/internal/repository"
)

type problemDetails struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func writeProblem(w http.ResponseWriter, status int, title, detail string) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(problemDetails{
		Type:   "https://webii.ufrn.br/errors/" + http.StatusText(status),
		Title:  title,
		Status: status,
		Detail: detail,
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// App recebe a interface — não sabe se é PostgreSQL ou memória.
type App struct {
	repo repository.ContactRepository
}

func NewRouter(repo repository.ContactRepository) http.Handler {
	a := &App{repo: repo}
	r := chi.NewRouter()
	r.Get("/contacts", a.listContacts)
	r.Post("/contacts", a.createContact)
	r.Get("/contacts/{id}", a.getContact)
	r.Delete("/contacts/{id}", a.deleteContact)
	return r
}

func (a *App) listContacts(w http.ResponseWriter, r *http.Request) {
	// TODO: substituir pela chamada correta usando a.repo
	// Dica: a.repo.List(r.Context())
	writeProblem(w, 501, "Not Implemented", "implemente este método")
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
	// TODO: substituir pela chamada correta usando a.repo
	// Dica: a.repo.Create(r.Context(), body.Name, body.Email)
	writeProblem(w, 501, "Not Implemented", "implemente este método")
}

func (a *App) getContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeProblem(w, 400, "Bad Request", "id must be an integer")
		return
	}
	// TODO: substituir pela chamada correta usando a.repo
	// Dica: a.repo.Get(r.Context(), int32(id))
	//       se errors.Is(err, repository.ErrNotFound) → 404
	_ = id
	_ = errors.Is
	writeProblem(w, 501, "Not Implemented", "implemente este método")
}

func (a *App) deleteContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeProblem(w, 400, "Bad Request", "id must be an integer")
		return
	}
	// TODO: substituir pela chamada correta usando a.repo
	// Dica: a.repo.Delete(r.Context(), int32(id))
	//       se errors.Is(err, repository.ErrNotFound) → 404
	_ = id
	writeProblem(w, 501, "Not Implemented", "implemente este método")
}
