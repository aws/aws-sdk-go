# Example

Demonstrates how the Go standard library `httptrace` can be used with the SDK
to collect HTTP request tracing timing using the SDK's API operation methods
like SNS's `PublishWithContext`.

The `trace.go` file demonstrates how the `httptrace` package's `ClientTrace`
can be created to gather timing information from HTTP requests made.

The `config.go` file provides additional configuration settings to control how
the HTTP client and its transport is configured. Such as, timeouts, and
keepalive.

## Usage

Run the example providing your SNS topic's ARN as the `-topic` parameter. This
example assumes that the region is provided via the environment variable and
the AWS shared credentials file (~/.aws/credentials)'s `default` provide
provides credentials.

```sh
AWS_REGION=us-west-2 go run -tags example . -topic arn:aws:sns:us-west-2:0123456789:mytopicname
```

Once the example starts you'll be prompted with a `Message:` statement. Input
the message that you'd like to send to the topic on a single line and hit
`enter` to send it.

```
Message: My Really cool Message
```

The example will output the http trace timing information for how long the request took.

```
2020/07/21 15:28:13 Latency: 290.859243ms, Validate: 20.607µs, Build: 114.17µs, Attempts: 1,
	Attempt: 0, Latency: 290.691814ms, Sign: 174.017µs, Send: 290.303676ms, Unmarshal: 196.662µs, WillRetry: false,
		HTTP: Latency: 290.304018ms, ConnReused: false, GetConn: 239.350422ms, DNS: 40.48639ms, Connect: 40.48639ms, TLS: 184.738452ms, WriteRequest: 310.311µs, WaitResponseFirstByte: 290.25327ms, ReadResponseHeader: 50.392µs,

Message: second
2020/07/21 15:28:15 Latency: 34.778924ms, Validate: 3.231µs, Build: 66.932µs, Attempts: 1,
	Attempt: 0, Latency: 34.685914ms, Sign: 92.238µs, Send: 33.879391ms, Unmarshal: 698.703µs, WillRetry: false,
		HTTP: Latency: 33.880956ms, ConnReused: true, GetConn: 29.447µs, WriteRequest: 151.002µs, WaitResponseFirstByte: 33.307018ms, ReadResponseHeader: 572.807µs,

Message: thrid
2020/07/21 15:28:16 Latency: 35.49632ms, Validate: 1.391µs, Build: 30.989µs, Attempts: 1,
	Attempt: 0, Latency: 35.446929ms, Sign: 59.237µs, Send: 35.218487ms, Unmarshal: 164.304µs, WillRetry: false,
		HTTP: Latency: 35.219014ms, ConnReused: true, GetConn: 36.817µs, WriteRequest: 160.361µs, WaitResponseFirstByte: 35.170978ms, ReadResponseHeader: 47.592µs,

Message: fourth
2020/07/21 15:28:21 Latency: 39.099871ms, Validate: 1.613µs, Build: 33.838µs, Attempts: 1,
	Attempt: 0, Latency: 39.037477ms, Sign: 59.805µs, Send: 38.84714ms, Unmarshal: 125.874µs, WillRetry: false,
		HTTP: Latency: 38.847806ms, ConnReused: true, GetConn: 47.417µs, WriteRequest: 224.433µs, WaitResponseFirstByte: 38.803854ms, ReadResponseHeader: 43.478µs,
```

