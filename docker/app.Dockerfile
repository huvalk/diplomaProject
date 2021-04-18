# Dockerfile relative to docker-compose.yml

FROM golang:alpine

COPY . /app

WORKDIR /app

RUN apk add make && make build-app


CMD ["sh", "-c", "/app/app"]