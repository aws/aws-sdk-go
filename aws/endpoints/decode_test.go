// +build go1.9

package endpoints

import (
	"strings"
	"testing"
)

func TestDecodeEndpoints_V3(t *testing.T) {
	const v3Doc = `{
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
      "dualstackDefaults": {
        "protocols": [
          "http",
          "https"
        ],
        "hostname": "{service}.{region}.{dualstackDnsSuffix}"
      },
      "dnsSuffix": "amazonaws.com",
      "dualstackDnsSuffix": "aws",
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
          "dualstackDefaults": {
            "protocols": [
              "http",
              "https"
            ],
            "hostname": "{service}.dualstack.{region}.{dualstackDnsSuffix}"
          },
          "dualstackDnsSuffix": "amazonaws.com",
          "endpoints": {
            "ap-northeast-1": {}
          },
          "dualstackEndpoints": {
            "us-west-2": {
              "hostname": "s3.dualstack.us-west-2.amazonaws.com",
              "signatureVersions": ["s3", "s3v4"]
            }
          }
        }
      }
    }
  ]
}`

	resolver, err := DecodeModel(strings.NewReader(v3Doc))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	endpoint, err := resolver.EndpointFor("acm", "ap-northeast-1")
	if err != nil {
		t.Fatalf("failed to resolve endpoint, %v", err)
	}

	if a, e := endpoint.URL, "https://acm.ap-northeast-1.amazonaws.com"; a != e {
		t.Errorf("expected %q URL got %q", e, a)
	}

	p := resolver.(partitions)[0]

	resolved, err := p.EndpointFor("s3", "us-west-2", func(options *Options) {
		options.UseDualStackEndpoint = DualStackEndpointStateEnabled
	})
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	assertEndpoint(t, resolved, "https://s3.dualstack.us-west-2.amazonaws.com", "s3", "us-west-2")

	ec2metaEndpoint := p.Services["ec2metadata"].Endpoints["aws-global"]
	if a, e := ec2metaEndpoint.Hostname, "169.254.169.254/latest"; a != e {
		t.Errorf("expect ec2metadata host to be %q, got %q", e, a)
	}
}

func assertEndpoint(t *testing.T, endpoint ResolvedEndpoint, expectedURL, expectedSigningName, expectedSigningRegion string) {
	t.Helper()

	if e, a := expectedURL, endpoint.URL; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

	if e, a := expectedSigningName, endpoint.SigningName; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

	if e, a := expectedSigningRegion, endpoint.SigningRegion; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestDecodeEndpoints_NoPartitions(t *testing.T) {
	const doc = `{ "version": 3 }`

	resolver, err := DecodeModel(strings.NewReader(doc))
	if err == nil {
		t.Fatalf("expected error")
	}

	if resolver != nil {
		t.Errorf("expect resolver to be nil")
	}
}

func TestDecodeEndpoints_UnsupportedVersion(t *testing.T) {
	const doc = `{ "version": 2 }`

	resolver, err := DecodeModel(strings.NewReader(doc))
	if err == nil {
		t.Fatalf("expected error decoding model")
	}

	if resolver != nil {
		t.Errorf("expect resolver to be nil")
	}
}

func TestDecodeModelOptionsSet(t *testing.T) {
	var actual DecodeModelOptions
	actual.Set(func(o *DecodeModelOptions) {
		o.SkipCustomizations = true
	})

	expect := DecodeModelOptions{
		SkipCustomizations: true,
	}

	if actual != expect {
		t.Errorf("expect %v options got %v", expect, actual)
	}
}

func TestCustFixAppAutoscalingChina(t *testing.T) {
	const doc = `
{
  "version": 3,
  "partitions": [{
    "defaults" : {
      "hostname" : "{service}.{region}.{dnsSuffix}",
      "protocols" : [ "https" ],
      "signatureVersions" : [ "v4" ]
    },
    "dnsSuffix" : "amazonaws.com.cn",
    "partition" : "aws-cn",
    "partitionName" : "AWS China",
    "regionRegex" : "^cn\\-\\w+\\-\\d+$",
    "regions" : {
      "cn-north-1" : {
        "description" : "China (Beijing)"
      },
      "cn-northwest-1" : {
        "description" : "China (Ningxia)"
      }
    },
    "services" : {
      "application-autoscaling" : {
        "defaults" : {
          "credentialScope" : {
            "service" : "application-autoscaling"
          },
          "hostname" : "autoscaling.{region}.amazonaws.com",
          "protocols" : [ "http", "https" ]
        },
        "endpoints" : {
          "cn-north-1" : { },
          "cn-northwest-1" : { }
        }
      }
	}
  }]
}`

	resolver, err := DecodeModel(strings.NewReader(doc))
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	endpoint, err := resolver.EndpointFor(
		"application-autoscaling", "cn-northwest-1",
	)
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if e, a := `https://autoscaling.cn-northwest-1.amazonaws.com.cn`, endpoint.URL; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestCustFixAppAutoscalingUsGov(t *testing.T) {
	const doc = `
{
  "version": 3,
  "partitions": [{
    "defaults" : {
      "hostname" : "{service}.{region}.{dnsSuffix}",
      "protocols" : [ "https" ],
      "signatureVersions" : [ "v4" ]
    },
    "dnsSuffix" : "amazonaws.com",
    "partition" : "aws-us-gov",
    "partitionName" : "AWS GovCloud (US)",
    "regionRegex" : "^us\\-gov\\-\\w+\\-\\d+$",
    "regions" : {
      "us-gov-east-1" : {
        "description" : "AWS GovCloud (US-East)"
      },
      "us-gov-west-1" : {
        "description" : "AWS GovCloud (US)"
      }
    },
    "services" : {
      "application-autoscaling" : {
        "endpoints" : {
          "us-gov-east-1" : { },
          "us-gov-west-1" : { }
        }
      }
	}
  }]
}`

	resolver, err := DecodeModel(strings.NewReader(doc))
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	endpoint, err := resolver.EndpointFor(
		"application-autoscaling", "us-gov-west-1",
	)
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if e, a := `https://autoscaling.us-gov-west-1.amazonaws.com`, endpoint.URL; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}
