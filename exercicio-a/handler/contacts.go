package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/webii-ufrn/aula-0505/exercicio-a/internal/db"
)

// problemDetails segue RFC 7807
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

// ── Handler com map (Sprint 1) ──────────────────────────────────────────────

// contactMap é o armazenamento em memória da Sprint 1.
// Vamos substituí-lo pelo db.Queries conectado ao PostgreSQL.
type contactMap struct {
	mu      sync.RWMutex
	items   map[int32]db.Contact
	nextID  int32
}

func newContactMap() *contactMap {
	return &contactMap{items: make(map[int32]db.Contact), nextID: 1}
}

// App agrupa o router e o storage.
// PASSO 3: você vai adicionar `queries *db.Queries` aqui e remover `store`.
type App struct {
	store *contactMap
}

func NewRouter() http.Handler {
	a := &App{store: newContactMap()}

	r := chi.NewRouter()
	r.Get("/contacts", a.listContacts)
	r.Post("/contacts", a.createContact)
	r.Get("/contacts/{id}", a.getContact)
	r.Delete("/contacts/{id}", a.deleteContact)
	return r
}

// listContacts lista todos os contatos.
func (a *App) listContacts(w http.ResponseWriter, r *http.Request) {
	// TODO (Passo 3): substituir pelo código abaixo e remover o bloco do map
	//
	//   contacts, err := a.queries.ListContacts(r.Context())
	//   if err != nil {
	//       writeProblem(w, 500, "Internal Server Error", "failed to list contacts")
	//       return
	//   }
	//   writeJSON(w, http.StatusOK, contacts)

	a.store.mu.RLock()
	defer a.store.mu.RUnlock()
	list := make([]db.Contact, 0, len(a.store.items))
	for _, c := range a.store.items {
		list = append(list, c)
	}
	writeJSON(w, http.StatusOK, list)
}

// createContact cria um novo contato.
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

	// TODO (Passo 3): substituir pelo código abaixo e remover o bloco do map
	//
	//   contact, err := a.queries.CreateContact(r.Context(), db.CreateContactParams{
	//       Name:  body.Name,
	//       Email: body.Email,
	//   })
	//   if err != nil {
	//       writeProblem(w, 500, "Internal Server Error", "failed to create contact")
	//       return
	//   }
	//   writeJSON(w, http.StatusCreated, contact)

	a.store.mu.Lock()
	defer a.store.mu.Unlock()
	c := db.Contact{ID: a.store.nextID, Name: body.Name, Email: body.Email}
	a.store.items[a.store.nextID] = c
	a.store.nextID++
	writeJSON(w, http.StatusCreated, c)
}

// getContact retorna um contato pelo ID.
func (a *App) getContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeProblem(w, 400, "Bad Request", "id must be an integer")
		return
	}

	// TODO (Passo 3): substituir pelo código abaixo e remover o bloco do map
	//
	//   contact, err := a.queries.GetContact(r.Context(), int32(id))
	//   if err != nil {
	//       writeProblem(w, 404, "Not Found", "contact not found")
	//       return
	//   }
	//   writeJSON(w, http.StatusOK, contact)

	a.store.mu.RLock()
	defer a.store.mu.RUnlock()
	c, ok := a.store.items[int32(id)]
	if !ok {
		writeProblem(w, 404, "Not Found", "contact not found")
		return
	}
	writeJSON(w, http.StatusOK, c)
}

// deleteContact remove um contato pelo ID.
func (a *App) deleteContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeProblem(w, 400, "Bad Request", "id must be an integer")
		return
	}

	// TODO (Passo 3): substituir pelo código abaixo e remover o bloco do map
	//
	//   if err := a.queries.DeleteContact(r.Context(), int32(id)); err != nil {
	//       writeProblem(w, 404, "Not Found", "contact not found")
	//       return
	//   }
	//   w.WriteHeader(http.StatusNoContent)

	a.store.mu.Lock()
	defer a.store.mu.Unlock()
	if _, ok := a.store.items[int32(id)]; !ok {
		writeProblem(w, 404, "Not Found", "contact not found")
		return
	}
	delete(a.store.items, int32(id))
	w.WriteHeader(http.StatusNoContent)
}
