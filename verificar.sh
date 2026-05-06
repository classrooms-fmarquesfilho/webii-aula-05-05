#!/usr/bin/env bash

set -e
PASS=0
FAIL=0

ok()   { echo "  ✅ $1"; PASS=$((PASS+1)); }
fail() { echo "  ❌ $1"; FAIL=$((FAIL+1)); }

echo ""
echo "=== Verificação pré-aula — Aula 05/05/2026 ==="
echo ""

# ── Go ──────────────────────────────────────────────────────────────────────
echo "[ Go ]"
if go version &>/dev/null; then
    ok "go instalado: $(go version | awk '{print $3}')"
else
    fail "go não encontrado"
fi

# ── sqlc ────────────────────────────────────────────────────────────────────
echo ""
echo "[ sqlc ]"
if sqlc version &>/dev/null; then
    ok "sqlc instalado: $(sqlc version)"
else
    fail "sqlc não encontrado — rode: go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest"
fi

# ── PostgreSQL ───────────────────────────────────────────────────────────────
echo ""
echo "[ PostgreSQL ]"
if pg_isready -U postgres &>/dev/null; then
    ok "PostgreSQL está rodando"
else
    fail "PostgreSQL não está rodando (pg_isready falhou)"
fi

if [ -n "$DATABASE_URL" ]; then
    ok "DATABASE_URL definida: $DATABASE_URL"
else
    fail "DATABASE_URL não definida"
fi

# ── Exercício A ──────────────────────────────────────────────────────────────
echo ""
echo "[ Exercício A ]"
cd exercicio-a

if go mod download &>/dev/null; then
    ok "go mod download OK"
else
    fail "go mod download falhou"
fi

# Aplica schema
if psql "$DATABASE_URL" -f db/schema/001_contacts.sql &>/dev/null; then
    ok "schema aplicado no banco"
else
    fail "falha ao aplicar schema (psql)"
fi

# Gera código sqlc (requer queries preenchidas — pula se não estiver)
if grep -q "???" db/queries/contacts.sql 2>/dev/null; then
    echo "  ⚠️  queries.sql ainda tem ???  (normal antes da aula)"
else
    if sqlc generate &>/dev/null; then
        ok "sqlc generate OK"
    else
        fail "sqlc generate falhou"
    fi
fi

# Testes com o map (versão inicial — devem passar sem banco)
if go test ./... &>/dev/null; then
    ok "go test ./... passou (versão map)"
else
    fail "go test ./... falhou"
fi

cd ..

# ── Exercício B ──────────────────────────────────────────────────────────────
echo ""
echo "[ Exercício B ]"
cd exercicio-b

if go mod download &>/dev/null; then
    ok "go mod download OK"
else
    fail "go mod download falhou"
fi

# Verifica compilação (testes vão falhar com 501 — é esperado antes da aula)
if go build ./... &>/dev/null; then
    ok "go build ./... OK"
else
    fail "go build ./... falhou (erro de compilação)"
fi

# Testa só o repositório (não depende do handler completo)
if go test ./internal/... &>/dev/null; then
    ok "go test ./internal/... passou (MemoryRepository OK)"
else
    fail "go test ./internal/... falhou"
fi

cd ..

# ── Resumo ───────────────────────────────────────────────────────────────────
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "  Passou: $PASS | Falhou: $FAIL"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
if [ $FAIL -eq 0 ]; then
    echo ""
    echo "  ✅ Tudo pronto para a aula!"
    echo ""
else
    echo ""
    echo "  ⚠️  Corrija os itens com ❌."
    echo ""
    exit 1
fi
