package aws

import (
	"os"
	"testing"
)

func TestEnvCreds(t *testing.T) {
	os.Clearenv()
	os.Setenv("AWS_ACCESS_KEY_ID", "access")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "secret")
	os.Setenv("AWS_SESSION_TOKEN", "token")

	auth, err := EnvCreds()
	if err != nil {
		t.Fatal(err)
	}

	if v, want := auth.AccessKeyID(), "access"; v != want {
		t.Errorf("Access key ID was %v, expected %v", v, want)
	}

	if v, want := auth.SecretAccessKey(), "secret"; v != want {
		t.Errorf("Secret access key was %v, expected %v", v, want)
	}

	if v, want := auth.SecurityToken(), "token"; v != want {
		t.Errorf("Security token was %v, expected %v", v, want)
	}
}

func TestEnvCredsNoAccessKeyID(t *testing.T) {
	os.Clearenv()
	os.Setenv("AWS_SECRET_ACCESS_KEY", "secret")

	auth, err := EnvCreds()
	if err != ErrAccessKeyIDNotFound {
		t.Fatalf("ErrAccessKeyIDNotFound expected, but was %#v/%#v", auth, err)
	}
}

func TestEnvCredsNoSecretAccessKey(t *testing.T) {
	os.Clearenv()
	os.Setenv("AWS_ACCESS_KEY_ID", "access")

	auth, err := EnvCreds()
	if err != ErrSecretAccessKeyNotFound {
		t.Fatalf("ErrSecretAccessKeyNotFound expected, but was %#v/%#v", auth, err)
	}
}

func TestEnvCredsAlternateNames(t *testing.T) {
	os.Clearenv()
	os.Setenv("AWS_ACCESS_KEY", "access")
	os.Setenv("AWS_SECRET_KEY", "secret")

	auth, err := EnvCreds()
	if err != nil {
		t.Fatal(err)
	}

	if v, want := auth.AccessKeyID(), "access"; v != want {
		t.Errorf("Access key ID was %v, expected %v", v, want)
	}

	if v, want := auth.SecretAccessKey(), "secret"; v != want {
		t.Errorf("Secret access key was %v, expected %v", v, want)
	}
}
