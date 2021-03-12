build-app:
	go build -o app cmd/main.go

start-latest:
	ENV=dev SSL='' REPO=huvalk TAG=latest docker-compose up

start-local:
	ENV=dev SSL='' REPO=huvalk TAG=local docker-compose up

start-db:
	docker-compose -f docker-compose.database.yml up

test-import-db:
	psql -h localhost -p 5432 -U postgres -d hhton -f db_test_dump

build-local:
	docker build -t huvalk/app:local -f docker/app.Dockerfile .

generate-model:
	cd application/models; easyjson -pkg -all

test-cover:
	go test ./... -coverprofile cover.out.tmp | grep -v "no test files"
	cat cover.out.tmp | grep -v "_easyjson.go" > cover.out
	go tool cover -func cover.out | grep "total"
	rm -f cover.out*

lint:
	golangci-lint run ./...