SHELL := /bin/bash

.PHONY: go rust

go:
	@echo "build specs for golang"
	pushd go           && \
	  go generate ./... && \
	  go build ./...    && \
	  popd
	pushd go       && \
	  go fmt ./... && \
	  go mod tidy  && \
	  popd

rust:
	@echo "build specs for rust"
	pushd rust && \
	  cargo build && \
	  popd