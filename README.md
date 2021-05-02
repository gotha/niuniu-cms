# niuniu-cms

Minimalistic serverless headless CMS with GraphQL support written in GO

## Run localy

### Via docker-compose

Just get docker-compose and execute:

```
docker build -t local/niuniu-cms:0.0.1
```

and then start with

```sh
docker-compose up -d
```

the api will be accessible on [http://localhost:9999](http://localhost:9999)

### Build it yourself

you can either run `go build` and run the binary with the respective env configs or:

you can use make and execute:
```sh
make build
make run
```

## Tests 

to update the generated mocks you need to have `moq`:

```sh
go get github.com/matryer/moq
```

then execute:

```sh
go generate ./...
```


### Integration tests

To run the integration tests you need postgres database. Start it with

```sh
docker-compose up -d db
```

and run the tests like this:

```sh
TZ=UTC \
	DB_TIMEZONE="UTC" \
	DB_HOST=localhost \
	DB_USERNAME=cms \
	DB_PASSWORD=4rfvbgt5 \
	DB_NAME=cms \
	DB_PORT=5432 \
	DB_SSL_MODE=disable \
	MIGRATE_DB=true \
	LAMBDA=false  \
	go test -v -tags=integration ./...
```
