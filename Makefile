all:run

build:
	@go build -o bin/glox

run: build
	@./bin/glox

test:
	@go test -v ./... 
