SHELL := /bin/bash

.PHONY: go

go:
	@echo "build specs for golang"
	pushd go            && \
	  go generate ./... && \
	  go fmt ./...      && \
	  go build ./...    && \
	  go mod tidy       && \
	  go run ./format   && \
	  popd
