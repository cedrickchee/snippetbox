SHELL:=/bin/bash -O extglob

.PHONY: dev test deps cert

dev:
	@go run cmd/web/!(*_test).go

cert:
	@go run $GOROOT/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost

test:
	@go test -v -cover -coverprofile=coverage.out -covermode=atomic ./...

deps:
	@go get ./...