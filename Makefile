default: generate

clean:
	rm -rfv aws/gen/**/*.go

install-gen:
	go install ./cmd/...

generate: clean install-gen
	aws-gen-goendpoints apis/_endpoints.json > aws/gen/endpoints/endpoints.go
	aws-gen-gocli AutoScaling apis/autoscaling/2011-01-01.api.json > aws/gen/autoscaling/autoscaling.go
	aws-gen-gocli CloudFormation apis/cloudformation/2010-05-15.api.json > aws/gen/cloudformation/cloudformation.go
	aws-gen-gocli CloudFront apis/cloudfront/2014-10-21.api.json > aws/gen/cloudfront/cloudfront.go
	aws-gen-gocli CloudTrail apis/cloudtrail/2013-11-01.api.json > aws/gen/cloudtrail/cloudtrail.go
	aws-gen-gocli CloudSearch apis/cloudsearch/2013-01-01.api.json > aws/gen/cloudsearch/cloudsearch.go
	aws-gen-gocli CloudWatch apis/cloudwatch/2010-08-01.api.json > aws/gen/cloudwatch/cloudwatch.go
	aws-gen-gocli CogitoIdentity apis/cognito-identity/2014-06-30.api.json > aws/gen/cognito/identity/identity.go
	aws-gen-gocli CodeDeploy apis/codedeploy/2014-10-06.api.json > aws/gen/codedeploy/codedeploy.go
	aws-gen-gocli Config apis/config/2014-11-12.api.json > aws/gen/config/config.go
	aws-gen-gocli DataPipeline apis/datapipeline/2012-10-29.api.json > aws/gen/datapipeline/datapipeline.go
	aws-gen-gocli DirectConnect apis/directconnect/2012-10-25.api.json > aws/gen/directconnect/directconnect.go
	aws-gen-gocli DynamoDB apis/dynamodb/2012-08-10.api.json > aws/gen/dynamodb/dynamodb.go
	aws-gen-gocli EC2 apis/ec2/2014-10-01.api.json > aws/gen/ec2/ec2.go
	aws-gen-gocli ElasticCache apis/elasticache/2014-09-30.api.json > aws/gen/elasticache/elasticache.go
	aws-gen-gocli ElasticBeanstalk apis/elasticbeanstalk/2010-12-01.api.json > aws/gen/elasticbeanstalk/elasticbeanstalk.go
	aws-gen-gocli ELB apis/elb/2012-06-01.api.json > aws/gen/elb/elb.go
	aws-gen-gocli EMR apis/emr/2009-03-31.api.json > aws/gen/emr/emr.go
	aws-gen-gocli IAM apis/iam/2010-05-08.api.json > aws/gen/iam/iam.go
	aws-gen-gocli ImportExport apis/importexport/2010-06-01.api.json > aws/gen/importexport/importexport.go
	aws-gen-gocli Kinesis apis/kinesis/2013-12-02.api.json > aws/gen/kinesis/kinesis.go
	aws-gen-gocli KMS apis/kms/2014-11-01.api.json > aws/gen/kms/kms.go
	aws-gen-gocli Logs apis/logs/2014-03-28.api.json > aws/gen/logs/logs.go
	aws-gen-gocli OpsWorks apis/opsworks/2013-02-18.api.json > aws/gen/opsworks/opsworks.go
	aws-gen-gocli RDS apis/rds/2014-09-01.api.json > aws/gen/rds/rds.go
	aws-gen-gocli RedShift apis/redshift/2012-12-01.api.json > aws/gen/redshift/redshift.go
	aws-gen-gocli Route53 apis/route53/2013-04-01.api.json > aws/gen/route53/route53.go
	aws-gen-gocli Route53Domains apis/route53domains/2014-05-15.api.json > aws/gen/route53domains/route53domains.go
	aws-gen-gocli S3 apis/s3/2006-03-01.api.json > aws/gen/s3/s3.go
	aws-gen-gocli SDB apis/sdb/2009-04-15.api.json > aws/gen/sdb/sdb.go
	aws-gen-gocli SES apis/ses/2010-12-01.api.json > aws/gen/ses/ses.go
	aws-gen-gocli SNS apis/sns/2010-03-31.api.json > aws/gen/sns/sns.go
	aws-gen-gocli SQS apis/sqs/2012-11-05.api.json > aws/gen/sqs/sqs.go
	aws-gen-gocli StorageGateway apis/storagegateway/2013-06-30.api.json > aws/gen/storagegateway/storagegateway.go
	aws-gen-gocli STS apis/sts/2011-06-15.api.json > aws/gen/sts/sts.go
	aws-gen-gocli Support apis/support/2013-04-15.api.json > aws/gen/support/support.go
	aws-gen-gocli SWF apis/swf/2012-01-25.api.json > aws/gen/swf/swf.go
	go install ./...
