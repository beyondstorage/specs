SHELL := /bin/bash

.PHONY: go

go:
	@echo "build specs for golang"
	@# execute twice to prevent spec file formatted
	@pushd go           && \
	  go generate ./... && \
	  go build ./...    && \
	  go run ./format   && \
	  popd
	pushd go            && \
	  go generate ./... && \
	  go build ./...    && \
	  go run ./format   && \
	  popd
	pushd go       && \
	  go fmt ./... && \
	  go mod tidy  && \
	  popd
