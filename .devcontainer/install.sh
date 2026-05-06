#!/usr/bin/env bash
set -e

apt-get update -qq
apt-get install -y -qq postgresql postgresql-client

# Configurar autenticação sem senha para o usuário postgres localmente
PG_HBA=$(find /etc/postgresql -name pg_hba.conf | head -1)
# Substituir peer por trust para conexões locais do usuário postgres
sed -i 's/^local\s\+all\s\+postgres\s\+peer/local all postgres trust/' "$PG_HBA"
echo "host all postgres 127.0.0.1/32 trust" >> "$PG_HBA"

echo "PostgreSQL instalado ✓"