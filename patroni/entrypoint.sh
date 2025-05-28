#!/bin/bash
set -e

echo "Starting Patroni entrypoint for node: ${PATRONI_NAME}"

if [ ! -d "$PATRONI_POSTGRESQL_DATA_DIR" ]; then
    echo "Creating data directory: $PATRONI_POSTGRESQL_DATA_DIR"
    mkdir -p "$PATRONI_POSTGRESQL_DATA_DIR"
fi

echo "Setting permissions for data directory"
chown -R postgres:postgres "$PATRONI_POSTGRESQL_DATA_DIR" || true
chmod 700 "$PATRONI_POSTGRESQL_DATA_DIR" || true

rm -rf "${PATRONI_POSTGRESQL_DATA_DIR}.failed" 2>/dev/null || true

mkdir -p /var/run/postgresql
chown postgres:postgres /var/run/postgresql

chown -R postgres:postgres /etc/patroni /var/log/patroni /home/postgres || true

echo "Creating pgpass file"
echo "*:*:*:${PATRONI_REPLICATION_USERNAME}:${PATRONI_REPLICATION_PASSWORD}" > /tmp/pgpass
echo "*:*:*:${PATRONI_SUPERUSER_USERNAME}:${PATRONI_SUPERUSER_PASSWORD}" >> /tmp/pgpass
chown postgres:postgres /tmp/pgpass
chmod 600 /tmp/pgpass

echo "Substituting environment variables in patroni.yml"
envsubst < /etc/patroni/patroni.yml > /tmp/patroni.yml
chown postgres:postgres /tmp/patroni.yml

echo "Final Patroni configuration:"
cat /tmp/patroni.yml

echo "Waiting for etcd to be ready..."
until curl -s http://${PATRONI_ETCD3_HOSTS}/health > /dev/null 2>&1; do
    echo "Waiting for etcd..."
    sleep 2
done
echo "etcd is ready"

echo "Starting Patroni as postgres user"
exec su-exec postgres patroni /tmp/patroni.yml