define DROP_ALL_TABLES
drop table event_users cascade;
drop table feed_users cascade;
drop table feed cascade;
drop table team_users cascade;
drop table prize_users cascade;
drop table prize cascade;
drop table notification cascade;
drop table invite cascade;
drop table team cascade;
drop table event cascade;
drop table job_skills_users cascade;
drop table users cascade;
drop table skills cascade;
drop table job cascade;
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
new-db-schema-local:
	psql -h localhost -p 8081 -U postgres -d hhton -c "$$DROP_ALL_TABLES"
	psql -h localhost -p 8081 -U postgres -d hhton -f pkg/infrastructure/postgres.sql -f config/hhton_public.sql

new-db-schema-dev:
	psql -h team-up.online -p 8081 -U postgres -d hhton -c "$$DROP_ALL_TABLES"
	psql -h team-up.online -p 8081 -U postgres -d hhton -f pkg/infrastructure/postgres.sql -f config/hhton_public.sql

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