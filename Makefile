.PHONY: fmt test clean

fmt:
	go fmt ./...

test: fmt
	go test ./...

clean:
	go clean -r ./...