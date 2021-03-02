build-app:
	go build -o app cmd/main.go

start-latest:
	ENV=dev SSL='' REPO=huvalk TAG=latest docker-compose up

start-local:
	ENV=dev SSL='' REPO=huvalk TAG=local docker-compose up

build-local:
	docker build -t huvalk/app:local -f docker/app.Dockerfile .

test-cover:
	go test ./... -coverprofile cover.out.tmp | grep -v "no test files"
	cat cover.out.tmp | grep -v "_easyjson.go" > cover.out
	go tool cover -func cover.out | grep "total"
	rm -f cover.out*

lint:
	golangci-lint run ./...