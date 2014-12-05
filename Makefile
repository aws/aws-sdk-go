build:
	go install ./...
	aws-go-generate CloudTrail apis/cloudtrail/2013-11-01.api.json > cloudtrail/cloudtrail.go
	aws-go-generate DynamoDB apis/dynamodb/2012-08-10.api.json > dynamodb/dynamodb.go
	aws-go-generate KMS apis/kms/2014-11-01.api.json > kms/kms.go
	aws-go-generate Support apis/support/2013-04-15.api.json > support/support.go
