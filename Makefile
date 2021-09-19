test:
	go test ./... | { grep -v 'no test files'; true; }

build:
	go build -o bin/go-rest-api .

db-gen:
	go generate ./ent