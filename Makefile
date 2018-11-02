LINTIGNOREDOT='awstesting/integration.+should not use dot imports'
LINTIGNOREDOC='service/[^/]+/(api|service|waiters)\.go:.+(comment on exported|should have comment or be unexported)'
LINTIGNORECONST='service/[^/]+/(api|service|waiters)\.go:.+(type|struct field|const|func) ([^ ]+) should be ([^ ]+)'
LINTIGNORESTUTTER='service/[^/]+/(api|service)\.go:.+(and that stutters)'
LINTIGNOREINFLECT='service/[^/]+/(api|errors|service)\.go:.+(method|const) .+ should be '
LINTIGNOREINFLECTS3UPLOAD='service/s3/s3manager/upload\.go:.+struct field SSEKMSKeyId should be '
LINTIGNOREENDPOINTS='aws/endpoints/defaults.go:.+(method|const) .+ should be '
LINTIGNOREDEPS='vendor/.+\.go'
LINTIGNOREPKGCOMMENT='service/[^/]+/doc_custom.go:.+package comment should be of the form'
UNIT_TEST_TAGS="example codegen awsinclude"

SDK_WITH_VENDOR_PKGS=$(shell go list -tags ${UNIT_TEST_TAGS} ./... | grep -v "/vendor/src")
SDK_ONLY_PKGS=$(shell go list ./... | grep -v "/vendor/")
SDK_UNIT_TEST_ONLY_PKGS=$(shell go list -tags ${UNIT_TEST_TAGS} ./... | grep -v "/vendor/")
SDK_GO_VERSION=$(shell go version | awk '''{print $$3}''' | tr -d '''\n''')

all: generate unit

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  api_info                to print a list of services and versions"
	@echo "  docs                    to build SDK documentation"
	@echo "  unit                    to run unit tests"
	@echo "  integration             to run integration tests"
	@echo "  performance             to run performance tests"
	@echo "  verify                  to verify tests"
	@echo "  lint                    to lint the SDK"
	@echo "  vet                     to vet the SDK"
	@echo "  generate                to go generate and make services"
	@echo "  gen-test                to generate protocol tests"
	@echo "  gen-services            to generate services"
	@echo "  get-deps                to go get the SDK dependencies"
	@echo "  get-deps-tests          to get the SDK's test dependencies"
	@echo "  get-deps-verify         to get the SDK's verification dependencies"

generate: cleanup-models gen-test gen-endpoints gen-services

gen-test: gen-protocol-test gen-codegen-test

gen-codegen-test: get-deps-codegen
	go generate ./private/model/api/codegentest/service

gen-services: get-deps-codegen
	go generate ./service

gen-protocol-test: get-deps-codegen
	go generate ./private/protocol/...

gen-endpoints: get-deps-codegen
	go generate ./models/endpoints/

cleanup-models:
	@echo "Cleaning up stale model versions"
	@./cleanup_models.sh

unit: get-deps verify
	@echo "go test SDK and vendor packages"
	@go test -tags ${UNIT_TEST_TAGS} $(SDK_UNIT_TEST_ONLY_PKGS)

unit-with-race-cover: get-deps verify
	@echo "go test SDK and vendor packages"
	@go test -tags ${UNIT_TEST_TAGS} -race -cpu=1,2,4 $(SDK_UNIT_TEST_ONLY_PKGS)

unit-old-go-race-cover: get-deps-tests
	@echo "go test SDK only packages for old Go versions"
	@go test -race -cpu=1,2,4 $(SDK_ONLY_PKGS)

ci-test: generate unit-with-race-cover ci-test-generate-validate

ci-test-generate-validate:
	@echo "CI test validate no generated code changes"
	@git add . -A
	@gitstatus=`git diff --cached --ignore-space-change`; \
	echo "$$gitstatus"; \
	if [ "$$gitstatus" != "" ] && [ "$$gitstatus" != "skipping validation" ]; then echo "$$gitstatus"; exit 1; fi

integration: integ-custom smoke-tests performance

integ-custom: get-deps-integ
	go test -tags "integration" -v ./awstesting/integration/customizations/...

cleanup-integ-buckets:
	go run -tags "integration" ./awstesting/cmd/bucket_cleanup/main.go "aws-sdk-go-integration"

smoke-tests: get-deps-integ
	gucumber -go-tags "integration" ./awstesting/integration/smoke

performance: get-deps-integ
	AWS_TESTING_LOG_RESULTS=${log-detailed} AWS_TESTING_REGION=$(region) AWS_TESTING_DB_TABLE=$(table) gucumber -go-tags "integration" ./awstesting/performance

sandbox-tests: sandbox-test-go15 sandbox-test-go15-novendorexp sandbox-test-go16 sandbox-test-go17 sandbox-test-go18 sandbox-test-go19 sandbox-test-gotip

sandbox-build-go15:
	docker build -f ./awstesting/sandbox/Dockerfile.test.go1.5 -t "aws-sdk-go-1.5" .
sandbox-go15: sandbox-build-go15
	docker run -i -t aws-sdk-go-1.5 bash
sandbox-test-go15: sandbox-build-go15
	docker run -t aws-sdk-go-1.5

sandbox-build-go15-novendorexp:
	docker build -f ./awstesting/sandbox/Dockerfile.test.go1.5-novendorexp -t "aws-sdk-go-1.5-novendorexp" .
sandbox-go15-novendorexp: sandbox-build-go15-novendorexp
	docker run -i -t aws-sdk-go-1.5-novendorexp bash
sandbox-test-go15-novendorexp: sandbox-build-go15-novendorexp
	docker run -t aws-sdk-go-1.5-novendorexp

sandbox-build-go16:
	docker build -f ./awstesting/sandbox/Dockerfile.test.go1.6 -t "aws-sdk-go-1.6" .
sandbox-go16: sandbox-build-go16
	docker run -i -t aws-sdk-go-1.6 bash
sandbox-test-go16: sandbox-build-go16
	docker run -t aws-sdk-go-1.6

sandbox-build-go17:
	docker build -f ./awstesting/sandbox/Dockerfile.test.go1.7 -t "aws-sdk-go-1.7" .
sandbox-go17: sandbox-build-go17
	docker run -i -t aws-sdk-go-1.7 bash
sandbox-test-go17: sandbox-build-go17
	docker run -t aws-sdk-go-1.7

sandbox-build-go18:
	docker build -f ./awstesting/sandbox/Dockerfile.test.go1.8 -t "aws-sdk-go-1.8" .
sandbox-go18: sandbox-build-go18
	docker run -i -t aws-sdk-go-1.8 bash
sandbox-test-go18: sandbox-build-go18
	docker run -t aws-sdk-go-1.8

sandbox-build-go19:
	docker build -f ./awstesting/sandbox/Dockerfile.test.go1.9 -t "aws-sdk-go-1.9" .
sandbox-go19: sandbox-build-go19
	docker run -i -t aws-sdk-go-1.9 bash
sandbox-test-go19: sandbox-build-go19
	docker run -t aws-sdk-go-1.9

sandbox-build-go110:
	docker build -f ./awstesting/sandbox/Dockerfile.test.go1.10 -t "aws-sdk-go-1.10" .
sandbox-go110: sandbox-build-go110
	docker run -i -t aws-sdk-go-1.10 bash
sandbox-test-go110: sandbox-build-go110
	docker run -t aws-sdk-go-1.10

sandbox-build-go111:
	docker build -f ./awstesting/sandbox/Dockerfile.test.go1.11 -t "aws-sdk-go-1.11" .
sandbox-go111: sandbox-build-go111
	docker run -i -t aws-sdk-go-1.11 bash
sandbox-test-go111: sandbox-build-go111
	docker run -t aws-sdk-go-1.11

sandbox-build-gotip:
	@echo "Run make update-aws-golang-tip, if this test fails because missing aws-golang:tip container"
	docker build -f ./awstesting/sandbox/Dockerfile.test.gotip -t "aws-sdk-go-tip" .
sandbox-gotip: sandbox-build-gotip
	docker run -i -t aws-sdk-go-tip bash
sandbox-test-gotip: sandbox-build-gotip
	docker run -t aws-sdk-go-tip

update-aws-golang-tip:
	docker build --no-cache=true -f ./awstesting/sandbox/Dockerfile.golang-tip -t "aws-golang:tip" .

verify: get-deps-verify lint vet

lint:
	@echo "go lint SDK and vendor packages"
	@lint=`golint ./...`; \
	lint=`echo "$$lint" | grep -E -v -e ${LINTIGNOREDOT} -e ${LINTIGNOREDOC} -e ${LINTIGNORECONST} -e ${LINTIGNORESTUTTER} -e ${LINTIGNOREINFLECT} -e ${LINTIGNOREDEPS} -e ${LINTIGNOREINFLECTS3UPLOAD} -e ${LINTIGNOREPKGCOMMENT} -e ${LINTIGNOREENDPOINTS}`; \
	echo "$$lint"; \
	if [ "$$lint" != "" ]; then exit 1; fi

SDK_BASE_FOLDERS=$(shell ls -d */ | grep -v vendor | grep -v awsmigrate)
GO_VET_CMD=go tool vet --all -shadow

vet:
	${GO_VET_CMD} ${SDK_BASE_FOLDERS}

get-deps: get-deps-tests get-deps-x-tests get-deps-codegen get-deps-verify

get-deps-tests:
	@echo "go get SDK testing dependencies"
	go get github.com/stretchr/testify

get-deps-x-tests:
	@echo "go get SDK testing golang.org/x dependencies"
	go get golang.org/x/net/http2

get-deps-integ: get-deps-tests
	@echo "go get SDK integ testing dependencies"
	go get github.com/gucumber/gucumber/cmd/gucumber

get-deps-codegen:
	@echo "go get SDK codegen dependencies"
	go get golang.org/x/net/html

get-deps-verify:
	@echo "go get SDK verification utilities"
	go get golang.org/x/lint/golint

bench:
	@echo "go bench SDK packages"
	@go test -run NONE -bench . -benchmem -tags 'bench' $(SDK_ONLY_PKGS)

bench-protocol:
	@echo "go bench SDK protocol marshallers"
	@go test -run NONE -bench . -benchmem -tags 'bench' ./private/protocol/...

docs:
	@echo "generate SDK docs"
	@# This env variable, DOCS, is for internal use
	@if [ -z ${AWS_DOC_GEN_TOOL} ]; then\
		rm -rf doc && bundle install && bundle exec yard;\
	else\
		$(AWS_DOC_GEN_TOOL) `pwd`;\
	fi

api_info:
	@go run private/model/cli/api-info/api-info.go
