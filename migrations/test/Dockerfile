FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
    curl \
    gnupg \
    ca-certificates \
    lsb-release \
    postgresql-client \
    && rm -rf /var/lib/apt/lists/*

RUN curl -fsSL https://realkarych.github.io/seqwall-apt/public.key \
    | tee /etc/apt/trusted.gpg.d/seqwall.asc

RUN echo "deb [arch=$(dpkg --print-architecture)] https://realkarych.github.io/seqwall-apt stable main" \
    | tee /etc/apt/sources.list.d/seqwall.list

RUN apt-get update && apt-get install -y seqwall && rm -rf /var/lib/apt/lists/*

ENV MIGRATE_VERSION=4.17.1
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz \
    | tar xvz && mv migrate /usr/local/bin/migrate

WORKDIR /workspace

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]