LINTIGNOREDOT='internal/features.+should not use dot imports'
LINTIGNOREDOC='service/[^/]+/(api)\.go:.+(comment on exported|should have comment or be unexported)'

default: generate

generate-protocol-test:
	go generate ./internal/protocol/...

generate-test: generate-protocol-test

generate:
	go generate ./aws
	make services

services:
	go generate ./service

integration: deps
	go test ./internal/test/integration/... -tags=integration
	gucumber

lint: deps
	@echo "golint ./..."
	@lint=`golint ./...`; \
	lint=`echo "$$lint" | grep -E -v -e ${LINTIGNOREDOT} -e ${LINTIGNOREDOC}`; \
	echo "$$lint"; \
	if [ "$$lint" != "" ]; then exit 1; fi

unit: deps build lint
	go test ./...

build:
	go build ./...

deps:
	@go get ./...
	@go get github.com/lsegal/gucumber/cmd/gucumber
	@go get github.com/golang/lint/golint

api_info:
	@go run internal/model/cli/api-info/api-info.go
