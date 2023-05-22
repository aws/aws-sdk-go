### SDK Features

### SDK Enhancements

### SDK Bugs
* `rest`: Revert removing unnecessary path normalization behavior.
  * This behavior would mutate request paths with URI-encoded characters, potentially resulting in misrouted requests.
* `config`: Revert deprecating `DisableRestProtocolURICleaning` config setting.
  * This setting will have an effect again. REST-protocol paths will now be normalized after serialization.
