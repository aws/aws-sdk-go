### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/crr`: Fixed a race condition that caused concurrent calls relying on endpoint discovery to share the same `url.URL` reference in their operation's `http.Request`.
