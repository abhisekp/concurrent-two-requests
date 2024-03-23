.PHONY: gen-mock test build clean run

test: gen-mock
	@go test -cover ./...

build: clean
	@go build -o dist/main .

clean:
	@rm -rf dist

run:
	@go run main.go

build-run: build
	@./dist/main

gen-mock:
	@mockery