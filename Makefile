.DEFAULT_GOAL := build

build: test exports
	go build
.PHONY: build

test:
	go test
.PHONY: test

exports:
	${$Env:PORT = '8000'}