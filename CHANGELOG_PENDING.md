### SDK Features
* `service/s3/s3crypto`: Introduces `EncryptionClientV2` and `DecryptionClientV2` encryption and decryption clients which support
a new key wrapping algorithm `kms+context`.
  * `DecryptionClientV2` maintains the ability to decrypt objects encrypted using the `EncryptionClient`.
  * Please see `s3crypto` documentation for migration details.

### SDK Enhancements

### SDK Bugs
