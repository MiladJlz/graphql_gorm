build:
	@go build -o bin/api server.go
run: build
	@./bin/api