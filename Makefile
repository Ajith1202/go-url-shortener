build:
	@go build -o bin/shurl

run: build
	@./bin/shurl