package s3

// http://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints
var hostedZoneIDsMap = map[string]string{
	"us-east-1":      "Z3AQBSTGFYJSTF",
	"us-west-2":      "Z3BJ6K6RIION7M",
	"us-west-1":      "Z2F56UZL2M1ACD",
	"eu-west-1":      "Z1BKCTXD74EZPE",
	"central-1":      "Z21DNDUVLTQW6Q",
	"ap-southeast-1": "Z3O0J2DXBE1FTB",
	"ap-southeast-2": "Z1WCIGYICN2BYD",
	"ap-northeast-1": "Z2M4EHUR26P7ZW",
	"sa-east-1":      "Z7KQH4QJS55SO",
	"us-gov-west-1":  "Z31GFT0UA1I2HV",
}

func HostedZoneIDForRegion(region string) string {
	if region == "" {
		region = "us-east-1"
	}

	return hostedZoneIDsMap[region]
}
