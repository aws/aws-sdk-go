### SDK Features

### SDK Enhancements
* `private/protocol/restjson/restjson`: Use json.Decoder to decrease memory allocation (#2141)
  * Update RESTJSON protocol unmarshaler to use json.Decoder instead of ioutil.ReadAll to reduce allocations.
* `private/protocol/jsonrpc/jsonrpc`: Use json.Decoder to decrease memory allocation (#2142)
  * Update JSONPRC protocol unmarshaler to use json.Decoder instead of ioutil.ReadAll to reduce allocations.

### SDK Bugs
