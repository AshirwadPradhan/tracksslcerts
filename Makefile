build:
	@go build -o bin/tracksslcerts

run: build fmt
	@npx tailwindcss-cli@latest build ./misc/app.css -o ./static/app.css
	@./bin/tracksslcerts

fmt:
	@go fmt ./...

test:
	@go test -v -race ./...