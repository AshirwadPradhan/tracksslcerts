build:
	@go build -o bin/tracksslcerts

run: build fmt
	@./bin/tracksslcerts

fmt:
	@go fmt ./...

test:
	@go test -v -race ./...