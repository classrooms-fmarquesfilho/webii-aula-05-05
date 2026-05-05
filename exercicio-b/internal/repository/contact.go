package repository

import (
	"context"
	"errors"
	"sync"
	"time"
)

// ErrNotFound é retornado quando um registro não existe.
var ErrNotFound = errors.New("not found")

// Contact é a entidade de domínio (sem dependência de sqlc).
type Contact struct {
	ID        int32
	Name      string
	Email     string
	CreatedAt time.Time
}

// ContactRepository define as operações de persistência.
// O handler depende desta interface, não de uma implementação concreta.
//
// TODO: adicione os dois métodos faltando:
//   - Create: recebe ctx, name string, email string → retorna Contact, error
//   - Delete: recebe ctx, id int32 → retorna error
type ContactRepository interface {
	List(ctx context.Context) ([]Contact, error)
	Get(ctx context.Context, id int32) (Contact, error)
	// ??? Create
	// ??? Delete
}

// ── MemoryRepository ─────────────────────────────────────────────────────────
// Implementação em memória — usada nos testes sem banco de dados.

type MemoryRepository struct {
	mu     sync.RWMutex
	items  map[int32]Contact
	nextID int32
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		items:  make(map[int32]Contact),
		nextID: 1,
	}
}

func (m *MemoryRepository) List(_ context.Context) ([]Contact, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := make([]Contact, 0, len(m.items))
	for _, c := range m.items {
		result = append(result, c)
	}
	return result, nil
}

func (m *MemoryRepository) Get(_ context.Context, id int32) (Contact, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.items[id]
	if !ok {
		return Contact{}, ErrNotFound
	}
	return c, nil
}

func (m *MemoryRepository) Create(_ context.Context, name, email string) (Contact, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	c := Contact{
		ID:        m.nextID,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}
	m.items[m.nextID] = c
	m.nextID++
	return c, nil
}

func (m *MemoryRepository) Delete(_ context.Context, id int32) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.items[id]; !ok {
		return ErrNotFound
	}
	delete(m.items, id)
	return nil
}
