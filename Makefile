build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/auth cmd/auth/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/create-transaction cmd/transaction/create-transaction/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/list-transactions cmd/transaction/list-transactions/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/report-transactions cmd/transaction/report-transactions/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/create-category cmd/category/create-category/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/list-categories cmd/category/list-categories/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/order-categories cmd/category/update-order/main.go

deploy-dev:
	AUTH0_DOMAIN=dev-kovicniqndmbavnk.us.auth0.com AUTH0_AUDIENCE=FinancialManagerAPI serverless deploy --aws-profile PERSONAL --stage dev

deploy-staging:
	AUTH0_DOMAIN=controlefamiliar-staging.us.auth0.com AUTH0_AUDIENCE=ControleFamiliarAPI serverless deploy --aws-profile PERSONAL --stage staging

deploy-prod:
	AUTH0_DOMAIN=controlefamiliar-prod.us.auth0.com AUTH0_AUDIENCE=ControleFamiliarAPI serverless deploy --aws-profile PERSONAL --stage prod
