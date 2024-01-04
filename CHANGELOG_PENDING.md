### SDK Features

### SDK Enhancements

### SDK Bugs
    The logging behavior in `aws/ec2metadata/token_provider.go` was updated: warnings about falling back to IMDSv1 are now logged only when LogLevel is set to `LogDebugWithDeprecated`. This change prevents unnecessary warnings when LogLevel is set to suppress messages.