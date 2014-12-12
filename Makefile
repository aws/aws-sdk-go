default: generate

clean:
	rm -rfv gen/*/**/*.go

install-gen:
	go install ./cmd/...

generate: clean install-gen
	go generate ./gen
	go install ./...
