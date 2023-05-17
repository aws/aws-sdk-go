### SDK Features

### SDK Enhancements

### SDK Bugs
* `rest`: Remove unnecessary path normalization behavior.
    * This behavior would incorrectly mutate request paths with URI-encoded characters, potentially resulting in misrouted requests.
* `config`: Deprecate `DisableRestProtocolURICleaning` config setting.
    * This setting no longer has any effect. REST-protocol paths will now never be normalized after serialization.