package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/webii-ufrn/aula-0505/exercicio-b/internal/repository"
)

func TestMemoryRepository_CreateAndGet(t *testing.T) {
	repo := repository.NewMemoryRepository()
	ctx := context.Background()

	c, err := repo.Create(ctx, "Fernando", "fernando@ufrn.br")
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if c.Name != "Fernando" {
		t.Fatalf("esperado Name=Fernando, got %q", c.Name)
	}
	if c.ID == 0 {
		t.Fatal("ID não foi atribuído")
	}

	got, err := repo.Get(ctx, c.ID)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if got.Email != "fernando@ufrn.br" {
		t.Fatalf("email errado: %q", got.Email)
	}
}

func TestMemoryRepository_GetNotFound(t *testing.T) {
	repo := repository.NewMemoryRepository()
	_, err := repo.Get(context.Background(), 9999)
	if !errors.Is(err, repository.ErrNotFound) {
		t.Fatalf("esperado ErrNotFound, got %v", err)
	}
}

func TestMemoryRepository_Delete(t *testing.T) {
	repo := repository.NewMemoryRepository()
	ctx := context.Background()

	c, _ := repo.Create(ctx, "Ana", "ana@ufrn.br")

	if err := repo.Delete(ctx, c.ID); err != nil {
		t.Fatalf("Delete: %v", err)
	}

	_, err := repo.Get(ctx, c.ID)
	if !errors.Is(err, repository.ErrNotFound) {
		t.Fatalf("após delete esperado ErrNotFound, got %v", err)
	}
}

func TestMemoryRepository_DeleteNotFound(t *testing.T) {
	repo := repository.NewMemoryRepository()
	err := repo.Delete(context.Background(), 9999)
	if !errors.Is(err, repository.ErrNotFound) {
		t.Fatalf("esperado ErrNotFound, got %v", err)
	}
}

func TestMemoryRepository_List(t *testing.T) {
	repo := repository.NewMemoryRepository()
	ctx := context.Background()

	repo.Create(ctx, "A", "a@x.com")
	repo.Create(ctx, "B", "b@x.com")

	list, err := repo.List(ctx)
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(list) != 2 {
		t.Fatalf("esperado 2 itens, got %d", len(list))
	}
}
