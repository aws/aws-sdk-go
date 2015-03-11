default: generate

generate-protocol-test:
	go generate ./internal/fixtures/protocol

generate:
	go generate ./aws
	go generate ./service
