/*
Package s3crypto provides encryption to S3 using KMS and AES GCM.

Keyproviders are interfaces that handle masterkeys. Masterkeys are used to encrypt and decrypt the randomly
generated cipher keys. The SDK currently uses KMS to do this. A user does not need to provide a master key
since all that information is hidden in KMS.

Modes are interfaces that handle content encryption and decryption. It is an abstraction layer that instantiates
the ciphers. If content is being encrypted we generate the key and iv of the cipher. For decryption, we use the
metadata stored either on the object or an instruction file object to decrypt the contents.

Ciphers are interfaces that handle encryption and decryption of data. This may be key wrap ciphers or content
ciphers.

Creating an S3 cryptography client

	cmkID := "<some key ID>"
	sess := session.Must(session.NewSession())
	// Create the KeyProvider
	handler := s3crypto.NewKMSContextKeyGenerator(kms.New(sess), cmkID)

	// Create an encryption and decryption client
	// We need to pass the session here so S3 can use it. In addition, any decryption that
	// occurs will use the KMS client.
	svc := s3crypto.NewEncryptionClientV2(sess, s3crypto.AESGCMContentCipherBuilder(handler))
	svc := s3crypto.NewDecryptionClientV2(sess)

Configuration of the S3 cryptography client

	sess := session.Must(session.NewSession())
	handler := s3crypto.NewKMSContextKeyGenerator(kms.New(sess), cmkID)
	svc := s3crypto.NewEncryptionClientV2(sess, s3crypto.AESGCMContentCipherBuilder(handler), func (o *s3crypto.EncryptionClientOptions) {
		// Save instruction files to separate objects
		o.SaveStrategy = NewS3SaveStrategy(sess, "")

		// Change instruction file suffix to .example
		o.InstructionFileSuffix = ".example"

		// Set temp folder path
		o.TempFolderPath = "/path/to/tmp/folder/"

		// Any content less than the minimum file size will use memory
		// instead of writing the contents to a temp file.
		o.MinFileSize = int64(1024 * 1024 * 1024)
	})

The default SaveStrategy is to the object's header.

The InstructionFileSuffix defaults to .instruction. Careful here though, if you do this, be sure you know
what that suffix is in grabbing data.  All requests will look for fooKey.example instead of fooKey.instruction.
This suffix only affects gets and not puts. Put uses the keyprovider's suffix.

Registration of new wrap or cek algorithms are also supported by the SDK. Let's say we want to support `AES Wrap`
and `AES CTR`. Let's assume we have already defined the functionality.

	svc := s3crypto.NewDecryptionClientV2(sess, func(o *s3crypto.DecryptionClientOptions) {
		o.WrapRegistry["CustomWrap"] = NewCustomWrap
		o.CEKRegistry["CustomCEK"] = NewCustomCEK
	})

We have now registered these new algorithms to the decryption client. When the client calls `GetObject` and sees
the wrap as `CustomWrap` then it'll use that wrap algorithm. This is also true for `CustomCEK`.

For encryption adding a custom content cipher builder and key handler will allow for encryption of custom
defined ciphers.

	// Our wrap algorithm, CustomWrap
	handler := NewCustomWrap(key, iv)
	// Our content cipher builder, NewCustomCEKContentBuilder
	svc := s3crypto.NewEncryptionClientV2(sess, NewCustomCEKContentBuilder(handler))

Deprecations

The EncryptionClient and DecryptionClient types and their associated constructor functions have been deprecated.
Users of these clients should migrate to EncryptionClientV2 and DecryptionClientV2 types and constructor functions.

EncryptionClientV2 removes encryption support of the following features
	* AES/CBC/PKCS5Padding (content cipher)
	* kms (key wrap algorithm)

Attempting to construct an EncryptionClientV2 with deprecated features will result in an error returned back to the
calling application during construction of the client.

Users of `AES/CBC/PKCS5Padding` will need to migrate usage to `AES/GCM/NoPadding`.
Users of `kms` key provider will need to migrate `kms+context`.

DecryptionClientV2 client adds support for the `kms+context` key provider and maintains backwards comparability with
objects encrypted with the deprecated EncryptionClient.

Migrating from V1 to V2 Clients

Examples of how to migrate usage of the V1 clients to the V2 equivalents have been documented as usage examples of
the NewEncryptionClientV2 and NewDecryptionClientV2 functions.
*/
package s3crypto
