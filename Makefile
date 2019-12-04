.PHONY: fmt test ci clean

fmt:
	go fmt ./...

ci:
	go test ./...

test: fmt ci

clean:
	go clean -r ./...