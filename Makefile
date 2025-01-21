build:
	@go build -o bin/GoSecureAuth cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/GoSecureAuth
