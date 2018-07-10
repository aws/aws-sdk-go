package ec2metadata

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/fullsailor/pkcs7"
)

const (
	// AWSCertificatePem is the official public certificate for AWS
	// copied from https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html
	AWSCertificatePem = `-----BEGIN CERTIFICATE-----
MIIC7TCCAq0CCQCWukjZ5V4aZzAJBgcqhkjOOAQDMFwxCzAJBgNVBAYTAlVTMRkw
FwYDVQQIExBXYXNoaW5ndG9uIFN0YXRlMRAwDgYDVQQHEwdTZWF0dGxlMSAwHgYD
VQQKExdBbWF6b24gV2ViIFNlcnZpY2VzIExMQzAeFw0xMjAxMDUxMjU2MTJaFw0z
ODAxMDUxMjU2MTJaMFwxCzAJBgNVBAYTAlVTMRkwFwYDVQQIExBXYXNoaW5ndG9u
IFN0YXRlMRAwDgYDVQQHEwdTZWF0dGxlMSAwHgYDVQQKExdBbWF6b24gV2ViIFNl
cnZpY2VzIExMQzCCAbcwggEsBgcqhkjOOAQBMIIBHwKBgQCjkvcS2bb1VQ4yt/5e
ih5OO6kK/n1Lzllr7D8ZwtQP8fOEpp5E2ng+D6Ud1Z1gYipr58Kj3nssSNpI6bX3
VyIQzK7wLclnd/YozqNNmgIyZecN7EglK9ITHJLP+x8FtUpt3QbyYXJdmVMegN6P
hviYt5JH/nYl4hh3Pa1HJdskgQIVALVJ3ER11+Ko4tP6nwvHwh6+ERYRAoGBAI1j
k+tkqMVHuAFcvAGKocTgsjJem6/5qomzJuKDmbJNu9Qxw3rAotXau8Qe+MBcJl/U
hhy1KHVpCGl9fueQ2s6IL0CaO/buycU1CiYQk40KNHCcHfNiZbdlx1E9rpUp7bnF
lRa2v1ntMX3caRVDdbtPEWmdxSCYsYFDk4mZrOLBA4GEAAKBgEbmeve5f8LIE/Gf
MNmP9CM5eovQOGx5ho8WqD+aTebs+k2tn92BBPqeZqpWRa5P/+jrdKml1qx4llHW
MXrs3IgIb6+hUIB+S8dz8/mmO0bpr76RoZVCXYab2CZedFut7qc3WUH9+EUAH5mw
vSeDCOUMYQR7R9LINYwouHIziqQYMAkGByqGSM44BAMDLwAwLAIUWXBlk40xTwSw
7HX32MxXYruse9ACFBNGmdX2ZBrVNGrN9N2f6ROk0k9K
-----END CERTIFICATE-----`
)

// GetMetadata uses the path provided to request information from the EC2
// instance metdata service. The content will be returned as a string, or
// error if the request failed.
func (c *EC2Metadata) GetMetadata(p string) (string, error) {
	op := &request.Operation{
		Name:       "GetMetadata",
		HTTPMethod: "GET",
		HTTPPath:   path.Join("/", "meta-data", p),
	}

	output := &metadataOutput{}
	req := c.NewRequest(op, nil, output)

	return output.Content, req.Send()
}

// GetUserData returns the userdata that was configured for the service. If
// there is no user-data setup for the EC2 instance a "NotFoundError" error
// code will be returned.
func (c *EC2Metadata) GetUserData() (string, error) {
	op := &request.Operation{
		Name:       "GetUserData",
		HTTPMethod: "GET",
		HTTPPath:   path.Join("/", "user-data"),
	}

	output := &metadataOutput{}
	req := c.NewRequest(op, nil, output)
	req.Handlers.UnmarshalError.PushBack(func(r *request.Request) {
		if r.HTTPResponse.StatusCode == http.StatusNotFound {
			r.Error = awserr.New("NotFoundError", "user-data not found", r.Error)
		}
	})

	return output.Content, req.Send()
}

// GetDynamicData uses the path provided to request information from the EC2
// instance metadata service for dynamic data. The content will be returned
// as a string, or error if the request failed.
func (c *EC2Metadata) GetDynamicData(p string) (string, error) {
	op := &request.Operation{
		Name:       "GetDynamicData",
		HTTPMethod: "GET",
		HTTPPath:   path.Join("/", "dynamic", p),
	}

	output := &metadataOutput{}
	req := c.NewRequest(op, nil, output)

	return output.Content, req.Send()
}

// Parses a PEM-encoded certificate
func parseCertificate(certPem string) (*x509.Certificate, error) {
	dec, _ := pem.Decode([]byte(certPem))
	if dec == nil {
		return nil, fmt.Errorf("failed to find any PEM-encoded data block")
	}
	return x509.ParseCertificate(dec.Bytes)
}

// GetInstanceIdentityDocument retrieves an identity document describing an
// instance and verifies its signature with provided certificate.
// Error is returned if failed to parse the certificate, or the request fails,
// or is unable to parse the response, or the signature verification fails.
func (c *EC2Metadata) GetInstanceIdentityDocument(certPem string) (EC2InstanceIdentityDocument, error) {
	cert, err := parseCertificate(certPem)
	if err != nil {
		return EC2InstanceIdentityDocument{},
			awserr.New("CertificateParsingError",
				"failed to parse the PEM-encoded certificate", err)
	}

	resp, err := c.GetDynamicData("instance-identity/pkcs7")
	if err != nil {
		return EC2InstanceIdentityDocument{},
			awserr.New("EC2MetadataRequestError",
				"failed to get EC2 instance PKCS7 signature", err)
	}

	dec, err := base64.StdEncoding.DecodeString(resp)
	if err != nil {
		return EC2InstanceIdentityDocument{},
			awserr.New("Base64DecodeError",
				"failed to base64-decode PKCS7 resopnse", err)
	}

	parsed, err := pkcs7.Parse(dec)
	if err != nil {
		return EC2InstanceIdentityDocument{},
			awserr.New("EC2MetadataRequestError",
				"failed to parse PKCS7 resopnse", err)
	}

	parsed.Certificates = []*x509.Certificate{cert}
	if err := parsed.Verify(); err != nil {
		return EC2InstanceIdentityDocument{},
			awserr.New("EC2MetadataVerificationError",
				"failed to verify PKCS7 signature with certificate", err)
	}

	doc := EC2InstanceIdentityDocument{}
	if err := json.NewDecoder(bytes.NewReader(parsed.Content)).Decode(&doc); err != nil {
		return EC2InstanceIdentityDocument{},
			awserr.New("SerializationError",
				"failed to decode EC2 instance identity document", err)
	}

	return doc, nil
}

// IAMInfo retrieves IAM info from the metadata API
func (c *EC2Metadata) IAMInfo() (EC2IAMInfo, error) {
	resp, err := c.GetMetadata("iam/info")
	if err != nil {
		return EC2IAMInfo{},
			awserr.New("EC2MetadataRequestError",
				"failed to get EC2 IAM info", err)
	}

	info := EC2IAMInfo{}
	if err := json.NewDecoder(strings.NewReader(resp)).Decode(&info); err != nil {
		return EC2IAMInfo{},
			awserr.New("SerializationError",
				"failed to decode EC2 IAM info", err)
	}

	if info.Code != "Success" {
		errMsg := fmt.Sprintf("failed to get EC2 IAM Info (%s)", info.Code)
		return EC2IAMInfo{},
			awserr.New("EC2MetadataError", errMsg, nil)
	}

	return info, nil
}

// Region returns the region the instance is running in.
func (c *EC2Metadata) Region() (string, error) {
	resp, err := c.GetMetadata("placement/availability-zone")
	if err != nil {
		return "", err
	}

	// returns region without the suffix. Eg: us-west-2a becomes us-west-2
	return resp[:len(resp)-1], nil
}

// Available returns if the application has access to the EC2 Metadata service.
// Can be used to determine if application is running within an EC2 Instance and
// the metadata service is available.
func (c *EC2Metadata) Available() bool {
	if _, err := c.GetMetadata("instance-id"); err != nil {
		return false
	}

	return true
}

// An EC2IAMInfo provides the shape for unmarshaling
// an IAM info from the metadata API
type EC2IAMInfo struct {
	Code               string
	LastUpdated        time.Time
	InstanceProfileArn string
	InstanceProfileID  string
}

// An EC2InstanceIdentityDocument provides the shape for unmarshaling
// an instance identity document
type EC2InstanceIdentityDocument struct {
	DevpayProductCodes []string  `json:"devpayProductCodes"`
	AvailabilityZone   string    `json:"availabilityZone"`
	PrivateIP          string    `json:"privateIp"`
	Version            string    `json:"version"`
	Region             string    `json:"region"`
	InstanceID         string    `json:"instanceId"`
	BillingProducts    []string  `json:"billingProducts"`
	InstanceType       string    `json:"instanceType"`
	AccountID          string    `json:"accountId"`
	PendingTime        time.Time `json:"pendingTime"`
	ImageID            string    `json:"imageId"`
	KernelID           string    `json:"kernelId"`
	RamdiskID          string    `json:"ramdiskId"`
	Architecture       string    `json:"architecture"`
}
