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
	go get github.com/lsegal/gucumber/cmd/gucumber
	go test ./internal/test/integration/... -tags=integration
	gucumber

lint: deps
	lint=`golint ./...`; \
	lint=`echo "$$lint" | grep -E -v -e ${LINTIGNOREDOT} -e ${LINTIGNOREDOC}`; \
	echo "$$lint"; \
	if [ "$$lint" != "" ]; then exit 1; fi

unit: deps lint
	go test ./...

deps:
	go get ./...

api_info:
	@go run internal/model/cli/api-info/api-info.go
