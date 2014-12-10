package aws

import (
	"os"
	"testing"
)

func TestEnvAuth(t *testing.T) {
	os.Clearenv()
	os.Setenv("AWS_SECRET_KEY", "secret")
	os.Setenv("AWS_ACCESS_KEY", "access")
	auth, err := EnvCreds()
	if err != nil {
		t.Fatal(err)
	}
	if auth.SecretAccessKey() != "secret" {
		t.Fatalf("Expected secret key 'secret', got %s.", auth.SecretAccessKey())
	}
	if auth.AccessKeyID() != "access" {
		t.Fatalf("Expected access key 'access', got %s.", auth.AccessKeyID())
	}
	if auth.SecurityToken() != "" {
		t.Fatalf("Expected security token '', got %s.", auth.SecurityToken())
	}
}

func TestEnvAuthWithOnlySecurityToken(t *testing.T) {
	os.Clearenv()
	os.Setenv("AWS_SESSION_TOKEN", "token")
	auth, err := EnvCreds()
	if err != nil {
		t.Fatal(err)
	}
	if auth.SecurityToken() != "token" {
		t.Fatalf("Expected security token 'token', got %s.", auth.SecurityToken())
	}
}
