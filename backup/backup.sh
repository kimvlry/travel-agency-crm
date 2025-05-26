#!/bin/bash

set -eo pipefail

TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKUP_DIR="/backups"
FILENAME="backup_$TIMESTAMP.sql.gz"

mkdir -p "$BACKUP_DIR"

PGPASSWORD="$POSTGRES_PASSWORD" pg_dump -h postgres \
        -U "$POSTGRES_USER" \
        -d "$POSTGRES_DB" \
        | gzip > "$BACKUP_DIR/$FILENAME"

cd "$BACKUP_DIR"

ls -tp | grep 'backup_.*\.sql\.gz$' | tail -n +$((BACKUP_RETENTION_COUNT + 1)) | xargs -r rm --
