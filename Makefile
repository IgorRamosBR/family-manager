build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/create-transaction cmd/transaction/create-transaction/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/list-transactions cmd/transaction/list-transactions/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/report-transactions cmd/transaction/report-transactions/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/create-category cmd/category/create-category/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/list-categories cmd/category/list-categories/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/order-categories cmd/category/update-order/main.go

deploy-dev:
	serverless deploy --aws-profile PERSONAL --stage dev