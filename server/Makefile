GIT_HASH = $(shell git describe --tags --dirty --always)

.PHONY: all run lint fmt test build clean gen-cert help

all: clean test build run ## clean dev env, run all linting and tests, build binary and run server

run: ## Run the server
	./bin/server

run-tls: gen-cert ## Run the server with self signed certificates
	./bin/server
	
test: lint fmt unit ## Execute all linting and tests 

unit: ## Execute unit tests 
	go test ./... -coverprofile cover.out
	
lint:
	go vet ./...

fmt:
	go fmt ./...

build: test ## Build server binary
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on GOSUMDB=off go build -o bin/server main.go

docker: ## Build server dockerfile
	docker build . -f docker/Dockerfile

gen-cert:
	openssl req -new -newkey rsa:4096 -x509 -sha256 -days 365 -nodes -out cert.pem -keyout key.pem
	
clean: ## Clean the development environment
	rm -rf cert.pem key.pem bin cover.out *.log

start:
	go run main.go 

start-nodemon:
	nodemon --exec go run main.go --signal SIGTERM

build-swagger:
	swag init --parseDependency --parseInternal

start-dev:
	nodemon --exec go run main.go --signal SIGTERM swag init --parseDependency --parseInternal
