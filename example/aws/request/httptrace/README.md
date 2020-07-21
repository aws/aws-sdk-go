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
Message: first message
2020/07/21 14:21:23 Latency: 292.687003ms, Validate: 23.261µs, Build: 92.249µs, Attempts: 1,
	Attempt: 0, Latency: 292.541266ms, Sign: 159.755µs, Send: 292.173529ms, Unmarshal: 203.204µs, WillRetry: false,
		HTTP: Latency: 292.174168ms, ConnReused: false, GetConn: 238.080409ms, DNS: 22.774015ms, Connect: 22.774015ms, TLS: 200.809752ms, WriteRequest: 424.376µs, WaitResponseFirstByte: 292.058664ms, ReadResponseHeader: 115.196µs,

Message: second message
2020/07/21 14:21:29 Latency: 47.824618ms, Validate: 2.91µs, Build: 110.86µs, Attempts: 1,
	Attempt: 0, Latency: 47.68784ms, Sign: 237.076µs, Send: 47.29333ms, Unmarshal: 151.831µs, WillRetry: false,
		HTTP: Latency: 47.29391ms, ConnReused: true, WriteRequest: 285.042µs, WaitResponseFirstByte: 47.233202ms, ReadResponseHeader: 60.252µs,

Message: third message
2020/07/21 14:21:33 Latency: 31.435353ms, Validate: 1.603µs, Build: 29.356µs, Attempts: 1,
	Attempt: 0, Latency: 31.39293ms, Sign: 63.152µs, Send: 31.245591ms, Unmarshal: 81.123µs, WillRetry: false,
		HTTP: Latency: 31.24588ms, ConnReused: true, WriteRequest: 151.738µs, WaitResponseFirstByte: 31.208877ms, ReadResponseHeader: 36.731µs,

Message: last message
2020/07/21 14:21:37 Latency: 42.643276ms, Validate: 1.903µs, Build: 33.829µs, Attempts: 1,
	Attempt: 0, Latency: 42.582587ms, Sign: 57.561µs, Send: 42.264062ms, Unmarshal: 246.32µs, WillRetry: false,
		HTTP: Latency: 42.265295ms, ConnReused: true, WriteRequest: 178.259µs, WaitResponseFirstByte: 42.142391ms, ReadResponseHeader: 122.393µs,
```

