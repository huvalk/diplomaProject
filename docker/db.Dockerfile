# Dockerfile relative to docker-compose.yml

FROM postgres:12.6

RUN apt-get update && apt-get -y install git postgresql-12-cron