build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/create cmd/create-transaction/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/list cmd/list-transactions/main.go
	