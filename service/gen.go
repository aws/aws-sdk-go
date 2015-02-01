// Package service contains automatically generated AWS clients.
package service

//go:generate ../../../../../bin/aws-gen-goendpoints ../apis/_endpoints.json endpoints/endpoints.go
//go:generate ../../../../../bin/aws-gen-gocli ../apis/autoscaling/2011-01-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/cloudformation/2010-05-15.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/cloudfront/2014-10-21.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/cloudtrail/2013-11-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/cloudsearch/2013-01-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/cloudsearchdomain/2013-01-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/cloudwatch/2010-08-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/cognito-identity/2014-06-30.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/cognito-sync/2014-06-30.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/codedeploy/2014-10-06.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/config/2014-11-12.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/datapipeline/2012-10-29.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/directconnect/2012-10-25.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/dynamodb/2012-08-10.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/ec2/2014-10-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/elasticache/2014-09-30.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/elasticbeanstalk/2010-12-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/elastictranscoder/2012-09-25.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/elb/2012-06-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/emr/2009-03-31.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/iam/2010-05-08.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/importexport/2010-06-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/kinesis/2013-12-02.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/kms/2014-11-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/lambda/2014-11-11.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/logs/2014-03-28.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/opsworks/2013-02-18.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/rds/2014-09-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/redshift/2012-12-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/route53/2013-04-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/route53domains/2014-05-15.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/s3/2006-03-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/sdb/2009-04-15.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/ses/2010-12-01.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/sns/2010-03-31.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/sqs/2012-11-05.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/storagegateway/2013-06-30.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/sts/2011-06-15.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/support/2013-04-15.normal.json
//go:generate ../../../../../bin/aws-gen-gocli ../apis/swf/2012-01-25.normal.json
