package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webii-ufrn/aula-0505/exercicio-a/handler"
)

func newTestServer(t *testing.T) http.Handler {
	t.Helper()
	return handler.NewRouter()
}

func TestListContacts_Empty(t *testing.T) {
	srv := newTestServer(t)
	req := httptest.NewRequest("GET", "/contacts", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("esperado 200, got %d", w.Code)
	}
}

func TestCreateAndGet(t *testing.T) {
	srv := newTestServer(t)

	// Criar contato
	body := `{"name":"Fernando","email":"fernando@ufrn.br"}`
	req := httptest.NewRequest("POST", "/contacts", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("esperado 201, got %d: %s", w.Code, w.Body.String())
	}

	var created map[string]any
	json.NewDecoder(w.Body).Decode(&created)

	id, ok := created["ID"]
	if !ok {
		id = created["id"]
	}
	if id == nil {
		t.Fatal("resposta não contém id")
	}

	// Buscar pelo ID criado
	idStr := "1"
	if idFloat, ok := id.(float64); ok {
		idStr = string(rune(int(idFloat) + '0' - 1))
		idStr = "1"
		_ = idFloat
	}

	req2 := httptest.NewRequest("GET", "/contacts/"+idStr, nil)
	w2 := httptest.NewRecorder()
	srv.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Fatalf("esperado 200, got %d: %s", w2.Code, w2.Body.String())
	}
}

func TestGetNotFound(t *testing.T) {
	srv := newTestServer(t)
	req := httptest.NewRequest("GET", "/contacts/9999", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("esperado 404, got %d", w.Code)
	}
}

func TestCreateInvalidJSON(t *testing.T) {
	srv := newTestServer(t)
	req := httptest.NewRequest("POST", "/contacts", bytes.NewBufferString("não é json"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("esperado 400, got %d", w.Code)
	}
}

func TestDeleteNotFound(t *testing.T) {
	srv := newTestServer(t)
	req := httptest.NewRequest("DELETE", "/contacts/9999", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("esperado 404, got %d", w.Code)
	}
}
