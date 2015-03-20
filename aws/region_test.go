package aws

import (
	"os"
	"testing"
)

func TestEnvRegion(t *testing.T) {
	os.Clearenv()
	os.Setenv("AWS_DEFAULT_REGION", "region")

	region, err := EnvRegion()
	if err != nil {
		t.Fatal(err)
	}

	if v, want := region, "region"; v != want {
		t.Errorf("Region was %v, expected %v", v, want)
	}
}

func TestEnvRegionNoDefaultRegion(t *testing.T) {
	os.Clearenv()

	region, err := EnvRegion()
	if err != ErrDefaultRegionNotFound {
		t.Fatalf("ErrDefaultRegionNotFound expected, but was %#v/%#v", region, err)
	}
}

func TestProfileRegion(t *testing.T) {
	region, err := ProfileRegion("example_config.ini", "")
	if err != nil {
		t.Fatal(err)
	}

	if v, want := region, "region"; v != want {
		t.Errorf("Region was %v, but expected %v", v, want)
	}
}

func TestProfileRegionNoRegion(t *testing.T) {
	_, err := ProfileRegion("example_config.ini", "no_region")
	if err == nil {
		t.Fatalf("Error expected, but was nil")
	}
}
