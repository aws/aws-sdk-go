default: generate

generate-test:
	go generate ./internal/fixtures

generate:
	go generate ./aws
	go generate ./service
