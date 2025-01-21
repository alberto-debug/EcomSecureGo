build:
	@go build -o bin/GoSecureAuth cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/GoSecureAuth

kill:
	@fuser -k 8080/tcp || true

