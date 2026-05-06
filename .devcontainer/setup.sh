#!/usr/bin/env bash
set -e

echo "==> Instalando PostgreSQL 15 via apt..."
sudo apt-get update -qq
sudo apt-get install -y -qq postgresql postgresql-client

echo "==> Configurando PostgreSQL..."
sudo service postgresql start

# Cria senha para o usuário postgres e o banco "aula"
sudo -u postgres psql -c "ALTER USER postgres WITH PASSWORD 'postgres';" 2>/dev/null || true
sudo -u postgres createdb aula 2>/dev/null || echo "  banco 'aula' já existe"

# Permite conexão local com senha
PG_HBA=$(sudo -u postgres psql -t -c "SHOW hba_file;" | tr -d ' ')
if ! grep -q "^host.*all.*postgres.*127" "$PG_HBA" 2>/dev/null; then
    echo "host all postgres 127.0.0.1/32 md5" | sudo tee -a "$PG_HBA" > /dev/null
    sudo service postgresql reload
fi

echo "==> Verificando conexão..."
psql "$DATABASE_URL" -c "SELECT version();" | head -1

echo "==> Instalando sqlc..."
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

echo "==> Instalando dependências dos exercícios..."
(cd exercicio-a && go mod download)
(cd exercicio-b && go mod download)

# Garante /go/bin no PATH
grep -q '/go/bin' ~/.bashrc || echo 'export PATH=$PATH:/go/bin' >> ~/.bashrc

echo ""
echo "✅ Ambiente pronto!"
echo "  DATABASE_URL=$DATABASE_URL"
