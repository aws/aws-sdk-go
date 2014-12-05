generate:
	rm -rfv ./gen/**/*.go
	go install ./cmd/...
	aws-go-generate CloudTrail apis/cloudtrail/2013-11-01.api.json > gen/cloudtrail/cloudtrail.go
	aws-go-generate CogitoIdentity apis/cognito-identity/2014-06-30.api.json > gen/cognito/identity/identity.go
	aws-go-generate CodeDeploy apis/codedeploy/2014-10-06.api.json > gen/codedeploy/codedeploy.go
	aws-go-generate Config apis/config/2014-11-12.api.json > gen/config/config.go
	aws-go-generate DataPipeline apis/datapipeline/2012-10-29.api.json > gen/datapipeline/datapipeline.go
	aws-go-generate DirectConnect apis/directconnect/2012-10-25.api.json > gen/directconnect/directconnect.go
	aws-go-generate DynamoDB apis/dynamodb/2012-08-10.api.json > gen/dynamodb/dynamodb.go
	aws-go-generate EMR apis/emr/2009-03-31.api.json > gen/emr/emr.go
	aws-go-generate Kinesis apis/kinesis/2013-12-02.api.json > gen/kinesis/kinesis.go
	aws-go-generate KMS apis/kms/2014-11-01.api.json > gen/kms/kms.go
	aws-go-generate Logs apis/logs/2014-03-28.api.json > gen/logs/logs.go
	aws-go-generate OpsWorks apis/opsworks/2013-02-18.api.json > gen/opsworks/opsworks.go
	aws-go-generate Route53Domains apis/route53domains/2014-05-15.api.json > gen/route53domains/route53domains.go
	aws-go-generate StorageGateway apis/storagegateway/2013-06-30.api.json > gen/storagegateway/storagegateway.go
	aws-go-generate Support apis/support/2013-04-15.api.json > gen/support/support.go
	aws-go-generate SWF apis/swf/2012-01-25.api.json > gen/swf/swf.go
	go install ./...
