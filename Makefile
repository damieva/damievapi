run_db:
	@docker run -p 5432:5432 postgres.local

build:
	@go build -o bin/damievapi main.go

run: build
	@./bin/damievapi
