### SDK Features

### SDK Enhancements

### SDK Bugs
* `service/lexruntimev2`: Add fix to prevent HTTP/2 event stream request hang on error with Go 1.15 through 1.17.
* `service/transcribestreamingservice`: Add fix to prevent HTTP/2 event stream request hang on error with Go 1.15 through 1.17.
  * Adds fix addressing an issue where SDK's bi-directional eventstream API operation request could hang when the HTTP/2 stream was closed by the service with a non-200 status code.
