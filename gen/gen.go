// Package gen contains automatically generated AWS clients.
package gen

//go:generate aws-gen-goendpoints ../apis/_endpoints.json endpoints/endpoints.go
//go:generate aws-gen-gocli AutoScaling ../apis/autoscaling/2011-01-01.api.json autoscaling/autoscaling.go
//go:generate aws-gen-gocli CloudFormation ../apis/cloudformation/2010-05-15.api.json cloudformation/cloudformation.go
//go:generate aws-gen-gocli CloudFront ../apis/cloudfront/2014-10-21.api.json cloudfront/cloudfront.go
//go:generate aws-gen-gocli CloudTrail ../apis/cloudtrail/2013-11-01.api.json cloudtrail/cloudtrail.go
//go:generate aws-gen-gocli CloudSearch ../apis/cloudsearch/2013-01-01.api.json cloudsearch/cloudsearch.go
//go:generate aws-gen-gocli CloudSearchDomain ../apis/cloudsearchdomain/2013-01-01.api.json cloudsearchdomain/cloudsearchdomain.go
//go:generate aws-gen-gocli CloudWatch ../apis/cloudwatch/2010-08-01.api.json cloudwatch/cloudwatch.go
//go:generate aws-gen-gocli CognitoIdentity ../apis/cognito-identity/2014-06-30.api.json cognito/identity/identity.go
//go:generate aws-gen-gocli CognitoSync ../apis/cognito-sync/2014-06-30.api.json cognito/sync/sync.go
//go:generate aws-gen-gocli CodeDeploy ../apis/codedeploy/2014-10-06.api.json codedeploy/codedeploy.go
//go:generate aws-gen-gocli Config ../apis/config/2014-11-12.api.json config/config.go
//go:generate aws-gen-gocli DataPipeline ../apis/datapipeline/2012-10-29.api.json datapipeline/datapipeline.go
//go:generate aws-gen-gocli DirectConnect ../apis/directconnect/2012-10-25.api.json directconnect/directconnect.go
//go:generate aws-gen-gocli DynamoDB ../apis/dynamodb/2012-08-10.api.json dynamodb/dynamodb.go
//go:generate aws-gen-gocli EC2 ../apis/ec2/2014-10-01.api.json ec2/ec2.go
//go:generate aws-gen-gocli ElasticCache ../apis/elasticache/2014-09-30.api.json elasticache/elasticache.go
//go:generate aws-gen-gocli ElasticBeanstalk ../apis/elasticbeanstalk/2010-12-01.api.json elasticbeanstalk/elasticbeanstalk.go
//go:generate aws-gen-gocli ElasticTranscoder ../apis/elastictranscoder/2012-09-25.api.json elastictranscoder/elastictranscoder.go
//go:generate aws-gen-gocli ELB ../apis/elb/2012-06-01.api.json elb/elb.go
//go:generate aws-gen-gocli EMR ../apis/emr/2009-03-31.api.json emr/emr.go
//go:generate aws-gen-gocli IAM ../apis/iam/2010-05-08.api.json iam/iam.go
//go:generate aws-gen-gocli ImportExport ../apis/importexport/2010-06-01.api.json importexport/importexport.go
//go:generate aws-gen-gocli Kinesis ../apis/kinesis/2013-12-02.api.json kinesis/kinesis.go
//go:generate aws-gen-gocli KMS ../apis/kms/2014-11-01.api.json kms/kms.go
//go:generate aws-gen-gocli Lambda ../apis/lambda/2014-11-11.api.json lambda/lambda.go
//go:generate aws-gen-gocli Logs ../apis/logs/2014-03-28.api.json logs/logs.go
//go:generate aws-gen-gocli OpsWorks ../apis/opsworks/2013-02-18.api.json opsworks/opsworks.go
//go:generate aws-gen-gocli RDS ../apis/rds/2014-09-01.api.json rds/rds.go
//go:generate aws-gen-gocli RedShift ../apis/redshift/2012-12-01.api.json redshift/redshift.go
//go:generate aws-gen-gocli Route53 ../apis/route53/2013-04-01.api.json route53/route53.go
//go:generate aws-gen-gocli Route53Domains ../apis/route53domains/2014-05-15.api.json route53domains/route53domains.go
//go:generate aws-gen-gocli S3 ../apis/s3/2006-03-01.api.json s3/s3.go
//go:generate aws-gen-gocli SDB ../apis/sdb/2009-04-15.api.json sdb/sdb.go
//go:generate aws-gen-gocli SES ../apis/ses/2010-12-01.api.json ses/ses.go
//go:generate aws-gen-gocli SNS ../apis/sns/2010-03-31.api.json sns/sns.go
//go:generate aws-gen-gocli SQS ../apis/sqs/2012-11-05.api.json sqs/sqs.go
//go:generate aws-gen-gocli StorageGateway ../apis/storagegateway/2013-06-30.api.json storagegateway/storagegateway.go
//go:generate aws-gen-gocli STS ../apis/sts/2011-06-15.api.json sts/sts.go
//go:generate aws-gen-gocli Support ../apis/support/2013-04-15.api.json support/support.go
//go:generate aws-gen-gocli SWF ../apis/swf/2012-01-25.api.json swf/swf.go
