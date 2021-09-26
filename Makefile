SERVER=./cmd/server/main.go

.PHONY: run
run: ## run the API server
	go run ${SERVER}

.PHONY: test
test: ## run unit tests
	go test -v ./... | { grep -v 'no test files'; true; }

.PHONY: build
build: ## build the API server binary
	go build -o bin/go-rest-api ${SERVER}

.PHONY: db-gen
db-gen: ## generate ent files
	go generate ./ent