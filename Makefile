generate:
	rm -rfv ./gen/**/*.go
	go install ./cmd/...
	aws-go-generate AutoScaling apis/autoscaling/2011-01-01.api.json > gen/autoscaling/autoscaling.go
	aws-go-generate CloudFormation apis/cloudformation/2010-05-15.api.json > gen/cloudformation/cloudformation.go
	aws-go-generate CloudTrail apis/cloudtrail/2013-11-01.api.json > gen/cloudtrail/cloudtrail.go
	aws-go-generate CloudSearch apis/cloudsearch/2013-01-01.api.json > gen/cloudsearch/cloudsearch.go
	aws-go-generate CloudWatch apis/cloudwatch/2010-08-01.api.json > gen/cloudwatch/cloudwatch.go
	aws-go-generate CogitoIdentity apis/cognito-identity/2014-06-30.api.json > gen/cognito/identity/identity.go
	aws-go-generate CodeDeploy apis/codedeploy/2014-10-06.api.json > gen/codedeploy/codedeploy.go
	aws-go-generate Config apis/config/2014-11-12.api.json > gen/config/config.go
	aws-go-generate DataPipeline apis/datapipeline/2012-10-29.api.json > gen/datapipeline/datapipeline.go
	aws-go-generate DirectConnect apis/directconnect/2012-10-25.api.json > gen/directconnect/directconnect.go
	aws-go-generate DynamoDB apis/dynamodb/2012-08-10.api.json > gen/dynamodb/dynamodb.go
	aws-go-generate ElasticCache apis/elasticache/2014-09-30.api.json > gen/elasticache/elasticache.go
	aws-go-generate ElasticBeanstalk apis/elasticbeanstalk/2010-12-01.api.json > gen/elasticbeanstalk/elasticbeanstalk.go
	aws-go-generate ELB apis/elb/2012-06-01.api.json > gen/elb/elb.go
	aws-go-generate EMR apis/emr/2009-03-31.api.json > gen/emr/emr.go
	aws-go-generate IAM apis/iam/2010-05-08.api.json > gen/iam/iam.go
	aws-go-generate ImportExport apis/importexport/2010-06-01.api.json > gen/importexport/importexport.go
	aws-go-generate Kinesis apis/kinesis/2013-12-02.api.json > gen/kinesis/kinesis.go
	aws-go-generate KMS apis/kms/2014-11-01.api.json > gen/kms/kms.go
	aws-go-generate Logs apis/logs/2014-03-28.api.json > gen/logs/logs.go
	aws-go-generate OpsWorks apis/opsworks/2013-02-18.api.json > gen/opsworks/opsworks.go
	aws-go-generate RDS apis/rds/2014-09-01.api.json > gen/rds/rds.go
	aws-go-generate RedShift apis/redshift/2012-12-01.api.json > gen/redshift/redshift.go
	aws-go-generate Route53Domains apis/route53domains/2014-05-15.api.json > gen/route53domains/route53domains.go
	aws-go-generate SDB apis/sdb/2009-04-15.api.json > gen/sdb/sdb.go
	aws-go-generate SES apis/ses/2010-12-01.api.json > gen/ses/ses.go
	aws-go-generate SNS apis/sns/2010-03-31.api.json > gen/sns/sns.go
	aws-go-generate SQS apis/sqs/2012-11-05.api.json > gen/sqs/sqs.go
	aws-go-generate StorageGateway apis/storagegateway/2013-06-30.api.json > gen/storagegateway/storagegateway.go
	aws-go-generate STS apis/sts/2011-06-15.api.json > gen/sts/sts.go
	aws-go-generate Support apis/support/2013-04-15.api.json > gen/support/support.go
	aws-go-generate SWF apis/swf/2012-01-25.api.json > gen/swf/swf.go
	go install ./...
