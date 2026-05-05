-- db/queries/contacts.sql
--
-- Substitua cada ??? pela anotação sqlc correta:
--   :many  → retorna vários registros ([]Contact, error)
--   :one   → retorna exatamente 1 registro (Contact, error)
--   :exec  → sem retorno além de error
--
-- Dica: RETURNING * junto com INSERT → use :one

-- name: ListContacts ???
SELECT * FROM contacts ORDER BY created_at DESC;

-- name: GetContact ???
SELECT * FROM contacts WHERE id = $1;

-- name: CreateContact ???
INSERT INTO contacts (name, email)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteContact ???
DELETE FROM contacts WHERE id = $1;
