LINTIGNOREDOT='awstesting/integration.+should not use dot imports'
LINTIGNOREDOC='service/[^/]+/(api|service)\.go:.+(comment on exported|should have comment or be unexported)'
LINTIGNORECONST='service/[^/]+/(api|service)\.go:.+(type|struct field|const|func) ([^ ]+) should be ([^ ]+)'
LINTIGNORESTUTTER='service/[^/]+/(api|service)\.go:.+(and that stutters)'
LINTIGNOREINFLECT='service/[^/]+/(api|service)\.go:.+method .+ should be '
LINTIGNOREDEPS='vendor/.+\.go'

default: generate unit

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  api_info                to print a list of services and versions"
	@echo "  docs                    to build SDK documentation"
	@echo "  build                   to go build the SDK"
	@echo "  unit                    to run unit tests"
	@echo "  integration             to run integration tests"
	@echo "  verify                  to verify tests"
	@echo "  lint                    to lint the SDK"
	@echo "  vet                     to vet the SDK"
	@echo "  generate                to go generate and make services"
	@echo "  gen-test                to generate protocol tests"
	@echo "  gen-services            to generate services"
	@echo "  get-deps                to go get the SDK dependencies"
	@echo "  get-deps-unit           to get the SDK's unit test dependencies"
	@echo "  get-deps-integ          to get the SDK's integration test dependencies"
	@echo "  get-deps-verify         to get the SDK's verification dependencies"

generate: gen-test gen-endpoints gen-services

gen-test: generate-protocol-test

gen-services:
	go generate ./service

gen-protocol-test:
	go generate ./private/protocol/...

gen-endpoints:
	go generate ./private/endpoints

build:
	go build ./...

unit: get-deps-unit build verify
	go test ./...

integration: get-deps-integ
	go test ./awstesting/integration/customizations/... -tags=integration
	gucumber ./awstesting/integration/smoke

verify: get-deps-verify lint

lint:
	@echo "golint ./..."
	@lint=`golint ./...`; \
	lint=`echo "$$lint" | grep -E -v -e ${LINTIGNOREDOT} -e ${LINTIGNOREDOC} -e ${LINTIGNORECONST} -e ${LINTIGNORESTUTTER} -e ${LINTIGNOREINFLECT} -e ${LINTIGNOREDEPS}`; \
	echo "$$lint"; \
	if [ "$$lint" != "" ]; then exit 1; fi

vet:
	@echo "go vet ./..."
	@go tool vet -all -shadow .

get-deps: get-deps-unit get-deps-integ get-deps-verify
	@go get -v ./...

get-deps-unit:
	@go get github.com/stretchr/testify

get-deps-integ: get-deps-unit
	@go get github.com/lsegal/gucumber/cmd/gucumber

get-deps-verify:
	@go get github.com/golang/lint/golint

bench:
	@go test -bench . -benchmem -tags 'bench' ./...

bench-protocol:
	@go test -bench . -benchmem -tags 'bench' ./private/protocol/...

docs:
	rm -rf doc && bundle install && bundle exec yard

api_info:
	@go run private/model/cli/api-info/api-info.go
