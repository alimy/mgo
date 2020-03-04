GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

LDFLAGS += -X "github.com/alimy/mgo/toolkit/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "github.com/alimy/mgo/toolkit/version.GitHash=$(shell git rev-parse HEAD)"

.PHONY: build
build: fmt
	go build  -ldflags '$(LDFLAGS)' -o mgo main.go

.PHONY: generate
generate:
	-rm -f cmd/templates_gen.go
	go generate cmd/templates.go
	$(GOFMT) -w cmd/templates_gen.go

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)
