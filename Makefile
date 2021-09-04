test:
	go test ./...

build:
	go build -o bin/go-rest-api .

db-gen:
	go generate ./ent