default: generate

install-gen:
	go install ./cmd/...

generate: install-gen
	go generate ./service
	go install ./service/...
