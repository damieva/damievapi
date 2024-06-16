build:
	@cd src && go build -o ../bin/damievapi main.go

run: build
	@./bin/damievapi

run_db:
	@docker run -p 5432:5432 postgres.local
