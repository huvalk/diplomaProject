define DROP_ALL_TABLES
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
endef

build-app:
	go build -o app cmd/main.go

start-latest:
	ENV=dev SSL='' REPO=huvalk TAG=latest docker-compose up

start-local:
	ENV=dev SSL='' REPO=huvalk TAG=local docker-compose up

start-db:
	docker-compose -f docker-compose.database.yml up

export DROP_ALL_TABLES
db-dump:
	pg_dump -U postgres -h localhost -p 8081 --column-inserts --data-only hhton > config/postgres/dump.sql

new-db-schema-from-dump:
	psql -h localhost -p 8081 -U postgres -d hhton -c "$$DROP_ALL_TABLES"
	psql -h localhost -p 8081 -U postgres -d hhton -f pkg/infrastructure/postgres.sql -f config/postgres/dump.sql -f pkg/infrastructure/funcs.sql

new-db-schema-local:
	psql -h localhost -p 8081 -U postgres -d hhton -c "$$DROP_ALL_TABLES"
	psql -h localhost -p 8081 -U postgres -d hhton -f pkg/infrastructure/postgres.sql -f config/postgres/hhton_public.sql -f pkg/infrastructure/funcs.sql

new-db-schema-dev:
	psql -h dev.team-up.online -p 8081 -U postgres -d hhton -c "$$DROP_ALL_TABLES"
	psql -h dev.team-up.online -p 8081 -U postgres -d hhton -f pkg/infrastructure/postgres.sql -f config/postgres/hhton_public.sql -f pkg/infrastructure/funcs.sql

refresh-db-local:
	make clear-db-local
	psql -h localhost -p 8081 -U postgres -d hhton -f config/postgres/hhton_public.sql

refresh-db-dev:
	make clear-db-dev
	sudo psql -h dev.team-up.online -p 8081 -U postgres -d hhton -f config/postgres/hhton_public.sql

build-local:
	docker build -t huvalk/app:local -f docker/app.Dockerfile .

generate-model:
	cd application/models; easyjson -pkg -all
	cd pkg/channel; easyjson -pkg

test-cover:
	go test ./... -coverprofile cover.out.tmp | grep -v "no test files"
	cat cover.out.tmp | grep -v "_easyjson.go" > cover.out
	go tool cover -func cover.out | grep "total"
	rm -f cover.out*

lint:
	golangci-lint run ./...

upload-nginx-conf:
	scp config/postgres/team-up.online.conf root@hahao.ru:/etc/nginx/conf.d/