LINTIGNOREDOT=-x 'internal/features.+should not use dot imports'

default: generate

generate-protocol-test:
	go generate ./internal/protocol/...

generate-test: generate-protocol-test

generate:
	go generate ./aws
	go generate ./service

integration:
	go get github.com/lsegal/gucumber/cmd/gucumber
	go test ./internal/test/integration/... -tags=integration
	gucumber

unit:
	lint=`golint ./aws/... && golint ./internal/...`; \
	lint=`echo "$$lint" | grep ${LINTIGNOREDOT}`; \
	echo $$lint; \
	if [ "$$lint" != "" ]; then exit 1; fi
	go test ./...
