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
