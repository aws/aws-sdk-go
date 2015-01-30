package aws

import (
	"testing"
)

func TestGlobalEndpoints(t *testing.T) {
	svcs := []string{"cloudfront", "iam", "importexport", "route53", "sts"}
	svc := Service{Config: &Config{Region: "mock-region-1"}}

	for _, name := range svcs {
		svc.ServiceName = name
		if svc.endpointForRegion() != name+".amazonaws.com" {
			t.Errorf("expected endpoint for %s to equal %s.amazonaws.com", name, name)
		}
	}
}

func TestServicesInCN(t *testing.T) {
	svcs := []string{"cloudfront", "iam", "importexport", "route53", "sts", "s3"}
	svc := Service{Config: &Config{Region: "cn-north-1"}}

	for _, name := range svcs {
		svc.ServiceName = name
		if svc.endpointForRegion() != name+"."+svc.Config.Region+".amazonaws.com.cn" {
			t.Errorf("expected endpoint for %s to equal %s.%s.amazonaws.com.cn", name, name, svc.Config.Region)
		}
	}
}
