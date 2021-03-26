build-app:
	go build -o app cmd/main.go

start-latest:
	ENV=dev SSL='' REPO=huvalk TAG=latest docker-compose up

start-local:
	ENV=dev SSL='' REPO=huvalk TAG=local docker-compose up

start-db:
	docker-compose -f docker-compose.database.yml up

clear-db-local:
	psql -h localhost -p 8081 -U postgres -d hhton -c 'truncate users, team, invite, event, feed, skills, job, notification cascade'

clear-db-dev:
	psql -h team-up.online -p 8081 -U postgres -d hhton -c 'truncate users, team, invite, event, feed, skills, job, notification cascade'

refresh-db-local:
	make clear-db-local
	psql -h localhost -p 8081 -U postgres -d hhton -f config/hhton_public.sql

refresh-db-dev:
	make clear-db-dev
	sudo psql -h team-up.online -p 8081 -U postgres -d hhton -f config/hhton_public.sql

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
	scp config/team-up.online.conf root@hahao.ru:/etc/nginx/conf.d/