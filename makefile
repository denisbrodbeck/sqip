.PHONY: build clean default test

build: clean
	@go build -o dist/sqip cmd/sqip/main.go

clean:
	@rm -rf dist/sqip*

test:
	go test ./...

default: build
