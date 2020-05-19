package s3crypto

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
)

type mockKMS struct {
	kmsiface.KMSAPI
}

func TestAESGCMContentCipherBuilder_isUsingDeprecatedFeatures(t *testing.T) {
	builder := AESGCMContentCipherBuilder(NewKMSKeyGenerator(mockKMS{}, "cmkID"))

	features, ok := builder.(deprecatedFeatures)
	if !ok {
		t.Errorf("expected to implement deprecatedFeatures interface")
	}

	err := features.isUsingDeprecatedFeatures()
	if err == nil {
		t.Errorf("expected to recieve error for using deprecated features")
	}

	builder = AESGCMContentCipherBuilder(NewKMSContextKeyGenerator(mockKMS{}, "cmkID"))

	features, ok = builder.(deprecatedFeatures)
	if !ok {
		t.Errorf("expected to implement deprecatedFeatures interface")
	}

	err = features.isUsingDeprecatedFeatures()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestAESCBCContentCipherBuilder_isUsingDeprecatedFeatures(t *testing.T) {
	builder := AESCBCContentCipherBuilder(nil, nil)

	features, ok := builder.(deprecatedFeatures)
	if !ok {
		t.Errorf("expected to implement deprecatedFeatures interface")
	}

	err := features.isUsingDeprecatedFeatures()
	if err == nil {
		t.Errorf("expected to recieve error for using deprecated features")
	}
}
