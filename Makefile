SHELL := /usr/bin/env bash

.PHONY: build
build: go_dependencies
	@bazel build //...

.PHONY: go_dependencies
go_dependencies:
	(go build ./... && bazel run gazelle -- update-repos -from_file=go.mod) &
	bazel run --logging=1 gazelle &
	wait
