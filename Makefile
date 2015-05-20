LINTIGNOREDOT=-x 'internal/features.+should not use dot imports'

default: generate

generate-protocol-test:
	go generate ./internal/protocol/...

generate-test: generate-protocol-test

generate:
	go generate ./aws
	make service

service:
	go generate ./service

integration: deps
	go get github.com/lsegal/gucumber/cmd/gucumber
	go test ./internal/test/integration/... -tags=integration
	gucumber

unit: deps
	lint=`golint ./aws/... && golint ./internal/...`; \
	lint=`echo "$$lint" | grep ${LINTIGNOREDOT}`; \
	echo $$lint; \
	if [ "$$lint" != "" ]; then exit 1; fi
	go test ./...

deps:
	go get ./...

api_info:
	@go run internal/model/cli/api-info/api-info.go
