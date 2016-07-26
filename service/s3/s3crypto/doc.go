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
	sess := session.New()
	// Create the KeyProvider
	kp, err = s3crypto.NewKMSKeyProvider(sess, cmkID, s3crypto.NewJSONMatDesc())
	if err != nil {
		return err
	}

	// Create the cryptography client
	// We need to pass the session here so S3 can use it. In addition, any decryption that
	// occurs will use the same session with KMS
	svc := s3crypto.New(sess, s3crypto.Authentication(kp))

Configuration of the S3 cryptography client

	cfg := s3crypto.Config{
		// Save instruction files to separate objects
		SaveStrategy: NewS3SaveStrategy(session.New(), ""),
		// Change instruction file suffix to .example
		InstructionFileSuffix: ".example",
		// Set temp folder path
		TempFolderPath: "/path/to/tmp/folder/",
	}

The default SaveStrategy is to the object's header.

The InstructionFileSuffix defaults to .instruction. Careful here though, if you do this, be sure you know
what that suffix is in grabbing data.  All requests will look for fooKey.example instead of fooKey.instruction.
This suffix only affects gets and not puts. Put uses the keyprovider's suffix.
*/
package s3crypto
