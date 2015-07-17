package dynamodb_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/internal/test/unit"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

var _ = unit.Imported
var db *dynamodb.DynamoDB

func TestMain(m *testing.M) {
	db = dynamodb.New(&aws.Config{
		MaxRetries: aws.Int(2),
	})
	db.Handlers.Send.Clear() // mock sending

	os.Exit(m.Run())
}

func mockCRCResponse(svc *dynamodb.DynamoDB, status int, body, crc string) (req *aws.Request) {
	header := http.Header{}
	header.Set("x-amz-crc32", crc)

	req, _ = svc.ListTablesRequest(nil)
	req.Handlers.Send.PushBack(func(*aws.Request) {
		req.HTTPResponse = &http.Response{
			StatusCode: status,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(body))),
			Header:     header,
		}
	})
	req.Send()
	return
}

func TestCustomRetryRules(t *testing.T) {
	d := dynamodb.New(&aws.Config{MaxRetries: aws.Int(-1)})
	assert.Equal(t, d.MaxRetries(), uint(10))
}

func TestValidateCRC32NoHeaderSkip(t *testing.T) {
	req := mockCRCResponse(db, 200, "{}", "")
	assert.NoError(t, req.Error)
}

func TestValidateCRC32InvalidHeaderSkip(t *testing.T) {
	req := mockCRCResponse(db, 200, "{}", "ABC")
	assert.NoError(t, req.Error)
}

func TestValidateCRC32AlreadyErrorSkip(t *testing.T) {
	req := mockCRCResponse(db, 400, "{}", "1234")
	assert.Error(t, req.Error)

	assert.NotEqual(t, "CRC32CheckFailed", req.Error.(awserr.Error).Code())
}

func TestValidateCRC32IsValid(t *testing.T) {
	req := mockCRCResponse(db, 200, `{"TableNames":["A"]}`, "3090163698")
	assert.NoError(t, req.Error)

	// CRC check does not affect output parsing
	out := req.Data.(*dynamodb.ListTablesOutput)
	assert.Equal(t, "A", *out.TableNames[0])
}

func TestValidateCRC32DoesNotMatch(t *testing.T) {
	req := mockCRCResponse(db, 200, "{}", "1234")
	assert.Error(t, req.Error)

	assert.Equal(t, "CRC32CheckFailed", req.Error.(awserr.Error).Code())
	assert.Equal(t, 2, int(req.RetryCount))
}

func TestValidateCRC32DoesNotMatchNoComputeChecksum(t *testing.T) {
	svc := dynamodb.New(&aws.Config{
		MaxRetries:              aws.Int(2),
		DisableComputeChecksums: aws.Boolean(true),
	})
	svc.Handlers.Send.Clear() // mock sending

	req := mockCRCResponse(svc, 200, `{"TableNames":["A"]}`, "1234")
	assert.NoError(t, req.Error)

	assert.Equal(t, 0, int(req.RetryCount))

	// CRC check disabled. Does not affect output parsing
	out := req.Data.(*dynamodb.ListTablesOutput)
	assert.Equal(t, "A", *out.TableNames[0])
}
