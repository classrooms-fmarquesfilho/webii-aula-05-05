# Exercício A — Do map ao sqlc

**Duração**: ~30 min | **Sprint 2** | Aula 05/05/2026

---

## O que você vai fazer

Sua API de contatos da Sprint 1 usa um `map` em memória.  
Ao reiniciar o servidor, tudo se perde.  
Neste exercício você migra para PostgreSQL usando sqlc.

O código inicial já tem a API funcionando com `map`.  
Você vai substituir o `map` por chamadas sqlc, passo a passo.

---

## Passo 1 — Aplicar o schema no banco (2 min)

```bash
psql $DATABASE_URL -f db/schema/001_contacts.sql
```

Verifique:
```bash
psql $DATABASE_URL -c "\dt"
# deve mostrar a tabela "contacts"
```

---

## Passo 2 — Preencher as queries (10 min)

Abra `db/queries/contacts.sql`.  
Você verá 4 queries com `???` nos lugares das anotações sqlc.

Substitua cada `???` pela anotação correta: `:one`, `:many` ou `:exec`.

Dica: consulte os slides ou a tabela abaixo:

| Anotação | Retorno | Quando usar |
|----------|---------|-------------|
| `:many`  | `[]T, error` | SELECT que retorna vários registros |
| `:one`   | `T, error`   | SELECT/INSERT que retorna 1 registro |
| `:exec`  | `error`      | DELETE/UPDATE sem retorno |

Depois de preencher, rode:
```bash
sqlc generate
```

Abra `internal/db/contacts.sql.go` e observe as funções geradas.

**Pergunta**: qual tipo Go foi gerado para `id SERIAL`? E para `created_at TIMESTAMPTZ`?

---

## Passo 3 — Conectar o handler ao banco (15 min)

Abra `handler/contacts.go`.  
Você verá os handlers usando `a.contacts` (o map).  
Há comentários `// TODO` marcando cada trecho a migrar.

Substitua cada `// TODO` pela chamada sqlc equivalente.

Depois, rode o servidor:
```bash
go run ./cmd/api
```

Teste em outro terminal:
```bash
# Criar contato
curl -s -X POST http://localhost:8080/contacts \
  -H "Content-Type: application/json" \
  -d '{"name":"Maria","email":"maria@example.com"}' | jq

# Listar
curl -s http://localhost:8080/contacts | jq

# Reinicie o servidor (Ctrl+C + go run ./cmd/api) e liste de novo
# Os dados devem PERSISTIR — essa é a diferença do map
```

---

## Passo 4 — Verificar (automático)

```bash
go test ./...
```

Os testes verificam os endpoints. Se passar, o exercício está completo.

---

## Gabarito

Se travar, o gabarito está em `handler/contacts_gabarito.go`.  
Tente sozinho primeiro — o gabarito só tem valor depois de você ter tentado.
