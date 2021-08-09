test:
	go test ./...
bench:
	go test ./... -bench=.
build:
	go build ./cmd/app
clean:
	rm -f ./app