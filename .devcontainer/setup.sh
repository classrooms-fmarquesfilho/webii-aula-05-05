#!/usr/bin/env bash
set -e

echo "==> Instalando sqlc..."
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

echo "==> Criando banco de dados 'aula'..."
# A feature postgres cria o usuário 'postgres' sem senha por padrão
# Aguarda o postgres estar pronto
for i in $(seq 1 10); do
    pg_isready -U postgres && break
    echo "  aguardando postgres... ($i)"
    sleep 2
done

createdb -U postgres aula 2>/dev/null || echo "  banco 'aula' já existe"

echo "==> Instalando dependências dos exercícios..."
(cd exercicio-a && go mod download) 2>/dev/null || true
(cd exercicio-b && go mod download) 2>/dev/null || true

echo ""
echo "✅ Ambiente pronto!"
echo ""
echo "  DATABASE_URL=$DATABASE_URL"
echo ""
echo "  Próximos passos:"
echo "    cd exercicio-a"
echo "    cat README.md"
