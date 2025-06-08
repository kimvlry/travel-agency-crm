#!/bin/bash

set -e

echo "Starting Seqwall staircase test with DEPTH=${DEPTH}"
echo "Database URL: $DATABASE_URL"

if [ ! -d "/workspace/sql" ] || [ -z "$(ls -A /workspace/sql/*.up.sql 2>/dev/null || true)" ]; then
    echo "ERROR: No .up.sql migrations found in /workspace/sql/"
    echo "Available files:"
    ls -la /workspace/sql/ 2>/dev/null || echo "Directory /workspace/sql/ not found"
    exit 1
fi

echo "Found migrations:"
ls -la /workspace/sql/*.sql

seqwall staircase \
  --postgres-url "$DATABASE_URL" \
  --migrations-path "/workspace/sql" \
  --migrations-extension ".up.sql" \
  --schema public \
  --depth "${DEPTH}" \
  --upgrade "migrate -path /workspace/sql -database \"$DATABASE_URL\" up 1" \
  --downgrade "migrate -path /workspace/sql -database \"$DATABASE_URL\" down 1" \
  --test-snapshots true
