package v4util

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
)

func TestSignWithHeader(t *testing.T) {
	server := setupServer()
	defer server.Close()

	creds := aws.Creds("DUMMY_KEY", "DUMMY_SECRET", "")
	config := &dynamodb.DynamoDBConfig{
		Config: &aws.Config{
			Credentials: creds,
			Endpoint:    server.URL,
		},
	}
	db := dynamodb.New(config)
	db.Handlers.Sign.PushBack(SignWithHeader)

	_, err := db.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignWithHeaderFailed(t *testing.T) {
	server := setupServer()
	defer server.Close()

	creds := aws.Creds("DUMMY_KEY", "DUMMY_SECRET", "")
	config := &dynamodb.DynamoDBConfig{
		Config: &aws.Config{
			Credentials: creds,
			Endpoint:    server.URL,
		},
	}
	db := dynamodb.New(config)
	db.Handlers.Sign.Init() // remove the v4.Sign handler
	db.Handlers.Sign.PushBack(SignWithHeader)

	_, err := db.ListTables(&dynamodb.ListTablesInput{})
	if err == nil {
		t.Fatalf("Expected failure")
	}
}

// setupServer creates test server which simply validates the presence of the
// Authorization header.
func setupServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	return server
}