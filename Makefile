SHELL := /bin/bash

.PHONY: go rust

go:
	@echo "build specs for golang"
	pushd go              && \
		go mod tidy       && \
		go generate ./... && \
		go build ./...    && \
		go fmt ./...      && \
		go test -v ./...     && \
	popd

rust:
	@echo "build specs for rust"
	pushd rust      && \
		cargo build && \
	popd