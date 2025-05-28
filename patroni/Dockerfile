FROM postgres:15-alpine

RUN apk add --no-cache \
    python3 \
    py3-pip \
    gcc \
    musl-dev \
    python3-dev \
    postgresql-dev \
    linux-headers \
    gettext \
    su-exec \
    bash \
    curl

RUN pip3 install --break-system-packages \
    patroni[etcd3]==3.2.2 \
    psycopg2-binary==2.9.9 \
    python-etcd==0.4.5 \
    urllib3

RUN mkdir -p /etc/patroni /var/log/patroni /home/postgres

COPY patroni.yml /etc/patroni/patroni.yml
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 5432 8008

ENTRYPOINT ["/entrypoint.sh"]