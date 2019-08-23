// +build go1.7

package endpoints

import (
	"encoding/json"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestUnmarshalRegionRegex(t *testing.T) {
	var input = []byte(`
{
    "regionRegex": "^(us|eu|ap|sa|ca)\\-\\w+\\-\\d+$"
}`)

	p := partition{}
	err := json.Unmarshal(input, &p)
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	expectRegexp, err := regexp.Compile(`^(us|eu|ap|sa|ca)\-\w+\-\d+$`)
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if e, a := expectRegexp.String(), p.RegionRegex.Regexp.String(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestUnmarshalRegion(t *testing.T) {
	var input = []byte(`
{
	"aws-global": {
	  "description": "AWS partition-global endpoint"
	},
	"us-east-1": {
	  "description": "US East (N. Virginia)"
	}
}`)

	rs := regions{}
	err := json.Unmarshal(input, &rs)
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if e, a := 2, len(rs); e != a {
		t.Errorf("expect %v len, got %v", e, a)
	}
	r, ok := rs["aws-global"]
	if !ok {
		t.Errorf("expect found, was not")
	}
	if e, a := "AWS partition-global endpoint", r.Description; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

	r, ok = rs["us-east-1"]
	if !ok {
		t.Errorf("expect found, was not")
	}
	if e, a := "US East (N. Virginia)", r.Description; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestUnmarshalServices(t *testing.T) {
	var input = []byte(`
{
	"acm": {
	  "endpoints": {
		"us-east-1": {}
	  }
	},
	"apigateway": {
      "isRegionalized": true,
	  "endpoints": {
		"us-east-1": {},
        "us-west-2": {}
	  }
	},
	"notRegionalized": {
      "isRegionalized": false,
	  "endpoints": {
		"us-east-1": {},
        "us-west-2": {}
	  }
	}
}`)

	ss := services{}
	err := json.Unmarshal(input, &ss)
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if e, a := 3, len(ss); e != a {
		t.Errorf("expect %v len, got %v", e, a)
	}
	s, ok := ss["acm"]
	if !ok {
		t.Errorf("expect found, was not")
	}
	if e, a := 1, len(s.Endpoints); e != a {
		t.Errorf("expect %v len, got %v", e, a)
	}
	if e, a := boxedBoolUnset, s.IsRegionalized; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

	s, ok = ss["apigateway"]
	if !ok {
		t.Errorf("expect found, was not")
	}
	if e, a := 2, len(s.Endpoints); e != a {
		t.Errorf("expect %v len, got %v", e, a)
	}
	if e, a := boxedTrue, s.IsRegionalized; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

	s, ok = ss["notRegionalized"]
	if !ok {
		t.Errorf("expect found, was not")
	}
	if e, a := 2, len(s.Endpoints); e != a {
		t.Errorf("expect %v len, got %v", e, a)
	}
	if e, a := boxedFalse, s.IsRegionalized; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestUnmarshalEndpoints(t *testing.T) {
	var inputs = []byte(`
{
	"aws-global": {
	  "hostname": "cloudfront.amazonaws.com",
	  "protocols": [
		"http",
		"https"
	  ],
	  "signatureVersions": [ "v4" ],
	  "credentialScope": {
		"region": "us-east-1",
		"service": "serviceName"
	  },
	  "sslCommonName": "commonName"
	},
	"us-east-1": {}
}`)

	es := endpoints{}
	err := json.Unmarshal(inputs, &es)
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if e, a := 2, len(es); e != a {
		t.Errorf("expect %v len, got %v", e, a)
	}
	s, ok := es["aws-global"]
	if !ok {
		t.Errorf("expect found, was not")
	}
	if e, a := "cloudfront.amazonaws.com", s.Hostname; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := []string{"http", "https"}, s.Protocols; !reflect.DeepEqual(e, a) {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := []string{"v4"}, s.SignatureVersions; !reflect.DeepEqual(e, a) {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := (credentialScope{"us-east-1", "serviceName"}), s.CredentialScope; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "commonName", s.SSLCommonName; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestEndpointResolve(t *testing.T) {
	defs := []endpoint{
		{
			Hostname:          "{service}.{region}.{dnsSuffix}",
			SignatureVersions: []string{"v2"},
			SSLCommonName:     "sslCommonName",
		},
		{
			Hostname:  "other-hostname",
			Protocols: []string{"http"},
			CredentialScope: credentialScope{
				Region:  "signing_region",
				Service: "signing_service",
			},
		},
	}

	e := endpoint{
		Hostname:          "{service}.{region}.{dnsSuffix}",
		Protocols:         []string{"http", "https"},
		SignatureVersions: []string{"v4"},
		SSLCommonName:     "new sslCommonName",
	}

	resolved := e.resolve("service", "region", "dnsSuffix",
		defs, Options{},
	)

	if e, a := "https://service.region.dnsSuffix", resolved.URL; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "signing_service", resolved.SigningName; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "signing_region", resolved.SigningRegion; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "v4", resolved.SigningMethod; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestEndpointMergeIn(t *testing.T) {
	expected := endpoint{
		Hostname:          "other hostname",
		Protocols:         []string{"http"},
		SignatureVersions: []string{"v4"},
		SSLCommonName:     "ssl common name",
		CredentialScope: credentialScope{
			Region:  "region",
			Service: "service",
		},
	}

	actual := endpoint{}
	actual.mergeIn(endpoint{
		Hostname:          "other hostname",
		Protocols:         []string{"http"},
		SignatureVersions: []string{"v4"},
		SSLCommonName:     "ssl common name",
		CredentialScope: credentialScope{
			Region:  "region",
			Service: "service",
		},
	})

	if e, a := expected, actual; !reflect.DeepEqual(e, a) {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestResolveEndpoint(t *testing.T) {
	resolved, err := testPartitions.EndpointFor("service2", "us-west-2")

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if e, a := "https://service2.us-west-2.amazonaws.com", resolved.URL; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "us-west-2", resolved.SigningRegion; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "service2", resolved.SigningName; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if resolved.SigningNameDerived {
		t.Errorf("expect the signing name not to be derived, but was")
	}
}

func TestResolveEndpoint_DisableSSL(t *testing.T) {
	resolved, err := testPartitions.EndpointFor("service2", "us-west-2", DisableSSLOption)

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if e, a := "http://service2.us-west-2.amazonaws.com", resolved.URL; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "us-west-2", resolved.SigningRegion; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "service2", resolved.SigningName; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if resolved.SigningNameDerived {
		t.Errorf("expect the signing name not to be derived, but was")
	}
}

func TestResolveEndpoint_UseDualStack(t *testing.T) {
	resolved, err := testPartitions.EndpointFor("service1", "us-west-2", UseDualStackOption)

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if e, a := "https://service1.dualstack.us-west-2.amazonaws.com", resolved.URL; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "us-west-2", resolved.SigningRegion; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "service1", resolved.SigningName; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if resolved.SigningNameDerived {
		t.Errorf("expect the signing name not to be derived, but was")
	}
}

func TestResolveEndpoint_HTTPProtocol(t *testing.T) {
	resolved, err := testPartitions.EndpointFor("httpService", "us-west-2")

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if e, a := "http://httpService.us-west-2.amazonaws.com", resolved.URL; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "us-west-2", resolved.SigningRegion; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "httpService", resolved.SigningName; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if !resolved.SigningNameDerived {
		t.Errorf("expect the signing name to be derived")
	}
}

func TestResolveEndpoint_UnknownService(t *testing.T) {
	_, err := testPartitions.EndpointFor("unknownservice", "us-west-2")

	if err == nil {
		t.Errorf("expect error, got none")
	}

	_, ok := err.(UnknownServiceError)
	if !ok {
		t.Errorf("expect error to be UnknownServiceError")
	}
}

func TestResolveEndpoint_ResolveUnknownService(t *testing.T) {
	resolved, err := testPartitions.EndpointFor("unknown-service", "us-region-1",
		ResolveUnknownServiceOption)

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if e, a := "https://unknown-service.us-region-1.amazonaws.com", resolved.URL; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "us-region-1", resolved.SigningRegion; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "unknown-service", resolved.SigningName; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if !resolved.SigningNameDerived {
		t.Errorf("expect the signing name to be derived")
	}
}

func TestResolveEndpoint_UnknownMatchedRegion(t *testing.T) {
	resolved, err := testPartitions.EndpointFor("service2", "us-region-1")

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if e, a := "https://service2.us-region-1.amazonaws.com", resolved.URL; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "us-region-1", resolved.SigningRegion; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "service2", resolved.SigningName; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if resolved.SigningNameDerived {
		t.Errorf("expect the signing name not to be derived, but was")
	}
}

func TestResolveEndpoint_UnknownRegion(t *testing.T) {
	resolved, err := testPartitions.EndpointFor("service2", "unknownregion")

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if e, a := "https://service2.unknownregion.amazonaws.com", resolved.URL; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "unknownregion", resolved.SigningRegion; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "service2", resolved.SigningName; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if resolved.SigningNameDerived {
		t.Errorf("expect the signing name not to be derived, but was")
	}
}

func TestResolveEndpoint_StrictPartitionUnknownEndpoint(t *testing.T) {
	_, err := testPartitions[0].EndpointFor("service2", "unknownregion", StrictMatchingOption)

	if err == nil {
		t.Errorf("expect error, got none")
	}

	_, ok := err.(UnknownEndpointError)
	if !ok {
		t.Errorf("expect error to be UnknownEndpointError")
	}
}

func TestResolveEndpoint_StrictPartitionsUnknownEndpoint(t *testing.T) {
	_, err := testPartitions.EndpointFor("service2", "us-region-1", StrictMatchingOption)

	if err == nil {
		t.Errorf("expect error, got none")
	}

	_, ok := err.(UnknownEndpointError)
	if !ok {
		t.Errorf("expect error to be UnknownEndpointError")
	}
}

func TestResolveEndpoint_NotRegionalized(t *testing.T) {
	resolved, err := testPartitions.EndpointFor("globalService", "us-west-2")

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if e, a := "https://globalService.amazonaws.com", resolved.URL; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "us-east-1", resolved.SigningRegion; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "globalService", resolved.SigningName; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if !resolved.SigningNameDerived {
		t.Errorf("expect the signing name to be derived")
	}
}

func TestResolveEndpoint_AwsGlobal(t *testing.T) {
	resolved, err := testPartitions.EndpointFor("globalService", "aws-global")

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if e, a := "https://globalService.amazonaws.com", resolved.URL; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "us-east-1", resolved.SigningRegion; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "globalService", resolved.SigningName; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if !resolved.SigningNameDerived {
		t.Errorf("expect the signing name to be derived")
	}
}
func Test_Regional_Flag(t *testing.T) {
	const v3Doc = `
{
  "version": 3,
  "partitions": [
    {
      "defaults": {
        "hostname": "{service}.{region}.{dnsSuffix}",
        "protocols": [
          "https"
        ],
        "signatureVersions": [
          "v4"
        ]
      },
      "dnsSuffix": "amazonaws.com",
      "partition": "aws",
      "partitionName": "AWS Standard",
      "regionRegex": "^(us|eu|ap|sa|ca)\\-\\w+\\-\\d+$",
      "regions": {
        "ap-northeast-1": {
          "description": "Asia Pacific (Tokyo)"
        }
      },
      "services": {
        "acm": {
          "endpoints": {
             "ap-northeast-1": {}
    	  }
        },
        "s3": {
          "endpoints": {
             "ap-northeast-1": {}
    	  }
        },
		 "sts" : {
			"defaults" : { },
			"endpoints" : {
			  "ap-east-1" : { },
			  "ap-northeast-1" : { },
			  "ap-northeast-2" : { },
			  "ap-south-1" : { },
			  "ap-southeast-1" : { },
			  "ap-southeast-2" : { },
			  "aws-global" : { 
				"credentialScope" : {
				  "region" : "us-east-1"
				},
				"hostname" : "sts.amazonaws.com"
			  },
			  "ca-central-1" : { },
			  "eu-central-1" : { },
			  "eu-north-1" : { },
			  "eu-west-1" : { },
			  "eu-west-2" : { },
			  "eu-west-3" : { },
			  "sa-east-1" : { },
			  "us-east-1" : { },
			  "us-east-1-fips" : {
				"credentialScope" : {
				  "region" : "us-east-1"
				},
				"hostname" : "sts-fips.us-east-1.amazonaws.com"
			  },
			  "us-east-2" : { },
			  "us-east-2-fips" : {
				"credentialScope" : {
				  "region" : "us-east-2"
				},
				"hostname" : "sts-fips.us-east-2.amazonaws.com"
			  },
			  "us-west-1" : { },
			  "us-west-1-fips" : {
				"credentialScope" : {
				  "region" : "us-west-1"
				},
				"hostname" : "sts-fips.us-west-1.amazonaws.com"
			  },
			  "us-west-2" : { },
			  "us-west-2-fips" : {
				"credentialScope" : {
				  "region" : "us-west-2"
				},
				"hostname" : "sts-fips.us-west-2.amazonaws.com"
			  }
			},
			"partitionEndpoint" : "aws-global"
		  }
      }
    }
  ]
}`

	resolver, err := DecodeModel(strings.NewReader(v3Doc))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	cases := map[string]struct {
		service, region                                     string
		regional                                            bool
		ExpectURL, ExpectSigningMethod, ExpectSigningRegion string
		ExpectSigningNameDerived                            bool
	}{
		"acm/ap-northeast-1/regional": {
			service:                  "acm",
			region:                   "ap-northeast-1",
			regional:                 true,
			ExpectURL:                "https://acm.ap-northeast-1.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "ap-northeast-1",
		},
		"acm/ap-northeast-1/legacy": {
			service:                  "acm",
			region:                   "ap-northeast-1",
			regional:                 false,
			ExpectURL:                "https://acm.ap-northeast-1.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "ap-northeast-1",
		},
		// STS Endpoints resolver tests :
		"sts/us-west-2/regional": {
			service:                  "sts",
			region:                   "us-west-2",
			regional:                 true,
			ExpectURL:                "https://sts.us-west-2.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-west-2",
		},
		"sts/us-west-2/legacy": {
			service:                  "sts",
			region:                   "us-west-2",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/ap-east-1/regional": {
			service:                  "sts",
			region:                   "ap-east-1",
			regional:                 true,
			ExpectURL:                "https://sts.ap-east-1.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "ap-east-1",
		},
		"sts/ap-east-1/legacy": {
			service:                  "sts",
			region:                   "ap-east-1",
			regional:                 false,
			ExpectURL:                "https://sts.ap-east-1.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "ap-east-1",
		},
		"sts/us-west-2-fips/regional": {
			service:                  "sts",
			region:                   "us-west-2-fips",
			regional:                 true,
			ExpectURL:                "https://sts-fips.us-west-2.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-west-2",
		},
		"sts/us-west-2-fips/legacy": {
			service:                  "sts",
			region:                   "us-west-2-fips",
			regional:                 false,
			ExpectURL:                "https://sts-fips.us-west-2.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-west-2",
		},
		"sts/aws-global/regional": {
			service:                  "sts",
			region:                   "aws-global",
			regional:                 true,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/aws-global/legacy": {
			service:                  "sts",
			region:                   "aws-global",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/ap-south-1/regional": {
			service:                  "sts",
			region:                   "ap-south-1",
			regional:                 true,
			ExpectURL:                "https://sts.ap-south-1.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "ap-south-1",
		},
		"sts/ap-south-1/legacy": {
			service:                  "sts",
			region:                   "ap-south-1",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/ap-northeast-1/regional": {
			service:                  "sts",
			region:                   "ap-northeast-1",
			regional:                 true,
			ExpectURL:                "https://sts.ap-northeast-1.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "ap-northeast-1",
		},
		"sts/ap-northeast-1/legacy": {
			service:                  "sts",
			region:                   "ap-northeast-1",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/ap-southeast-1/regional": {
			service:                  "sts",
			region:                   "ap-southeast-1",
			regional:                 true,
			ExpectURL:                "https://sts.ap-southeast-1.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "ap-southeast-1",
		},
		"sts/ap-southeast-1/legacy": {
			service:                  "sts",
			region:                   "ap-southeast-1",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/ca-central-1/regional": {
			service:                  "sts",
			region:                   "ca-central-1",
			regional:                 true,
			ExpectURL:                "https://sts.ca-central-1.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "ca-central-1",
		},
		"sts/ca-central-1/legacy": {
			service:                  "sts",
			region:                   "ca-central-1",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/eu-central-1/regional": {
			service:                  "sts",
			region:                   "eu-central-1",
			regional:                 true,
			ExpectURL:                "https://sts.eu-central-1.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "eu-central-1",
		},
		"sts/eu-central-1/legacy": {
			service:                  "sts",
			region:                   "eu-central-1",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/eu-north-1/regional": {
			service:                  "sts",
			region:                   "eu-north-1",
			regional:                 true,
			ExpectURL:                "https://sts.eu-north-1.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "eu-north-1",
		},
		"sts/eu-north-1/legacy": {
			service:                  "sts",
			region:                   "eu-north-1",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/eu-west-1/regional": {
			service:                  "sts",
			region:                   "eu-west-1",
			regional:                 true,
			ExpectURL:                "https://sts.eu-west-1.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "eu-west-1",
		},
		"sts/eu-west-1/legacy": {
			service:                  "sts",
			region:                   "eu-west-1",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/eu-west-2/regional": {
			service:                  "sts",
			region:                   "eu-west-2",
			regional:                 true,
			ExpectURL:                "https://sts.eu-west-2.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "eu-west-2",
		},
		"sts/eu-west-2/legacy": {
			service:                  "sts",
			region:                   "eu-west-2",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/eu-west-3/regional": {
			service:                  "sts",
			region:                   "eu-west-3",
			regional:                 true,
			ExpectURL:                "https://sts.eu-west-3.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "eu-west-3",
		},
		"sts/eu-west-3/legacy": {
			service:                  "sts",
			region:                   "eu-west-3",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/sa-east-1/regional": {
			service:                  "sts",
			region:                   "sa-east-1",
			regional:                 true,
			ExpectURL:                "https://sts.sa-east-1.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "sa-east-1",
		},
		"sts/sa-east-1/legacy": {
			service:                  "sts",
			region:                   "sa-east-1",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/us-east-1/regional": {
			service:                  "sts",
			region:                   "us-east-1",
			regional:                 true,
			ExpectURL:                "https://sts.us-east-1.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/us-east-1/legacy": {
			service:                  "sts",
			region:                   "us-east-1",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/us-east-2/regional": {
			service:                  "sts",
			region:                   "us-east-2",
			regional:                 true,
			ExpectURL:                "https://sts.us-east-2.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-2",
		},
		"sts/us-east-2/legacy": {
			service:                  "sts",
			region:                   "us-east-2",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
		"sts/us-west-1/regional": {
			service:                  "sts",
			region:                   "us-west-1",
			regional:                 true,
			ExpectURL:                "https://sts.us-west-1.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-west-1",
		},
		"sts/us-west-1/legacy": {
			service:                  "sts",
			region:                   "us-west-1",
			regional:                 false,
			ExpectURL:                "https://sts.amazonaws.com",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "us-east-1",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var optionSlice []func(o *Options)
			optionSlice = append(optionSlice, func(o *Options) {
				if c.regional {
					o.STSRegionalEndpoint = RegionalSTSEndpoint
				}
			})

			actual, err := resolver.EndpointFor(c.service, c.region, optionSlice...)
			if err != nil {
				t.Fatalf("failed to resolve endpoint, %v", err)
			}

			if e, a := c.ExpectURL, actual.URL; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}

			if e, a := c.ExpectSigningMethod, actual.SigningMethod; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}

			if e, a := c.ExpectSigningNameDerived, actual.SigningNameDerived; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}

			if e, a := c.ExpectSigningRegion, actual.SigningRegion; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}

		})
	}
}

func Test_Regional_Flag_CN(t *testing.T) {

	resolver := AwsCnPartition()

	cases := map[string]struct {
		service, region                                     string
		regional                                            bool
		ExpectURL, ExpectSigningMethod, ExpectSigningRegion string
		ExpectSigningNameDerived                            bool
	}{
		"sts/cn-north-1/regional": {
			service:                  "sts",
			region:                   "cn-north-1",
			regional:                 true,
			ExpectURL:                "https://sts.cn-north-1.amazonaws.com.cn",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "cn-north-1",
		},
		"sts/cn-north-1/legacy": {
			service:                  "sts",
			region:                   "cn-north-1",
			regional:                 false,
			ExpectURL:                "https://sts.cn-north-1.amazonaws.com.cn",
			ExpectSigningMethod:      "v4",
			ExpectSigningNameDerived: true,
			ExpectSigningRegion:      "cn-north-1",
		},
	}

	for name, c := range cases {
		var optionSlice []func(o *Options)
		t.Run(name, func(t *testing.T) {
			if c.regional {
				optionSlice = append(optionSlice, STSRegionalEndpointOption)
			}
			actual, err := resolver.EndpointFor(c.service, c.region, optionSlice...)
			if err != nil {
				t.Fatalf("failed to resolve endpoint, %v", err)
			}

			if e, a := c.ExpectURL, actual.URL; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
			if e, a := c.ExpectSigningMethod, actual.SigningMethod; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
			if e, a := c.ExpectSigningNameDerived, actual.SigningNameDerived; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
			if e, a := c.ExpectSigningRegion, actual.SigningRegion; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}
