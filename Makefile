SHELL := /bin/bash

.PHONY: go rust

go:
	@echo "build specs for golang"
	pushd go           && \
		go mod tidy    && \
		popd
	pushd go              && \
		go generate ./... && \
		go build ./...    && \
		popd
	pushd go         && \
		go fmt ./... && \
		popd

rust:
	@echo "build specs for rust"
	pushd rust      && \
		cargo build && \
		popd