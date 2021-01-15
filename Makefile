SHELL := /bin/bash

.PHONY: go

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
