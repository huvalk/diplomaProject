build-app:
	go build -o app cmd/main.go

test-cover:
	go test ./... -coverprofile cover.out.tmp | grep -v "no test files"
	cat cover.out.tmp | grep -v "_easyjson.go" > cover.out
	go tool cover -func cover.out | grep "total"
	rm -f cover.out*
