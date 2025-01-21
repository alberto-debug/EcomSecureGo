build:
	@go build -o go/GoSecureAuth cmd/main.go

test:
	@go test -v ./...

run: build
	@./go/GoSecureAuth
