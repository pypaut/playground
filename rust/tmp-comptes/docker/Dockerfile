FROM postgres:14.18-bookworm

ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=pass
ENV POSTGRES_DB=db

COPY init.sql /docker-entrypoint-initdb.d/
