FROM mcr.microsoft.com/devcontainers/go:1.22
RUN apt-get update && apt-get install -y postgresql postgresql-client \
    && rm -rf /var/lib/apt/lists/*
RUN PG_HBA=$(find /etc/postgresql -name pg_hba.conf | head -1) \
    && sed -i 's/^local\s\+all\s\+postgres\s\+peer/local all postgres trust/' "$PG_HBA" \
    && echo "host all postgres 127.0.0.1/32 trust" >> "$PG_HBA"