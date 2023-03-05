build-backend:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/auth cmd/auth/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/create-transactions cmd/transaction/create-transactions/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/list-transactions cmd/transaction/list-transactions/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/report-transactions cmd/transaction/report-transactions/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/create-categories cmd/category/create-categories/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/list-categories cmd/category/list-categories/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/order-categories cmd/category/order-categories/main.go

deploy-backend-dev:
	AUTH0_DOMAIN=dev-kovicniqndmbavnk.us.auth0.com AUTH0_AUDIENCE=FinancialManagerAPI serverless deploy --aws-profile PERSONAL --stage dev

deploy-backend-staging:
	AUTH0_DOMAIN=controlefamiliar-staging.us.auth0.com AUTH0_AUDIENCE=ControleFamiliarAPI serverless deploy --aws-profile PERSONAL --stage staging

deploy-backend-prod:
	AUTH0_DOMAIN=controlefamiliar-prod.us.auth0.com AUTH0_AUDIENCE=ControleFamiliarAPI serverless deploy --aws-profile PERSONAL --stage prod

remove-backend-dev:
	AUTH0_DOMAIN=dev-kovicniqndmbavnk.us.auth0.com AUTH0_AUDIENCE=FinancialManagerAPI serverless remove  --aws-profile PERSONAL --stage dev

remove-backend-staging:
	AUTH0_DOMAIN=controlefamiliar-staging.us.auth0.com AUTH0_AUDIENCE=ControleFamiliarAPI serverless remove  --aws-profile PERSONAL --stage staging

deploy-frontend-staging:
	cd web/fm-dashboard && npm run build-staging && npm run deploy-staging

deploy-frontend-prod:
	cd web/fm-dashboard && npm run build-prod && npm run deploy-prod