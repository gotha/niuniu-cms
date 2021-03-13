.PHONY: build clean deploy run

build_lambda:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/lambda .
build:
	export GO111MODULE=on
	go build -ldflags="-s -w" -o bin/server .

run:
	DB_HOST=localhost DB_USERNAME=gocms DB_PASSWORD=4rfvbgt5 DB_NAME=gocms DB_PORT=5432 DB_SSL_MODE=disable MIGRATE_DB=true LAMBDA=false ./bin/server

generate_gql:
	go run github.com/99designs/gqlgen generate

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose
