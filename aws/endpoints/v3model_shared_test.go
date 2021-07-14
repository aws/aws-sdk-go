// +build go1.9

package endpoints

import "regexp"

var testPartitions = partitions{
	partition{
		ID:                 "part-id",
		Name:               "partitionName",
		DNSSuffix:          "amazonaws.com",
		DualStackDNSSuffix: "aws",
		RegionRegex: regionRegex{
			Regexp: func() *regexp.Regexp {
				reg, _ := regexp.Compile("^(us|eu|ap|sa|ca)\\-\\w+\\-\\d+$")
				return reg
			}(),
		},
		Defaults: endpoint{
			Hostname:          "{service}.{region}.{dnsSuffix}",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		DualStackDefaults: endpoint{
			Hostname:          "{service}.{region}.{dualstackDnsSuffix}",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		Regions: regions{
			"us-east-1": region{
				Description: "region description",
			},
			"us-west-2": region{},
		},
		Services: services{
			"s3": service{
				Defaults: endpoint{
					CredentialScope: credentialScope{
						Service: "s3",
					},
				},
				DualStackDefaults: endpoint{
					Hostname:          "{service}.dualstack.{region}.{dualstackDnsSuffix}",
					Protocols:         []string{"https"},
					SignatureVersions: []string{"v4"},
					CredentialScope: credentialScope{
						Service: "s3",
					},
				},
				DualStackDNSSuffix: "amazonaws.com",
				DualStackEndpoints: endpoints{
					"us-west-1": endpoint{
						Hostname: "s3-foo.dualstack.us-west-1.amazonaws.com",
					},
				},
			},
			"s3-control": service{
				Defaults: endpoint{
					CredentialScope: credentialScope{
						Service: "s3-control",
					},
				},
				DualStackDefaults: endpoint{
					Hostname:          "{service}.dualstack.{region}.{dualstackDnsSuffix}",
					Protocols:         []string{"https"},
					SignatureVersions: []string{"v4"},
					CredentialScope: credentialScope{
						Service: "s3-control",
					},
				},
				DualStackDNSSuffix: "amazonaws.com",
				DualStackEndpoints: endpoints{
					"us-west-1": endpoint{
						Hostname: "s3-control-foo.dualstack.us-west-1.amazonaws.com",
					},
				},
			},
			"service1": service{
				Defaults: endpoint{
					CredentialScope: credentialScope{
						Service: "service1",
					},
				},
				Endpoints: endpoints{
					"us-east-1": {},
					"us-west-2": {},
				},
			},
			"service2": service{
				Defaults: endpoint{
					CredentialScope: credentialScope{
						Service: "service2",
					},
				},
			},
			"httpService": service{
				Defaults: endpoint{
					Protocols: []string{"http"},
				},
			},
			"globalService": service{
				IsRegionalized:    boxedFalse,
				PartitionEndpoint: "aws-global",
				Endpoints: endpoints{
					"aws-global": endpoint{
						CredentialScope: credentialScope{
							Region: "us-east-1",
						},
						Hostname: "globalService.amazonaws.com",
					},
					"fips-aws-global": endpoint{
						CredentialScope: credentialScope{
							Region: "us-east-1",
						},
						Hostname: "globalService-fips.amazonaws.com",
					},
				},
				DualStackEndpoints: endpoints{
					"aws-global": endpoint{
						CredentialScope: credentialScope{
							Region: "us-east-1",
						},
						Hostname: "globalService.global.aws",
					},
					"fips-aws-global": endpoint{
						CredentialScope: credentialScope{
							Region: "us-east-1",
						},
						Hostname: "globalService-fips.global.aws",
					},
				},
			},
		},
	},
	partition{
		ID:        "part-id-2",
		Name:      "partitionNumber2",
		DNSSuffix: "amazonaws.com.cn",
		RegionRegex: regionRegex{
			Regexp: func() *regexp.Regexp {
				reg, _ := regexp.Compile("^(cn)\\-\\w+\\-\\d+$")
				return reg
			}(),
		},
		Defaults: endpoint{
			Hostname:          "{service}.{region}.{dnsSuffix}",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		Services: services{
			"service1": service{
				Defaults: endpoint{
					CredentialScope: credentialScope{
						Service: "s3",
					},
				},
				DualStackEndpoints: map[string]endpoint{
					"cn-north-1": {
						Hostname: "service1.cn-north-1.aws.cn",
						CredentialScope: credentialScope{
							Service: "service1",
							Region:  "cn-north-1",
						},
					},
				},
			},
		},
	},
}
