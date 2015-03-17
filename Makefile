default: generate

generate-protocol-test:
	go generate ./internal/protocol/...

generate-integration-test:
	go generate ./internal/fixtures/integration

generate-test: generate-protocol-test generate-integration-test

generate:
	go generate ./aws
	go generate ./service

test: generate-test
	go test ./... -tags=integration
