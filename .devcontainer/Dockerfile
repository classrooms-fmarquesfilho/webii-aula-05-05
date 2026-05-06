FROM mcr.microsoft.com/devcontainers/go:1.22

RUN rm -f /etc/apt/sources.list.d/yarn.list /etc/apt/sources.list.d/yarn.list.save && \
    apt-get update && \
    apt-get install -y postgresql postgresql-client && \
    rm -rf /var/lib/apt/lists/*

RUN PG_HBA=$(find /etc/postgresql -name pg_hba.conf | head -1) && \
    echo "local all postgres trust" | cat - "$PG_HBA" > /tmp/pg_hba_new && \
    mv /tmp/pg_hba_new "$PG_HBA"