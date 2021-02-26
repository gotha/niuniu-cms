.PHONY: build clean deploy run

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/lambda .

run:
	DB_HOST=localhost DB_USERNAME=gocms DB_PASSWORD=4rfvbgt5 DB_NAME=gocms DB_PORT=5432 DB_SSL_MODE=disable MIGRATE_DB=true LAMBDA=false go run ./main.go

generate_gql:
	go run github.com/99designs/gqlgen generate

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose
