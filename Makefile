.PHONY: build clean deploy gomodgen

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/lambda .

run:
	go build -ldflags="-s -w" -o bin/local .
	./bin/local

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

