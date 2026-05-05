package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webii-ufrn/aula-0505/exercicio-b/handler"
	"github.com/webii-ufrn/aula-0505/exercicio-b/internal/repository"
)

func newTestServer(t *testing.T) http.Handler {
	t.Helper()
	// Usa MemoryRepository — sem banco de dados!
	repo := repository.NewMemoryRepository()
	return handler.NewRouter(repo)
}

func TestListContacts_Empty(t *testing.T) {
	srv := newTestServer(t)
	req := httptest.NewRequest("GET", "/contacts", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("esperado 200, got %d: %s", w.Code, w.Body.String())
	}
}

func TestCreateContact(t *testing.T) {
	srv := newTestServer(t)
	body := `{"name":"Fernando","email":"fernando@ufrn.br"}`
	req := httptest.NewRequest("POST", "/contacts", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("esperado 201, got %d: %s", w.Code, w.Body.String())
	}
	var c map[string]any
	json.NewDecoder(w.Body).Decode(&c)
	if c["Name"] != "Fernando" && c["name"] != "Fernando" {
		t.Fatalf("nome esperado 'Fernando', got %v", c)
	}
}

func TestGetContact_NotFound(t *testing.T) {
	srv := newTestServer(t)
	req := httptest.NewRequest("GET", "/contacts/9999", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("esperado 404, got %d", w.Code)
	}
}

func TestDeleteContact(t *testing.T) {
	srv := newTestServer(t)

	// Criar
	body := `{"name":"Ana","email":"ana@ufrn.br"}`
	req := httptest.NewRequest("POST", "/contacts", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("create: esperado 201, got %d", w.Code)
	}

	// Deletar o ID 1
	req2 := httptest.NewRequest("DELETE", "/contacts/1", nil)
	w2 := httptest.NewRecorder()
	srv.ServeHTTP(w2, req2)
	if w2.Code != http.StatusNoContent {
		t.Fatalf("delete: esperado 204, got %d: %s", w2.Code, w2.Body.String())
	}

	// Get após delete → 404
	req3 := httptest.NewRequest("GET", "/contacts/1", nil)
	w3 := httptest.NewRecorder()
	srv.ServeHTTP(w3, req3)
	if w3.Code != http.StatusNotFound {
		t.Fatalf("get após delete: esperado 404, got %d", w3.Code)
	}
}

func TestCreateContact_MissingFields(t *testing.T) {
	srv := newTestServer(t)
	body := `{"name":"Sem email"}`
	req := httptest.NewRequest("POST", "/contacts", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("esperado 422, got %d", w.Code)
	}
}
