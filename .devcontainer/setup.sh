#!/usr/bin/env bash
set -e
echo "==> Criando banco 'aula'..."
createdb -U postgres aula 2>/dev/null || echo "  banco 'aula' já existe"
psql "$DATABASE_URL" -c "SELECT version();" | head -1
echo "==> Instalando sqlc v1.26.0..."
go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.26.0
echo "==> Instalando dependências..."
(cd exercicio-a && go mod download)
(cd exercicio-b && go mod download)
grep -q '/go/bin' ~/.bashrc || echo 'export PATH=$PATH:/go/bin' >> ~/.bashrc
echo "✅ Ambiente pronto!"