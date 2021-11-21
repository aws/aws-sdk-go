FROM golang AS base

ENV GOPRIVATE=github.com/blinkops

WORKDIR /go/src/github.com/aws/aws-sdk-go

COPY go.mod go.sum ./
RUN go mod download
COPY .. .

FROM base AS action-builder

WORKDIR /go/src/github.com/aws/aws-sdk-go
RUN go build -o ./aws-gen-api -tags codegen ./private/model/cli/gen-api/main.go && \
    chmod a+x aws-gen-api && \
    find './models/apis' -maxdepth 1 -exec ./aws-gen-api "{}/*/api-2.json" \;

FROM base AS aws-builder

WORKDIR /go/src/github.com/aws/aws-sdk-go
# Build the Integration Pack
RUN go build -ldflags="-s -w" -o /go/bin/blink-aws ./plugin.go

FROM alpine:3.14.0 AS plugin

WORKDIR /blink-aws

RUN apk add --no-cache && apk add libc6-compat

# Copy only the required artifacts.
COPY --from=action-builder /go/src/github.com/aws/aws-sdk-go/actions actions/
COPY --from=aws-builder /go/bin/blink-aws .

COPY --from=action-builder /go/src/github.com/aws/aws-sdk-go/custom_actions actions/

COPY config.yaml plugin.yaml ./

# Expose the gRPC port.
EXPOSE 1337

RUN chmod a+x blink-aws

ENTRYPOINT ./blink-aws
