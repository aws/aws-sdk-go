### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/signer/v4`: Sign `X-Amz-Server-Side-Encryption-Context` header.
  * Fixes signing for PutObject requests that set `SSEKMSEncryptionContext`.
