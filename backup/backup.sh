#!/bin/bash

set -eo pipefail

TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKUP_DIR="/backups"
FILENAME="backup_$TIMESTAMP.sql.gz"

mkdir -p "$BACKUP_DIR"

find_primary() {
    local primary=""
    for node in node1 node2; do
        echo "Checking $node via REST API..."
        if command -v curl >/dev/null 2>&1; then
            response=$(curl -s http://$node:8008 2>/dev/null || echo "failed")
            if echo "$response" | grep -q '"role":"master"'; then
                primary=$node
                echo "Found primary via REST API: $primary"
                break
            fi
        fi
    done

    if [ -z "$primary" ]; then
        echo "REST API check failed, trying direct connections..."
        for node in node1 node2; do
            echo "Testing direct connection to $node..."
            if PGPASSWORD="$POSTGRES_PASSWORD" pg_isready -h "$node" -p 5432 -U "$POSTGRES_USER" >/dev/null 2>&1; then
                if PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$node" -p 5432 -U "$POSTGRES_USER" -d postgres -t -c "SELECT NOT pg_is_in_recovery();" 2>/dev/null | grep -q 't'; then
                    primary=$node
                    echo "Found primary via direct connection: $primary"
                    break
                fi
            fi
        done
    fi

    if [ -z "$primary" ]; then
        echo "Could not determine primary, using DB_HOST: $DB_HOST"
        primary="$DB_HOST"
    fi

    echo "$primary"
}

DB_HOST=$(find_primary)

echo "Using database host: $DB_HOST"

echo "Testing connection to $DB_HOST..."
if ! PGPASSWORD="$POSTGRES_PASSWORD" pg_isready -h "$DB_HOST" -p 5432 -U "$POSTGRES_USER"; then
    echo "ERROR: Cannot connect to database at $DB_HOST"
    exit 1
fi

echo "Creating backup: $FILENAME"

PGPASSWORD="$POSTGRES_PASSWORD" pg_dump \
    -h "$DB_HOST" \
    -p 5432 \
    -U "$POSTGRES_USER" \
    -d "$POSTGRES_DB" \
    --verbose \
    --no-password \
    | gzip > "$BACKUP_DIR/$FILENAME"

if [ ! -f "$BACKUP_DIR/$FILENAME" ] || [ ! -s "$BACKUP_DIR/$FILENAME" ]; then
    echo "ERROR: Backup file was not created or is empty"
    exit 1
fi

echo "Backup created successfully: $FILENAME ($(du -h "$BACKUP_DIR/$FILENAME" | cut -f1))"

cd "$BACKUP_DIR"
backup_count=$(ls -1 backup_*.sql.gz 2>/dev/null | wc -l)
echo "Current backup count: $backup_count"

if [ "$backup_count" -gt "${BACKUP_RETENTION_COUNT:-7}" ]; then
    echo "Cleaning old backups (keeping ${BACKUP_RETENTION_COUNT:-7} most recent)..."
    ls -tp backup_*.sql.gz | tail -n +$((${BACKUP_RETENTION_COUNT:-7} + 1)) | xargs -r rm -v
fi

echo "Backup process completed successfully"