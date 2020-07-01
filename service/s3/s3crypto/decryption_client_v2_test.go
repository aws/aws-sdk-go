package s3crypto_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/s3/s3crypto"
)

func TestDecryptionClientV2_CheckDeprecatedFeatures(t *testing.T) {
	// AES/GCM/NoPadding with kms+context => allowed
	builder := s3crypto.AESGCMContentCipherBuilder(s3crypto.NewKMSContextKeyGenerator(kms.New(unit.Session), "cmkID"))
	_, err := s3crypto.NewEncryptionClientV2(unit.Session, builder)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// AES/GCM/NoPadding with kms => not allowed
	builder = s3crypto.AESGCMContentCipherBuilder(s3crypto.NewKMSKeyGenerator(kms.New(unit.Session), "cmkID"))
	_, err = s3crypto.NewEncryptionClientV2(unit.Session, builder)
	if err == nil {
		t.Error("expected error, but got nil")
	}

	// AES/CBC/PKCS5Padding with kms => not allowed
	builder = s3crypto.AESCBCContentCipherBuilder(s3crypto.NewKMSKeyGenerator(kms.New(unit.Session), "cmkID"), s3crypto.NewPKCS7Padder(128))
	_, err = s3crypto.NewEncryptionClientV2(unit.Session, builder)
	if err == nil {
		t.Error("expected error, but got nil")
	}

	// AES/CBC/PKCS5Padding with kms+context => not allowed
	builder = s3crypto.AESCBCContentCipherBuilder(s3crypto.NewKMSContextKeyGenerator(kms.New(unit.Session), "cmkID"), s3crypto.NewPKCS7Padder(128))
	_, err = s3crypto.NewEncryptionClientV2(unit.Session, builder)
	if err == nil {
		t.Error("expected error, but got nil")
	}
}
