# Exercício B — Interface de Repositório

**Duração**: ~10 min | Continuação do Exercício A

---

## O problema

No Exercício A, o handler ficou assim depois da migração:

```go
type App struct {
    queries *db.Queries   // concreto — acoplado ao sqlc
}
```

Isso funciona, mas torna impossível testar o handler sem banco de dados real.

## O objetivo

Extrair uma **interface** `ContactRepository` para que o handler não saiba se está
falando com PostgreSQL, memória, ou qualquer outra implementação.

---

## Passo 1 — Defina a interface (3 min)

O arquivo `internal/repository/contact.go` tem a interface com dois métodos faltando.  
Preencha os `???`:

```go
type ContactRepository interface {
    List(ctx context.Context) ([]Contact, error)
    Get(ctx context.Context, id int32) (Contact, error)
    ???   // Create: recebe name e email string, retorna Contact e error
    ???   // Delete: recebe id int32, retorna error
}
```

---

## Passo 2 — Ajuste o handler (5 min)

O arquivo `handler/contacts.go` tem o campo `repo ContactRepository` já declarado,
mas os métodos ainda chamam `a.queries.*`.

Substitua cada chamada `a.queries.XYZ(...)` pela chamada `a.repo.XYZ(...)` equivalente.

---

## Passo 3 — Verifique (2 min)

```bash
go test ./...
```

Os testes instanciam o handler com `MemoryRepository` — sem banco real.  
Se passar, o handler está desacoplado do PostgreSQL.

---

## Pergunta para reflexão

O `*db.Queries` gerado pelo sqlc satisfaz automaticamente a interface `ContactRepository`?  
Por quê sim ou por quê não? O que precisaria mudar?

> Resposta: não satisfaz diretamente porque as assinaturas diferem
> (sqlc usa `CreateContactParams` em vez de `name, email string`).
> Para satisfazer, você criaria um wrapper — ou ajustaria a interface para usar params.
> Isso é o que o ex02 da Lista 4 explora em detalhes.
