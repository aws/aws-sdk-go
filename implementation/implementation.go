package implementation

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/accessanalyzer"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/aws/aws-sdk-go/service/acmpca"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness"
	"github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/amplifybackend"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/appconfig"
	"github.com/aws/aws-sdk-go/service/appflow"
	"github.com/aws/aws-sdk-go/service/appintegrationsservice"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go/service/applicationcostprofiler"
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice"
	"github.com/aws/aws-sdk-go/service/applicationinsights"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/aws/aws-sdk-go/service/appregistry"
	"github.com/aws/aws-sdk-go/service/apprunner"
	"github.com/aws/aws-sdk-go/service/appstream"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/auditmanager"
	"github.com/aws/aws-sdk-go/service/augmentedairuntime"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
	"github.com/aws/aws-sdk-go/service/backup"
	"github.com/aws/aws-sdk-go/service/batch"
	"github.com/aws/aws-sdk-go/service/braket"
	"github.com/aws/aws-sdk-go/service/budgets"
	"github.com/aws/aws-sdk-go/service/chime"
	"github.com/aws/aws-sdk-go/service/cloud9"
	"github.com/aws/aws-sdk-go/service/clouddirectory"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/cloudhsm"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go/service/cloudsearch"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/codeartifact"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/aws/aws-sdk-go/service/codecommit"
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"github.com/aws/aws-sdk-go/service/codeguruprofiler"
	"github.com/aws/aws-sdk-go/service/codegurureviewer"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/aws/aws-sdk-go/service/codestar"
	"github.com/aws/aws-sdk-go/service/codestarconnections"
	"github.com/aws/aws-sdk-go/service/codestarnotifications"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitosync"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/service/comprehendmedical"
	"github.com/aws/aws-sdk-go/service/computeoptimizer"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/aws/aws-sdk-go/service/connectcontactlens"
	"github.com/aws/aws-sdk-go/service/connectparticipant"
	"github.com/aws/aws-sdk-go/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/customerprofiles"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/dataexchange"
	"github.com/aws/aws-sdk-go/service/datapipeline"
	"github.com/aws/aws-sdk-go/service/datasync"
	"github.com/aws/aws-sdk-go/service/dax"
	"github.com/aws/aws-sdk-go/service/detective"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/service/devopsguru"
	"github.com/aws/aws-sdk-go/service/directconnect"
	"github.com/aws/aws-sdk-go/service/directoryservice"
	"github.com/aws/aws-sdk-go/service/dlm"
	"github.com/aws/aws-sdk-go/service/docdb"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/service/ebs"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecrpublic"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go/service/elasticinference"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/aws/aws-sdk-go/service/emrcontainers"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/aws/aws-sdk-go/service/finspace"
	"github.com/aws/aws-sdk-go/service/finspacedata"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/fis"
	"github.com/aws/aws-sdk-go/service/fms"
	"github.com/aws/aws-sdk-go/service/forecastqueryservice"
	"github.com/aws/aws-sdk-go/service/forecastservice"
	"github.com/aws/aws-sdk-go/service/frauddetector"
	"github.com/aws/aws-sdk-go/service/fsx"
	"github.com/aws/aws-sdk-go/service/gamelift"
	"github.com/aws/aws-sdk-go/service/glacier"
	"github.com/aws/aws-sdk-go/service/globalaccelerator"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/aws/aws-sdk-go/service/gluedatabrew"
	"github.com/aws/aws-sdk-go/service/greengrass"
	"github.com/aws/aws-sdk-go/service/greengrassv2"
	"github.com/aws/aws-sdk-go/service/groundstation"
	"github.com/aws/aws-sdk-go/service/guardduty"
	"github.com/aws/aws-sdk-go/service/health"
	"github.com/aws/aws-sdk-go/service/healthlake"
	"github.com/aws/aws-sdk-go/service/honeycode"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/identitystore"
	"github.com/aws/aws-sdk-go/service/imagebuilder"
	"github.com/aws/aws-sdk-go/service/inspector"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iot1clickdevicesservice"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects"
	"github.com/aws/aws-sdk-go/service/iotanalytics"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"github.com/aws/aws-sdk-go/service/iotdeviceadvisor"
	"github.com/aws/aws-sdk-go/service/iotevents"
	"github.com/aws/aws-sdk-go/service/ioteventsdata"
	"github.com/aws/aws-sdk-go/service/iotfleethub"
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane"
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling"
	"github.com/aws/aws-sdk-go/service/iotsitewise"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph"
	"github.com/aws/aws-sdk-go/service/iotwireless"
	"github.com/aws/aws-sdk-go/service/ivs"
	"github.com/aws/aws-sdk-go/service/kafka"
	"github.com/aws/aws-sdk-go/service/kendra"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2"
	"github.com/aws/aws-sdk-go/service/kinesisvideo"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideosignalingchannels"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/lakeformation"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice"
	"github.com/aws/aws-sdk-go/service/lexmodelsv2"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice"
	"github.com/aws/aws-sdk-go/service/lexruntimev2"
	"github.com/aws/aws-sdk-go/service/licensemanager"
	"github.com/aws/aws-sdk-go/service/lightsail"
	"github.com/aws/aws-sdk-go/service/locationservice"
	"github.com/aws/aws-sdk-go/service/lookoutequipment"
	"github.com/aws/aws-sdk-go/service/lookoutforvision"
	"github.com/aws/aws-sdk-go/service/lookoutmetrics"
	"github.com/aws/aws-sdk-go/service/machinelearning"
	"github.com/aws/aws-sdk-go/service/macie"
	"github.com/aws/aws-sdk-go/service/macie2"
	"github.com/aws/aws-sdk-go/service/managedblockchain"
	"github.com/aws/aws-sdk-go/service/marketplacecatalog"
	"github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics"
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"github.com/aws/aws-sdk-go/service/marketplacemetering"
	"github.com/aws/aws-sdk-go/service/mediaconnect"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/aws/aws-sdk-go/service/medialive"
	"github.com/aws/aws-sdk-go/service/mediapackage"
	"github.com/aws/aws-sdk-go/service/mediapackagevod"
	"github.com/aws/aws-sdk-go/service/mediastore"
	"github.com/aws/aws-sdk-go/service/mediastoredata"
	"github.com/aws/aws-sdk-go/service/mediatailor"
	"github.com/aws/aws-sdk-go/service/mgn"
	"github.com/aws/aws-sdk-go/service/migrationhub"
	"github.com/aws/aws-sdk-go/service/migrationhubconfig"
	"github.com/aws/aws-sdk-go/service/mobile"
	"github.com/aws/aws-sdk-go/service/mobileanalytics"
	"github.com/aws/aws-sdk-go/service/mq"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/aws/aws-sdk-go/service/mwaa"
	"github.com/aws/aws-sdk-go/service/neptune"
	"github.com/aws/aws-sdk-go/service/networkfirewall"
	"github.com/aws/aws-sdk-go/service/networkmanager"
	"github.com/aws/aws-sdk-go/service/nimblestudio"
	"github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/aws/aws-sdk-go/service/opsworkscm"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/aws/aws-sdk-go/service/outposts"
	"github.com/aws/aws-sdk-go/service/personalize"
	"github.com/aws/aws-sdk-go/service/personalizeevents"
	"github.com/aws/aws-sdk-go/service/personalizeruntime"
	"github.com/aws/aws-sdk-go/service/pi"
	"github.com/aws/aws-sdk-go/service/pinpoint"
	"github.com/aws/aws-sdk-go/service/pinpointemail"
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/aws/aws-sdk-go/service/pricing"
	"github.com/aws/aws-sdk-go/service/prometheusservice"
	"github.com/aws/aws-sdk-go/service/qldb"
	"github.com/aws/aws-sdk-go/service/qldbsession"
	"github.com/aws/aws-sdk-go/service/quicksight"
	"github.com/aws/aws-sdk-go/service/ram"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/aws/aws-sdk-go/service/redshiftdataapiservice"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/resourcegroups"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/service/robomaker"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3control"
	"github.com/aws/aws-sdk-go/service/s3outposts"
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/aws/aws-sdk-go/service/sagemakeredgemanager"
	"github.com/aws/aws-sdk-go/service/sagemakerfeaturestoreruntime"
	"github.com/aws/aws-sdk-go/service/sagemakerruntime"
	"github.com/aws/aws-sdk-go/service/savingsplans"
	"github.com/aws/aws-sdk-go/service/schemas"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/securityhub"
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository"
	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"github.com/aws/aws-sdk-go/service/servicediscovery"
	"github.com/aws/aws-sdk-go/service/servicequotas"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sesv2"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/shield"
	"github.com/aws/aws-sdk-go/service/signer"
	"github.com/aws/aws-sdk-go/service/simpledb"
	"github.com/aws/aws-sdk-go/service/sms"
	"github.com/aws/aws-sdk-go/service/snowball"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssmcontacts"
	"github.com/aws/aws-sdk-go/service/ssmincidents"
	"github.com/aws/aws-sdk-go/service/sso"
	"github.com/aws/aws-sdk-go/service/ssoadmin"
	"github.com/aws/aws-sdk-go/service/ssooidc"
	"github.com/aws/aws-sdk-go/service/storagegateway"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
	"github.com/aws/aws-sdk-go/service/support"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/aws/aws-sdk-go/service/synthetics"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/aws/aws-sdk-go/service/timestreamquery"
	"github.com/aws/aws-sdk-go/service/timestreamwrite"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice"
	"github.com/aws/aws-sdk-go/service/transfer"
	"github.com/aws/aws-sdk-go/service/translate"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
	"github.com/aws/aws-sdk-go/service/wafv2"
	"github.com/aws/aws-sdk-go/service/wellarchitected"
	"github.com/aws/aws-sdk-go/service/workdocs"
	"github.com/aws/aws-sdk-go/service/worklink"
	"github.com/aws/aws-sdk-go/service/workmail"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow"
	"github.com/aws/aws-sdk-go/service/workspaces"
	"github.com/aws/aws-sdk-go/service/xray"
	"github.com/blinkops/blink-sdk/plugin"
	"github.com/blinkops/blink-sdk/plugin/actions"
	"github.com/blinkops/blink-sdk/plugin/config"
	"github.com/blinkops/blink-sdk/plugin/connections"
	description2 "github.com/blinkops/blink-sdk/plugin/description"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	packageActionSeparator = "_"
	serviceKey             = "_Service"
	awsRegionKey           = "awsRegion"
	runOnAllRegions        = "*"
	awsAccessKeyId         = "aws_access_key_id"
	awsSecretAccessKey     = "aws_secret_access_key"
	roleArn                = "role_arn"
	externalID             = "external_id"
)

type ActionExecutor func(map[string]interface{}) (map[string]interface{}, error)
type ActionHandlersMap map[string]func(map[string]interface{}) (map[string]interface{}, error)

type PackageParameters struct {
	packageName string
	awsSession  *session.Session // if nil returns ActionMap
}

type ActionParameters struct {
	Parameters map[string]interface{}
}

func NewActionParameters() *ActionParameters {
	return &ActionParameters{
		Parameters: make(map[string]interface{}),
	}
}

func NewActionParametersDeepCopy(parameters map[string]interface{}) *ActionParameters {
	actionParameters := make(map[string]interface{})
	for key, value := range parameters {
		actionParameters[key] = value
	}

	return &ActionParameters{
		Parameters: actionParameters,
	}
}

func (a *ActionParameters) Get(parameter string) interface{} {
	return a.Parameters[parameter]
}

func (a *ActionParameters) Set(parameter string, value interface{}) {
	a.Parameters[parameter] = value
}

// resolvePackage is a function maps package names to the real aws packages
// to renew the cases run the following command
// find './models/apis' -maxdepth 1 -execdir printf "\ncase \"{}\":\n\tif parameters.awsSession != nil {{ return {}.New(parameters.awsSession), nil }}\n\treturn {}.ActionMap, nil" \;
func resolvePackage(parameters *PackageParameters) (interface{}, error) {
	if parameters == nil {
		return nil, errors.New("failed to resolve package, no parameters provided")
	}

	switch strings.ToLower(parameters.packageName) {
	case "iotdataplane":
		if parameters.awsSession != nil {
			return iotdataplane.New(parameters.awsSession), nil
		}
		return iotdataplane.ActionMap, nil
	case "iotwireless":
		if parameters.awsSession != nil {
			return iotwireless.New(parameters.awsSession), nil
		}
		return iotwireless.ActionMap, nil
	case "emrcontainers":
		if parameters.awsSession != nil {
			return emrcontainers.New(parameters.awsSession), nil
		}
		return emrcontainers.ActionMap, nil
	case "timestreamwrite":
		if parameters.awsSession != nil {
			return timestreamwrite.New(parameters.awsSession), nil
		}
		return timestreamwrite.ActionMap, nil
	case "codebuild":
		if parameters.awsSession != nil {
			return codebuild.New(parameters.awsSession), nil
		}
		return codebuild.ActionMap, nil
	case "iotdeviceadvisor":
		if parameters.awsSession != nil {
			return iotdeviceadvisor.New(parameters.awsSession), nil
		}
		return iotdeviceadvisor.ActionMap, nil
	case "ssmcontacts":
		if parameters.awsSession != nil {
			return ssmcontacts.New(parameters.awsSession), nil
		}
		return ssmcontacts.ActionMap, nil
	case "wafregional":
		if parameters.awsSession != nil {
			return wafregional.New(parameters.awsSession), nil
		}
		return wafregional.ActionMap, nil
	case "lexmodelbuildingservice":
		if parameters.awsSession != nil {
			return lexmodelbuildingservice.New(parameters.awsSession), nil
		}
		return lexmodelbuildingservice.ActionMap, nil
	case "codeguruprofiler":
		if parameters.awsSession != nil {
			return codeguruprofiler.New(parameters.awsSession), nil
		}
		return codeguruprofiler.ActionMap, nil
	case "kinesis":
		if parameters.awsSession != nil {
			return kinesis.New(parameters.awsSession), nil
		}
		return kinesis.ActionMap, nil
	case "kinesisvideo":
		if parameters.awsSession != nil {
			return kinesisvideo.New(parameters.awsSession), nil
		}
		return kinesisvideo.ActionMap, nil
	case "pinpoint":
		if parameters.awsSession != nil {
			return pinpoint.New(parameters.awsSession), nil
		}
		return pinpoint.ActionMap, nil
	case "chime":
		if parameters.awsSession != nil {
			return chime.New(parameters.awsSession), nil
		}
		return chime.ActionMap, nil
	case "organizations":
		if parameters.awsSession != nil {
			return organizations.New(parameters.awsSession), nil
		}
		return organizations.ActionMap, nil
	case "licensemanager":
		if parameters.awsSession != nil {
			return licensemanager.New(parameters.awsSession), nil
		}
		return licensemanager.ActionMap, nil
	case "shield":
		if parameters.awsSession != nil {
			return shield.New(parameters.awsSession), nil
		}
		return shield.ActionMap, nil
	case "ssm":
		if parameters.awsSession != nil {
			return ssm.New(parameters.awsSession), nil
		}
		return ssm.ActionMap, nil
	case "mediastoredata":
		if parameters.awsSession != nil {
			return mediastoredata.New(parameters.awsSession), nil
		}
		return mediastoredata.ActionMap, nil
	case "sagemakerruntime":
		if parameters.awsSession != nil {
			return sagemakerruntime.New(parameters.awsSession), nil
		}
		return sagemakerruntime.ActionMap, nil
	case "signer":
		if parameters.awsSession != nil {
			return signer.New(parameters.awsSession), nil
		}
		return signer.ActionMap, nil
	case "servicecatalog":
		if parameters.awsSession != nil {
			return servicecatalog.New(parameters.awsSession), nil
		}
		return servicecatalog.ActionMap, nil
	case "lakeformation":
		if parameters.awsSession != nil {
			return lakeformation.New(parameters.awsSession), nil
		}
		return lakeformation.ActionMap, nil
	case "secretsmanager":
		if parameters.awsSession != nil {
			return secretsmanager.New(parameters.awsSession), nil
		}
		return secretsmanager.ActionMap, nil
	case "mediaconnect":
		if parameters.awsSession != nil {
			return mediaconnect.New(parameters.awsSession), nil
		}
		return mediaconnect.ActionMap, nil
	case "mwaa":
		if parameters.awsSession != nil {
			return mwaa.New(parameters.awsSession), nil
		}
		return mwaa.ActionMap, nil
	case "kms":
		if parameters.awsSession != nil {
			return kms.New(parameters.awsSession), nil
		}
		return kms.ActionMap, nil
	case "quicksight":
		if parameters.awsSession != nil {
			return quicksight.New(parameters.awsSession), nil
		}
		return quicksight.ActionMap, nil
	case "appregistry":
		if parameters.awsSession != nil {
			return appregistry.New(parameters.awsSession), nil
		}
		return appregistry.ActionMap, nil
	case "workmail":
		if parameters.awsSession != nil {
			return workmail.New(parameters.awsSession), nil
		}
		return workmail.ActionMap, nil
	case "eventbridge":
		if parameters.awsSession != nil {
			return eventbridge.New(parameters.awsSession), nil
		}
		return eventbridge.ActionMap, nil
	case "frauddetector":
		if parameters.awsSession != nil {
			return frauddetector.New(parameters.awsSession), nil
		}
		return frauddetector.ActionMap, nil
	case "elastictranscoder":
		if parameters.awsSession != nil {
			return elastictranscoder.New(parameters.awsSession), nil
		}
		return elastictranscoder.ActionMap, nil
	case "elasticinference":
		if parameters.awsSession != nil {
			return elasticinference.New(parameters.awsSession), nil
		}
		return elasticinference.ActionMap, nil
	case "lookoutequipment":
		if parameters.awsSession != nil {
			return lookoutequipment.New(parameters.awsSession), nil
		}
		return lookoutequipment.ActionMap, nil
	case "pinpointsmsvoice":
		if parameters.awsSession != nil {
			return pinpointsmsvoice.New(parameters.awsSession), nil
		}
		return pinpointsmsvoice.ActionMap, nil
	case "cloudwatch":
		if parameters.awsSession != nil {
			return cloudwatch.New(parameters.awsSession), nil
		}
		return cloudwatch.ActionMap, nil
	case "glue":
		if parameters.awsSession != nil {
			return glue.New(parameters.awsSession), nil
		}
		return glue.ActionMap, nil
	case "servicequotas":
		if parameters.awsSession != nil {
			return servicequotas.New(parameters.awsSession), nil
		}
		return servicequotas.ActionMap, nil
	case "marketplaceentitlementservice":
		if parameters.awsSession != nil {
			return marketplaceentitlementservice.New(parameters.awsSession), nil
		}
		return marketplaceentitlementservice.ActionMap, nil
	case "s3":
		if parameters.awsSession != nil {
			return s3.New(parameters.awsSession), nil
		}
		return s3.ActionMap, nil
	case "sesv2":
		if parameters.awsSession != nil {
			return sesv2.New(parameters.awsSession), nil
		}
		return sesv2.ActionMap, nil
	case "emr":
		if parameters.awsSession != nil {
			return emr.New(parameters.awsSession), nil
		}
		return emr.ActionMap, nil
	case "configservice":
		if parameters.awsSession != nil {
			return configservice.New(parameters.awsSession), nil
		}
		return configservice.ActionMap, nil
	case "iotfleethub":
		if parameters.awsSession != nil {
			return iotfleethub.New(parameters.awsSession), nil
		}
		return iotfleethub.ActionMap, nil
	case "personalize":
		if parameters.awsSession != nil {
			return personalize.New(parameters.awsSession), nil
		}
		return personalize.ActionMap, nil
	case "outposts":
		if parameters.awsSession != nil {
			return outposts.New(parameters.awsSession), nil
		}
		return outposts.ActionMap, nil
	case "workdocs":
		if parameters.awsSession != nil {
			return workdocs.New(parameters.awsSession), nil
		}
		return workdocs.ActionMap, nil
	case "networkmanager":
		if parameters.awsSession != nil {
			return networkmanager.New(parameters.awsSession), nil
		}
		return networkmanager.ActionMap, nil
	case "mediapackage":
		if parameters.awsSession != nil {
			return mediapackage.New(parameters.awsSession), nil
		}
		return mediapackage.ActionMap, nil
	case "medialive":
		if parameters.awsSession != nil {
			return medialive.New(parameters.awsSession), nil
		}
		return medialive.ActionMap, nil
	case "mediaconvert":
		if parameters.awsSession != nil {
			return mediaconvert.New(parameters.awsSession), nil
		}
		return mediaconvert.ActionMap, nil
	case "cognitosync":
		if parameters.awsSession != nil {
			return cognitosync.New(parameters.awsSession), nil
		}
		return cognitosync.ActionMap, nil
	case "sns":
		if parameters.awsSession != nil {
			return sns.New(parameters.awsSession), nil
		}
		return sns.ActionMap, nil
	case "datasync":
		if parameters.awsSession != nil {
			return datasync.New(parameters.awsSession), nil
		}
		return datasync.ActionMap, nil
	case "macie":
		if parameters.awsSession != nil {
			return macie.New(parameters.awsSession), nil
		}
		return macie.ActionMap, nil
	case "greengrassv2":
		if parameters.awsSession != nil {
			return greengrassv2.New(parameters.awsSession), nil
		}
		return greengrassv2.ActionMap, nil
	case "pinpointemail":
		if parameters.awsSession != nil {
			return pinpointemail.New(parameters.awsSession), nil
		}
		return pinpointemail.ActionMap, nil
	case "iotanalytics":
		if parameters.awsSession != nil {
			return iotanalytics.New(parameters.awsSession), nil
		}
		return iotanalytics.ActionMap, nil
	case "groundstation":
		if parameters.awsSession != nil {
			return groundstation.New(parameters.awsSession), nil
		}
		return groundstation.ActionMap, nil
	case "fis":
		if parameters.awsSession != nil {
			return fis.New(parameters.awsSession), nil
		}
		return fis.ActionMap, nil
	case "cloudhsm":
		if parameters.awsSession != nil {
			return cloudhsm.New(parameters.awsSession), nil
		}
		return cloudhsm.ActionMap, nil
	case "ecrpublic":
		if parameters.awsSession != nil {
			return ecrpublic.New(parameters.awsSession), nil
		}
		return ecrpublic.ActionMap, nil
	case "sms":
		if parameters.awsSession != nil {
			return sms.New(parameters.awsSession), nil
		}
		return sms.ActionMap, nil
	case "cognitoidentity":
		if parameters.awsSession != nil {
			return cognitoidentity.New(parameters.awsSession), nil
		}
		return cognitoidentity.ActionMap, nil
	case "inspector":
		if parameters.awsSession != nil {
			return inspector.New(parameters.awsSession), nil
		}
		return inspector.ActionMap, nil
	case "translate":
		if parameters.awsSession != nil {
			return translate.New(parameters.awsSession), nil
		}
		return translate.ActionMap, nil
	case "fms":
		if parameters.awsSession != nil {
			return fms.New(parameters.awsSession), nil
		}
		return fms.ActionMap, nil
	case "ssmincidents":
		if parameters.awsSession != nil {
			return ssmincidents.New(parameters.awsSession), nil
		}
		return ssmincidents.ActionMap, nil
	case "s3control":
		if parameters.awsSession != nil {
			return s3control.New(parameters.awsSession), nil
		}
		return s3control.ActionMap, nil
	case "kinesisanalyticsv2":
		if parameters.awsSession != nil {
			return kinesisanalyticsv2.New(parameters.awsSession), nil
		}
		return kinesisanalyticsv2.ActionMap, nil
	case "marketplacecommerceanalytics":
		if parameters.awsSession != nil {
			return marketplacecommerceanalytics.New(parameters.awsSession), nil
		}
		return marketplacecommerceanalytics.ActionMap, nil
	case "synthetics":
		if parameters.awsSession != nil {
			return synthetics.New(parameters.awsSession), nil
		}
		return synthetics.ActionMap, nil
	case "prometheusservice":
		if parameters.awsSession != nil {
			return prometheusservice.New(parameters.awsSession), nil
		}
		return prometheusservice.ActionMap, nil
	case "costexplorer":
		if parameters.awsSession != nil {
			return costexplorer.New(parameters.awsSession), nil
		}
		return costexplorer.ActionMap, nil
	case "iotsecuretunneling":
		if parameters.awsSession != nil {
			return iotsecuretunneling.New(parameters.awsSession), nil
		}
		return iotsecuretunneling.ActionMap, nil
	case "cloudfront":
		if parameters.awsSession != nil {
			return cloudfront.New(parameters.awsSession), nil
		}
		return cloudfront.ActionMap, nil
	case "wafv2":
		if parameters.awsSession != nil {
			return wafv2.New(parameters.awsSession), nil
		}
		return wafv2.ActionMap, nil
	case "transcribeservice":
		if parameters.awsSession != nil {
			return transcribeservice.New(parameters.awsSession), nil
		}
		return transcribeservice.ActionMap, nil
	case "ec2instanceconnect":
		if parameters.awsSession != nil {
			return ec2instanceconnect.New(parameters.awsSession), nil
		}
		return ec2instanceconnect.ActionMap, nil
	case "iotthingsgraph":
		if parameters.awsSession != nil {
			return iotthingsgraph.New(parameters.awsSession), nil
		}
		return iotthingsgraph.ActionMap, nil
	case "health":
		if parameters.awsSession != nil {
			return health.New(parameters.awsSession), nil
		}
		return health.ActionMap, nil
	case "workmailmessageflow":
		if parameters.awsSession != nil {
			return workmailmessageflow.New(parameters.awsSession), nil
		}
		return workmailmessageflow.ActionMap, nil
	case "comprehendmedical":
		if parameters.awsSession != nil {
			return comprehendmedical.New(parameters.awsSession), nil
		}
		return comprehendmedical.ActionMap, nil
	case "sagemakeredgemanager":
		if parameters.awsSession != nil {
			return sagemakeredgemanager.New(parameters.awsSession), nil
		}
		return sagemakeredgemanager.ActionMap, nil
	case "accessanalyzer":
		if parameters.awsSession != nil {
			return accessanalyzer.New(parameters.awsSession), nil
		}
		return accessanalyzer.ActionMap, nil
	case "glacier":
		if parameters.awsSession != nil {
			return glacier.New(parameters.awsSession), nil
		}
		return glacier.ActionMap, nil
	case "kinesisvideoarchivedmedia":
		if parameters.awsSession != nil {
			return kinesisvideoarchivedmedia.New(parameters.awsSession), nil
		}
		return kinesisvideoarchivedmedia.ActionMap, nil
	case "lightsail":
		if parameters.awsSession != nil {
			return lightsail.New(parameters.awsSession), nil
		}
		return lightsail.ActionMap, nil
	case "imagebuilder":
		if parameters.awsSession != nil {
			return imagebuilder.New(parameters.awsSession), nil
		}
		return imagebuilder.ActionMap, nil
	case "migrationhub":
		if parameters.awsSession != nil {
			return migrationhub.New(parameters.awsSession), nil
		}
		return migrationhub.ActionMap, nil
	case "elasticbeanstalk":
		if parameters.awsSession != nil {
			return elasticbeanstalk.New(parameters.awsSession), nil
		}
		return elasticbeanstalk.ActionMap, nil
	case "cloudsearchdomain":
		if parameters.awsSession != nil {
			return cloudsearchdomain.New(parameters.awsSession), nil
		}
		return cloudsearchdomain.ActionMap, nil
	case "neptune":
		if parameters.awsSession != nil {
			return neptune.New(parameters.awsSession), nil
		}
		return neptune.ActionMap, nil
	case "transfer":
		if parameters.awsSession != nil {
			return transfer.New(parameters.awsSession), nil
		}
		return transfer.ActionMap, nil
	case "braket":
		if parameters.awsSession != nil {
			return braket.New(parameters.awsSession), nil
		}
		return braket.ActionMap, nil
	case "resourcegroups":
		if parameters.awsSession != nil {
			return resourcegroups.New(parameters.awsSession), nil
		}
		return resourcegroups.ActionMap, nil
	case "qldb":
		if parameters.awsSession != nil {
			return qldb.New(parameters.awsSession), nil
		}
		return qldb.ActionMap, nil
	case "ecr":
		if parameters.awsSession != nil {
			return ecr.New(parameters.awsSession), nil
		}
		return ecr.ActionMap, nil
	case "dynamodb":
		if parameters.awsSession != nil {
			return dynamodb.New(parameters.awsSession), nil
		}
		return dynamodb.ActionMap, nil
	case "qldbsession":
		if parameters.awsSession != nil {
			return qldbsession.New(parameters.awsSession), nil
		}
		return qldbsession.ActionMap, nil
	case "route53domains":
		if parameters.awsSession != nil {
			return route53domains.New(parameters.awsSession), nil
		}
		return route53domains.ActionMap, nil
	case "macie2":
		if parameters.awsSession != nil {
			return macie2.New(parameters.awsSession), nil
		}
		return macie2.ActionMap, nil
	case "applicationautoscaling":
		if parameters.awsSession != nil {
			return applicationautoscaling.New(parameters.awsSession), nil
		}
		return applicationautoscaling.ActionMap, nil
	case "s3outposts":
		if parameters.awsSession != nil {
			return s3outposts.New(parameters.awsSession), nil
		}
		return s3outposts.ActionMap, nil
	case "honeycode":
		if parameters.awsSession != nil {
			return honeycode.New(parameters.awsSession), nil
		}
		return honeycode.ActionMap, nil
	case "storagegateway":
		if parameters.awsSession != nil {
			return storagegateway.New(parameters.awsSession), nil
		}
		return storagegateway.ActionMap, nil
	case "ioteventsdata":
		if parameters.awsSession != nil {
			return ioteventsdata.New(parameters.awsSession), nil
		}
		return ioteventsdata.ActionMap, nil
	case "lookoutforvision":
		if parameters.awsSession != nil {
			return lookoutforvision.New(parameters.awsSession), nil
		}
		return lookoutforvision.ActionMap, nil
	case "ecs":
		if parameters.awsSession != nil {
			return ecs.New(parameters.awsSession), nil
		}
		return ecs.ActionMap, nil
	case "connectcontactlens":
		if parameters.awsSession != nil {
			return connectcontactlens.New(parameters.awsSession), nil
		}
		return connectcontactlens.ActionMap, nil
	case "cloudsearch":
		if parameters.awsSession != nil {
			return cloudsearch.New(parameters.awsSession), nil
		}
		return cloudsearch.ActionMap, nil
	case "cloudwatchlogs":
		if parameters.awsSession != nil {
			return cloudwatchlogs.New(parameters.awsSession), nil
		}
		return cloudwatchlogs.ActionMap, nil
	case "gluedatabrew":
		if parameters.awsSession != nil {
			return gluedatabrew.New(parameters.awsSession), nil
		}
		return gluedatabrew.ActionMap, nil
	case "directoryservice":
		if parameters.awsSession != nil {
			return directoryservice.New(parameters.awsSession), nil
		}
		return directoryservice.ActionMap, nil
	case "rdsdataservice":
		if parameters.awsSession != nil {
			return rdsdataservice.New(parameters.awsSession), nil
		}
		return rdsdataservice.ActionMap, nil
	case "route53resolver":
		if parameters.awsSession != nil {
			return route53resolver.New(parameters.awsSession), nil
		}
		return route53resolver.ActionMap, nil
	case "workspaces":
		if parameters.awsSession != nil {
			return workspaces.New(parameters.awsSession), nil
		}
		return workspaces.ActionMap, nil
	case "machinelearning":
		if parameters.awsSession != nil {
			return machinelearning.New(parameters.awsSession), nil
		}
		return machinelearning.ActionMap, nil
	case "iot1clickdevicesservice":
		if parameters.awsSession != nil {
			return iot1clickdevicesservice.New(parameters.awsSession), nil
		}
		return iot1clickdevicesservice.ActionMap, nil
	case "fsx":
		if parameters.awsSession != nil {
			return fsx.New(parameters.awsSession), nil
		}
		return fsx.ActionMap, nil
	case "codepipeline":
		if parameters.awsSession != nil {
			return codepipeline.New(parameters.awsSession), nil
		}
		return codepipeline.ActionMap, nil
	case "elasticsearchservice":
		if parameters.awsSession != nil {
			return elasticsearchservice.New(parameters.awsSession), nil
		}
		return elasticsearchservice.ActionMap, nil
	case "elb":
		if parameters.awsSession != nil {
			return elb.New(parameters.awsSession), nil
		}
		return elb.ActionMap, nil
	case "codestarnotifications":
		if parameters.awsSession != nil {
			return codestarnotifications.New(parameters.awsSession), nil
		}
		return codestarnotifications.ActionMap, nil
	case "schemas":
		if parameters.awsSession != nil {
			return schemas.New(parameters.awsSession), nil
		}
		return schemas.ActionMap, nil
	case "sqs":
		if parameters.awsSession != nil {
			return sqs.New(parameters.awsSession), nil
		}
		return sqs.ActionMap, nil
	case "appmesh":
		if parameters.awsSession != nil {
			return appmesh.New(parameters.awsSession), nil
		}
		return appmesh.ActionMap, nil
	case "iot":
		if parameters.awsSession != nil {
			return iot.New(parameters.awsSession), nil
		}
		return iot.ActionMap, nil
	case "ebs":
		if parameters.awsSession != nil {
			return ebs.New(parameters.awsSession), nil
		}
		return ebs.ActionMap, nil
	case "amplify":
		if parameters.awsSession != nil {
			return amplify.New(parameters.awsSession), nil
		}
		return amplify.ActionMap, nil
	case "wellarchitected":
		if parameters.awsSession != nil {
			return wellarchitected.New(parameters.awsSession), nil
		}
		return wellarchitected.ActionMap, nil
	case "redshift":
		if parameters.awsSession != nil {
			return redshift.New(parameters.awsSession), nil
		}
		return redshift.ActionMap, nil
	case "locationservice":
		if parameters.awsSession != nil {
			return locationservice.New(parameters.awsSession), nil
		}
		return locationservice.ActionMap, nil
	case "appflow":
		if parameters.awsSession != nil {
			return appflow.New(parameters.awsSession), nil
		}
		return appflow.ActionMap, nil
	case "gamelift":
		if parameters.awsSession != nil {
			return gamelift.New(parameters.awsSession), nil
		}
		return gamelift.ActionMap, nil
	case "cloudtrail":
		if parameters.awsSession != nil {
			return cloudtrail.New(parameters.awsSession), nil
		}
		return cloudtrail.ActionMap, nil
	case "forecastqueryservice":
		if parameters.awsSession != nil {
			return forecastqueryservice.New(parameters.awsSession), nil
		}
		return forecastqueryservice.ActionMap, nil
	case "applicationinsights":
		if parameters.awsSession != nil {
			return applicationinsights.New(parameters.awsSession), nil
		}
		return applicationinsights.ActionMap, nil
	case "mediatailor":
		if parameters.awsSession != nil {
			return mediatailor.New(parameters.awsSession), nil
		}
		return mediatailor.ActionMap, nil
	case "forecastservice":
		if parameters.awsSession != nil {
			return forecastservice.New(parameters.awsSession), nil
		}
		return forecastservice.ActionMap, nil
	case "pi":
		if parameters.awsSession != nil {
			return pi.New(parameters.awsSession), nil
		}
		return pi.ActionMap, nil
	case "appconfig":
		if parameters.awsSession != nil {
			return appconfig.New(parameters.awsSession), nil
		}
		return appconfig.ActionMap, nil
	case "lexruntimeservice":
		if parameters.awsSession != nil {
			return lexruntimeservice.New(parameters.awsSession), nil
		}
		return lexruntimeservice.ActionMap, nil
	case "sagemakerfeaturestoreruntime":
		if parameters.awsSession != nil {
			return sagemakerfeaturestoreruntime.New(parameters.awsSession), nil
		}
		return sagemakerfeaturestoreruntime.ActionMap, nil
	case "computeoptimizer":
		if parameters.awsSession != nil {
			return computeoptimizer.New(parameters.awsSession), nil
		}
		return computeoptimizer.ActionMap, nil
	case "connectparticipant":
		if parameters.awsSession != nil {
			return connectparticipant.New(parameters.awsSession), nil
		}
		return connectparticipant.ActionMap, nil
	case "mgn":
		if parameters.awsSession != nil {
			return mgn.New(parameters.awsSession), nil
		}
		return mgn.ActionMap, nil
	case "applicationcostprofiler":
		if parameters.awsSession != nil {
			return applicationcostprofiler.New(parameters.awsSession), nil
		}
		return applicationcostprofiler.ActionMap, nil
	case "iam":
		if parameters.awsSession != nil {
			return iam.New(parameters.awsSession), nil
		}
		return iam.ActionMap, nil
	case "networkfirewall":
		if parameters.awsSession != nil {
			return networkfirewall.New(parameters.awsSession), nil
		}
		return networkfirewall.ActionMap, nil
	case "mediastore":
		if parameters.awsSession != nil {
			return mediastore.New(parameters.awsSession), nil
		}
		return mediastore.ActionMap, nil
	case "cloud9":
		if parameters.awsSession != nil {
			return cloud9.New(parameters.awsSession), nil
		}
		return cloud9.ActionMap, nil
	case "sso":
		if parameters.awsSession != nil {
			return sso.New(parameters.awsSession), nil
		}
		return sso.ActionMap, nil
	case "auditmanager":
		if parameters.awsSession != nil {
			return auditmanager.New(parameters.awsSession), nil
		}
		return auditmanager.ActionMap, nil
	case "snowball":
		if parameters.awsSession != nil {
			return snowball.New(parameters.awsSession), nil
		}
		return snowball.ActionMap, nil
	case "kinesisanalytics":
		if parameters.awsSession != nil {
			return kinesisanalytics.New(parameters.awsSession), nil
		}
		return kinesisanalytics.ActionMap, nil
	case "opsworks":
		if parameters.awsSession != nil {
			return opsworks.New(parameters.awsSession), nil
		}
		return opsworks.ActionMap, nil
	case "identitystore":
		if parameters.awsSession != nil {
			return identitystore.New(parameters.awsSession), nil
		}
		return identitystore.ActionMap, nil
	case "textract":
		if parameters.awsSession != nil {
			return textract.New(parameters.awsSession), nil
		}
		return textract.ActionMap, nil
	case "eks":
		if parameters.awsSession != nil {
			return eks.New(parameters.awsSession), nil
		}
		return eks.ActionMap, nil
	case "support":
		if parameters.awsSession != nil {
			return support.New(parameters.awsSession), nil
		}
		return support.ActionMap, nil
	case "mturk":
		if parameters.awsSession != nil {
			return mturk.New(parameters.awsSession), nil
		}
		return mturk.ActionMap, nil
	case "apigatewayv2":
		if parameters.awsSession != nil {
			return apigatewayv2.New(parameters.awsSession), nil
		}
		return apigatewayv2.ActionMap, nil
	case "devopsguru":
		if parameters.awsSession != nil {
			return devopsguru.New(parameters.awsSession), nil
		}
		return devopsguru.ActionMap, nil
	case "redshiftdataapiservice":
		if parameters.awsSession != nil {
			return redshiftdataapiservice.New(parameters.awsSession), nil
		}
		return redshiftdataapiservice.ActionMap, nil
	case "migrationhubconfig":
		if parameters.awsSession != nil {
			return migrationhubconfig.New(parameters.awsSession), nil
		}
		return migrationhubconfig.ActionMap, nil
	case "kafka":
		if parameters.awsSession != nil {
			return kafka.New(parameters.awsSession), nil
		}
		return kafka.ActionMap, nil
	case "mobile":
		if parameters.awsSession != nil {
			return mobile.New(parameters.awsSession), nil
		}
		return mobile.ActionMap, nil
	case "codedeploy":
		if parameters.awsSession != nil {
			return codedeploy.New(parameters.awsSession), nil
		}
		return codedeploy.ActionMap, nil
	case "cloudhsmv2":
		if parameters.awsSession != nil {
			return cloudhsmv2.New(parameters.awsSession), nil
		}
		return cloudhsmv2.ActionMap, nil
	case "batch":
		if parameters.awsSession != nil {
			return batch.New(parameters.awsSession), nil
		}
		return batch.ActionMap, nil
	case "simpledb":
		if parameters.awsSession != nil {
			return simpledb.New(parameters.awsSession), nil
		}
		return simpledb.ActionMap, nil
	case "iot1clickprojects":
		if parameters.awsSession != nil {
			return iot1clickprojects.New(parameters.awsSession), nil
		}
		return iot1clickprojects.ActionMap, nil
	case "kinesisvideosignalingchannels":
		if parameters.awsSession != nil {
			return kinesisvideosignalingchannels.New(parameters.awsSession), nil
		}
		return kinesisvideosignalingchannels.ActionMap, nil
	case "savingsplans":
		if parameters.awsSession != nil {
			return savingsplans.New(parameters.awsSession), nil
		}
		return savingsplans.ActionMap, nil
	case "appsync":
		if parameters.awsSession != nil {
			return appsync.New(parameters.awsSession), nil
		}
		return appsync.ActionMap, nil
	case "dlm":
		if parameters.awsSession != nil {
			return dlm.New(parameters.awsSession), nil
		}
		return dlm.ActionMap, nil
	case "amplifybackend":
		if parameters.awsSession != nil {
			return amplifybackend.New(parameters.awsSession), nil
		}
		return amplifybackend.ActionMap, nil
	case "budgets":
		if parameters.awsSession != nil {
			return budgets.New(parameters.awsSession), nil
		}
		return budgets.ActionMap, nil
	case "finspace":
		if parameters.awsSession != nil {
			return finspace.New(parameters.awsSession), nil
		}
		return finspace.ActionMap, nil
	case "detective":
		if parameters.awsSession != nil {
			return detective.New(parameters.awsSession), nil
		}
		return detective.ActionMap, nil
	case "lambda":
		if parameters.awsSession != nil {
			return lambda.New(parameters.awsSession), nil
		}
		return lambda.ActionMap, nil
	case "ssooidc":
		if parameters.awsSession != nil {
			return ssooidc.New(parameters.awsSession), nil
		}
		return ssooidc.ActionMap, nil
	case "applicationdiscoveryservice":
		if parameters.awsSession != nil {
			return applicationdiscoveryservice.New(parameters.awsSession), nil
		}
		return applicationdiscoveryservice.ActionMap, nil
	case "nimblestudio":
		if parameters.awsSession != nil {
			return nimblestudio.New(parameters.awsSession), nil
		}
		return nimblestudio.ActionMap, nil
	case "iotevents":
		if parameters.awsSession != nil {
			return iotevents.New(parameters.awsSession), nil
		}
		return iotevents.ActionMap, nil
	case "managedblockchain":
		if parameters.awsSession != nil {
			return managedblockchain.New(parameters.awsSession), nil
		}
		return managedblockchain.ActionMap, nil
	case "servicediscovery":
		if parameters.awsSession != nil {
			return servicediscovery.New(parameters.awsSession), nil
		}
		return servicediscovery.ActionMap, nil
	case "waf":
		if parameters.awsSession != nil {
			return waf.New(parameters.awsSession), nil
		}
		return waf.ActionMap, nil
	case "mobileanalytics":
		if parameters.awsSession != nil {
			return mobileanalytics.New(parameters.awsSession), nil
		}
		return mobileanalytics.ActionMap, nil
	case "ivs":
		if parameters.awsSession != nil {
			return ivs.New(parameters.awsSession), nil
		}
		return ivs.ActionMap, nil
	case "directconnect":
		if parameters.awsSession != nil {
			return directconnect.New(parameters.awsSession), nil
		}
		return directconnect.ActionMap, nil
	case "mq":
		if parameters.awsSession != nil {
			return mq.New(parameters.awsSession), nil
		}
		return mq.ActionMap, nil
	case "iotsitewise":
		if parameters.awsSession != nil {
			return iotsitewise.New(parameters.awsSession), nil
		}
		return iotsitewise.ActionMap, nil
	case "codestar":
		if parameters.awsSession != nil {
			return codestar.New(parameters.awsSession), nil
		}
		return codestar.ActionMap, nil
	case "lexmodelsv2":
		if parameters.awsSession != nil {
			return lexmodelsv2.New(parameters.awsSession), nil
		}
		return lexmodelsv2.ActionMap, nil
	case "lexruntimev2":
		if parameters.awsSession != nil {
			return lexruntimev2.New(parameters.awsSession), nil
		}
		return lexruntimev2.ActionMap, nil
	case "serverlessapplicationrepository":
		if parameters.awsSession != nil {
			return serverlessapplicationrepository.New(parameters.awsSession), nil
		}
		return serverlessapplicationrepository.ActionMap, nil
	case "clouddirectory":
		if parameters.awsSession != nil {
			return clouddirectory.New(parameters.awsSession), nil
		}
		return clouddirectory.ActionMap, nil
	case "mediapackagevod":
		if parameters.awsSession != nil {
			return mediapackagevod.New(parameters.awsSession), nil
		}
		return mediapackagevod.ActionMap, nil
	case "databasemigrationservice":
		if parameters.awsSession != nil {
			return databasemigrationservice.New(parameters.awsSession), nil
		}
		return databasemigrationservice.ActionMap, nil
	case "codestarconnections":
		if parameters.awsSession != nil {
			return codestarconnections.New(parameters.awsSession), nil
		}
		return codestarconnections.ActionMap, nil
	case "codeartifact":
		if parameters.awsSession != nil {
			return codeartifact.New(parameters.awsSession), nil
		}
		return codeartifact.ActionMap, nil
	case "guardduty":
		if parameters.awsSession != nil {
			return guardduty.New(parameters.awsSession), nil
		}
		return guardduty.ActionMap, nil
	case "worklink":
		if parameters.awsSession != nil {
			return worklink.New(parameters.awsSession), nil
		}
		return worklink.ActionMap, nil
	case "customerprofiles":
		if parameters.awsSession != nil {
			return customerprofiles.New(parameters.awsSession), nil
		}
		return customerprofiles.ActionMap, nil
	case "dax":
		if parameters.awsSession != nil {
			return dax.New(parameters.awsSession), nil
		}
		return dax.ActionMap, nil
	case "opsworkscm":
		if parameters.awsSession != nil {
			return opsworkscm.New(parameters.awsSession), nil
		}
		return opsworkscm.ActionMap, nil
	case "docdb":
		if parameters.awsSession != nil {
			return docdb.New(parameters.awsSession), nil
		}
		return docdb.ActionMap, nil
	case "acmpca":
		if parameters.awsSession != nil {
			return acmpca.New(parameters.awsSession), nil
		}
		return acmpca.ActionMap, nil
	case "firehose":
		if parameters.awsSession != nil {
			return firehose.New(parameters.awsSession), nil
		}
		return firehose.ActionMap, nil
	case "dynamodbstreams":
		if parameters.awsSession != nil {
			return dynamodbstreams.New(parameters.awsSession), nil
		}
		return dynamodbstreams.ActionMap, nil
	case "globalaccelerator":
		if parameters.awsSession != nil {
			return globalaccelerator.New(parameters.awsSession), nil
		}
		return globalaccelerator.ActionMap, nil
	case "ses":
		if parameters.awsSession != nil {
			return ses.New(parameters.awsSession), nil
		}
		return ses.ActionMap, nil
	case "codegurureviewer":
		if parameters.awsSession != nil {
			return codegurureviewer.New(parameters.awsSession), nil
		}
		return codegurureviewer.ActionMap, nil
	case "alexaforbusiness":
		if parameters.awsSession != nil {
			return alexaforbusiness.New(parameters.awsSession), nil
		}
		return alexaforbusiness.ActionMap, nil
	case "robomaker":
		if parameters.awsSession != nil {
			return robomaker.New(parameters.awsSession), nil
		}
		return robomaker.ActionMap, nil
	case "autoscaling":
		if parameters.awsSession != nil {
			return autoscaling.New(parameters.awsSession), nil
		}
		return autoscaling.ActionMap, nil
	case "iotjobsdataplane":
		if parameters.awsSession != nil {
			return iotjobsdataplane.New(parameters.awsSession), nil
		}
		return iotjobsdataplane.ActionMap, nil
	case "elbv2":
		if parameters.awsSession != nil {
			return elbv2.New(parameters.awsSession), nil
		}
		return elbv2.ActionMap, nil
	case "augmentedairuntime":
		if parameters.awsSession != nil {
			return augmentedairuntime.New(parameters.awsSession), nil
		}
		return augmentedairuntime.ActionMap, nil
	case "greengrass":
		if parameters.awsSession != nil {
			return greengrass.New(parameters.awsSession), nil
		}
		return greengrass.ActionMap, nil
	case "securityhub":
		if parameters.awsSession != nil {
			return securityhub.New(parameters.awsSession), nil
		}
		return securityhub.ActionMap, nil
	case "timestreamquery":
		if parameters.awsSession != nil {
			return timestreamquery.New(parameters.awsSession), nil
		}
		return timestreamquery.ActionMap, nil
	case "backup":
		if parameters.awsSession != nil {
			return backup.New(parameters.awsSession), nil
		}
		return backup.ActionMap, nil
	case "cloudformation":
		if parameters.awsSession != nil {
			return cloudformation.New(parameters.awsSession), nil
		}
		return cloudformation.ActionMap, nil
	case "kendra":
		if parameters.awsSession != nil {
			return kendra.New(parameters.awsSession), nil
		}
		return kendra.ActionMap, nil
	case "connect":
		if parameters.awsSession != nil {
			return connect.New(parameters.awsSession), nil
		}
		return connect.ActionMap, nil
	case "elasticache":
		if parameters.awsSession != nil {
			return elasticache.New(parameters.awsSession), nil
		}
		return elasticache.ActionMap, nil
	case "sfn":
		if parameters.awsSession != nil {
			return sfn.New(parameters.awsSession), nil
		}
		return sfn.ActionMap, nil
	case "cognitoidentityprovider":
		if parameters.awsSession != nil {
			return cognitoidentityprovider.New(parameters.awsSession), nil
		}
		return cognitoidentityprovider.ActionMap, nil
	case "costandusagereportservice":
		if parameters.awsSession != nil {
			return costandusagereportservice.New(parameters.awsSession), nil
		}
		return costandusagereportservice.ActionMap, nil
	case "comprehend":
		if parameters.awsSession != nil {
			return comprehend.New(parameters.awsSession), nil
		}
		return comprehend.ActionMap, nil
	case "marketplacemetering":
		if parameters.awsSession != nil {
			return marketplacemetering.New(parameters.awsSession), nil
		}
		return marketplacemetering.ActionMap, nil
	case "devicefarm":
		if parameters.awsSession != nil {
			return devicefarm.New(parameters.awsSession), nil
		}
		return devicefarm.ActionMap, nil
	case "rekognition":
		if parameters.awsSession != nil {
			return rekognition.New(parameters.awsSession), nil
		}
		return rekognition.ActionMap, nil
	case "appstream":
		if parameters.awsSession != nil {
			return appstream.New(parameters.awsSession), nil
		}
		return appstream.ActionMap, nil
	case "polly":
		if parameters.awsSession != nil {
			return polly.New(parameters.awsSession), nil
		}
		return polly.ActionMap, nil
	case "appintegrationsservice":
		if parameters.awsSession != nil {
			return appintegrationsservice.New(parameters.awsSession), nil
		}
		return appintegrationsservice.ActionMap, nil
	case "rds":
		if parameters.awsSession != nil {
			return rds.New(parameters.awsSession), nil
		}
		return rds.ActionMap, nil
	case "pricing":
		if parameters.awsSession != nil {
			return pricing.New(parameters.awsSession), nil
		}
		return pricing.ActionMap, nil
	case "swf":
		if parameters.awsSession != nil {
			return swf.New(parameters.awsSession), nil
		}
		return swf.ActionMap, nil
	case "cloudwatchevents":
		if parameters.awsSession != nil {
			return cloudwatchevents.New(parameters.awsSession), nil
		}
		return cloudwatchevents.ActionMap, nil
	case "transcribestreamingservice":
		if parameters.awsSession != nil {
			return transcribestreamingservice.New(parameters.awsSession), nil
		}
		return transcribestreamingservice.ActionMap, nil
	case "autoscalingplans":
		if parameters.awsSession != nil {
			return autoscalingplans.New(parameters.awsSession), nil
		}
		return autoscalingplans.ActionMap, nil
	case "datapipeline":
		if parameters.awsSession != nil {
			return datapipeline.New(parameters.awsSession), nil
		}
		return datapipeline.ActionMap, nil
	case "personalizeruntime":
		if parameters.awsSession != nil {
			return personalizeruntime.New(parameters.awsSession), nil
		}
		return personalizeruntime.ActionMap, nil
	case "codecommit":
		if parameters.awsSession != nil {
			return codecommit.New(parameters.awsSession), nil
		}
		return codecommit.ActionMap, nil
	case "resourcegroupstaggingapi":
		if parameters.awsSession != nil {
			return resourcegroupstaggingapi.New(parameters.awsSession), nil
		}
		return resourcegroupstaggingapi.ActionMap, nil
	case "healthlake":
		if parameters.awsSession != nil {
			return healthlake.New(parameters.awsSession), nil
		}
		return healthlake.ActionMap, nil
	case "personalizeevents":
		if parameters.awsSession != nil {
			return personalizeevents.New(parameters.awsSession), nil
		}
		return personalizeevents.ActionMap, nil
	case "apigatewaymanagementapi":
		if parameters.awsSession != nil {
			return apigatewaymanagementapi.New(parameters.awsSession), nil
		}
		return apigatewaymanagementapi.ActionMap, nil
	case "xray":
		if parameters.awsSession != nil {
			return xray.New(parameters.awsSession), nil
		}
		return xray.ActionMap, nil
	case "ssoadmin":
		if parameters.awsSession != nil {
			return ssoadmin.New(parameters.awsSession), nil
		}
		return ssoadmin.ActionMap, nil
	case "apigateway":
		if parameters.awsSession != nil {
			return apigateway.New(parameters.awsSession), nil
		}
		return apigateway.ActionMap, nil
	case "ram":
		if parameters.awsSession != nil {
			return ram.New(parameters.awsSession), nil
		}
		return ram.ActionMap, nil
	case "efs":
		if parameters.awsSession != nil {
			return efs.New(parameters.awsSession), nil
		}
		return efs.ActionMap, nil
	case "kinesisvideomedia":
		if parameters.awsSession != nil {
			return kinesisvideomedia.New(parameters.awsSession), nil
		}
		return kinesisvideomedia.ActionMap, nil
	case "dataexchange":
		if parameters.awsSession != nil {
			return dataexchange.New(parameters.awsSession), nil
		}
		return dataexchange.ActionMap, nil
	case "sts":
		if parameters.awsSession != nil {
			return sts.New(parameters.awsSession), nil
		}
		return sts.ActionMap, nil
	case "sagemaker":
		if parameters.awsSession != nil {
			return sagemaker.New(parameters.awsSession), nil
		}
		return sagemaker.ActionMap, nil
	case "finspacedata":
		if parameters.awsSession != nil {
			return finspacedata.New(parameters.awsSession), nil
		}
		return finspacedata.ActionMap, nil
	case "marketplacecatalog":
		if parameters.awsSession != nil {
			return marketplacecatalog.New(parameters.awsSession), nil
		}
		return marketplacecatalog.ActionMap, nil
	case "acm":
		if parameters.awsSession != nil {
			return acm.New(parameters.awsSession), nil
		}
		return acm.ActionMap, nil
	case "athena":
		if parameters.awsSession != nil {
			return athena.New(parameters.awsSession), nil
		}
		return athena.ActionMap, nil
	case "route53":
		if parameters.awsSession != nil {
			return route53.New(parameters.awsSession), nil
		}
		return route53.ActionMap, nil
	case "ec2":
		if parameters.awsSession != nil {
			return ec2.New(parameters.awsSession), nil
		}
		return ec2.ActionMap, nil
	case "apprunner":
		if parameters.awsSession != nil {
			return apprunner.New(parameters.awsSession), nil
		}
		return apprunner.ActionMap, nil
	case "lookoutmetrics":
		if parameters.awsSession != nil {
			return lookoutmetrics.New(parameters.awsSession), nil
		}
		return lookoutmetrics.ActionMap, nil
	}

	return nil, errors.New("provided package name is not supported: " + parameters.packageName)
}

func resolvePackageServiceByContext(packageName string, region string, ctx *plugin.ActionContext, timeout int32) (interface{}, error) {
	awsSession, err := createAWSSessionByContext(region, ctx, timeout)
	if err != nil {
		return nil, err
	}

	parameters := &PackageParameters{packageName: packageName, awsSession: awsSession}

	return resolvePackage(parameters)
}

func resolvePackageServiceByCredentials(packageName string, region string, awsCredentials map[string]interface{}) (interface{}, error) {
	awsSession, err := createAWSSessionByCredentials(region, awsCredentials, 0)
	if err != nil {
		return nil, err
	}

	parameters := &PackageParameters{packageName: packageName, awsSession: awsSession}

	return resolvePackage(parameters)
}

func resolvePackageActionMap(packageName string) (ActionHandlersMap, error) {
	parameters := &PackageParameters{packageName: packageName, awsSession: nil}

	resolvedActionMap, err := resolvePackage(parameters)
	if err != nil {
		return nil, err
	}

	actionMap, ok := resolvedActionMap.(map[string]func(map[string]interface{}) (map[string]interface{}, error))
	if !ok {
		return nil, fmt.Errorf("failed to access %v actions", packageName)
	}

	return actionMap, nil
}

func resolveActionExecutor(packageName string, actionFunctionName string) (ActionExecutor, error) {
	actionMap, err := resolvePackageActionMap(packageName)
	if err != nil {
		return nil, err
	}

	actionExecutor, ok := actionMap[actionFunctionName]
	if !ok {
		log.Error("Unknown action: ", actionFunctionName)
		return nil, errors.New("unknown action " + actionFunctionName)
	}

	return actionExecutor, nil
}

func resolvePackageNameAndActionFunctionName(requestActionName string) (string, string, error) {
	actionNameComponents := strings.Split(requestActionName, packageActionSeparator)
	if len(actionNameComponents) != 2 {
		return "", "", errors.New("invalid action name structure, failed to infer component name")
	}

	return actionNameComponents[0], actionNameComponents[1], nil
}

type AWSPlugin struct {
	actions     []plugin.Action
	description plugin.Description
}

func (p *AWSPlugin) Describe() plugin.Description {
	log.Debug("Handling Describe request!")
	return p.description
}

func (p *AWSPlugin) GetActions() []plugin.Action {
	log.Debug("Handling GetActions request!")
	return p.actions
}

func executeActionOnRegions(packageName string, ctx *plugin.ActionContext, actionExecutor ActionExecutor, actionParameters ActionParameters, regions []string) map[string]interface{} {
	syncResults := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(len(regions))

	for _, region := range regions {
		go func(region string, results *sync.Map, group *sync.WaitGroup, parameters ActionParameters) {
			copiedParameters := NewActionParametersDeepCopy(parameters.Parameters)
			copiedParameters.Set(awsRegionKey, strings.TrimSpace(region))
			regionResult, err := executeAction(packageName, ctx, actionExecutor, *copiedParameters, 0)

			if err != nil {
				results.Store(region, map[string]interface{}{
					"error": err.Error(),
				})
			} else {
				results.Store(region, regionResult)
			}
			group.Done()
		}(region, &syncResults, &wg, actionParameters)
	}

	wg.Wait()

	result := make(map[string]interface{})
	syncResults.Range(func(key, value interface{}) bool {
		resultKey := ""
		switch key.(type) {
		case string:
			resultKey = key.(string)
		default:
			resultKey = fmt.Sprintf("%v", key)
		}
		result[resultKey] = value

		return true
	})

	return result
}

func executeAction(packageName string, ctx *plugin.ActionContext, actionExecutor ActionExecutor, actionParameters ActionParameters, timeout int32) (map[string]interface{}, error) {
	actionRegion := actionParameters.Get(awsRegionKey).(string)
	if actionRegion == runOnAllRegions {
		availableRegions := awsutil.GetServiceRegions(packageName)
		return executeActionOnRegions(packageName, ctx, actionExecutor, actionParameters, availableRegions), nil
	}

	if strings.Contains(actionRegion, ",") {
		availableRegions := strings.Split(actionRegion, ",")
		return executeActionOnRegions(packageName, ctx, actionExecutor, actionParameters, availableRegions), nil
	}

	if err := appendServiceToParametersByContext(packageName, ctx, actionParameters.Parameters, timeout); err != nil {
		return nil, fmt.Errorf("failed to append service to parameters, error: %v", err)
	}

	log.Tracef("Received parameters for action: %v", actionParameters)
	output, err := actionExecutor(actionParameters.Parameters)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (p *AWSPlugin) ExecuteAction(ctx *plugin.ActionContext, request *plugin.ExecuteActionRequest) (*plugin.ExecuteActionResponse, error) {
	log.Debugf("Requested to execute action: \n %v", *request)

	packageName, actionFunctionName, err := resolvePackageNameAndActionFunctionName(request.Name)
	if err != nil {
		return nil, err
	}

	awsActionParameters, err := getActionParameters(request)
	if err != nil {
		return nil, fmt.Errorf("failed to get action parameters, error: %v", err)
	}

	actionExecutor, err := resolveActionExecutor(packageName, actionFunctionName)
	if err != nil {
		return nil, fmt.Errorf("failed to get action executor, error: %v", err)
	}

	output, err := executeAction(packageName, ctx, actionExecutor, *awsActionParameters, request.Timeout)
	if err != nil {
		return nil, err
	}

	marshaledResponse, err := json.Marshal(output)
	if err != nil {
		log.Error("Failed to marshal output with error: ", err)
		return nil, err
	}

	log.Debug("Finished executing action: ", request.Name)

	return &plugin.ExecuteActionResponse{
		ErrorCode: 0,
		Result:    marshaledResponse,
	}, nil
}

func (p *AWSPlugin) TestCredentials(credentialsMap map[string]connections.ConnectionInstance) (*plugin.CredentialsValidationResponse, error) {
	log.Debugf("Requested to test credentials: \n %v", credentialsMap)

	for _, connInstance := range credentialsMap {
		awsCredentials, err := connInstance.ResolveCredentials()
		if err != nil {
			return &plugin.CredentialsValidationResponse{
				AreCredentialsValid:   false,
				RawValidationResponse: []byte(err.Error()),
			}, err
		}

		serviceName := "sts"
		serviceRegions := awsutil.GetServiceRegions(serviceName)

		if len(serviceRegions) == 0 {
			return &plugin.CredentialsValidationResponse{
				AreCredentialsValid:   false,
				RawValidationResponse: []byte("failed to get service regions to test connection with"),
			}, fmt.Errorf("failed to get service regions to test connection with")
		}

		awsActionParameters := map[string]interface{}{
			awsRegionKey: serviceRegions[0],
		}

		if err := appendServiceToParametersByCredentials(serviceName, awsCredentials, awsActionParameters); err != nil {
			return &plugin.CredentialsValidationResponse{
				AreCredentialsValid:   false,
				RawValidationResponse: []byte(err.Error()),
			}, fmt.Errorf("failed to append service to parameters, error: %v", err)
		}

		output, err := sts.ExecuteGetCallerIdentity(awsActionParameters)
		if err != nil {
			log.Debugf("failed on credentials validation, got: %v", output)
			return &plugin.CredentialsValidationResponse{
				AreCredentialsValid:   false,
				RawValidationResponse: []byte("failed on credentials validation"),
			}, fmt.Errorf("failed on credentials validation, got: %v and error: %w", output, err)
		}

		log.Debugf("Credentials are valid, continue to the next instnace")
	}

	log.Debugf("All credentials are valid...")
	return &plugin.CredentialsValidationResponse{
		AreCredentialsValid:   true,
		RawValidationResponse: []byte("Credentials are valid!"),
	}, nil
}

func NewAWSPlugin(rootPluginDirectory string) (*AWSPlugin, error) {

	pluginConfig := config.GetConfig()

	actionFolderPath := path.Join(rootPluginDirectory, pluginConfig.Plugin.ActionsFolderPath)

	logrus.Info("Loading actions from folder: ", actionFolderPath)
	actionsFromDisk, err := actions.LoadActionsFromDisk(actionFolderPath)
	if err != nil {
		return nil, err
	}

	description, err := description2.LoadPluginDescriptionFromDisk(path.Join(rootPluginDirectory, pluginConfig.Plugin.PluginDescriptionFilePath))
	if err != nil {
		return nil, err
	}

	loadedConnections, err := connections.LoadConnectionsFromDisk(path.Join(rootPluginDirectory, pluginConfig.Plugin.PluginDescriptionFilePath))
	if err != nil {
		return nil, err
	}

	log.Infof("Loaded %d connections from disk", len(loadedConnections))
	description.Connections = loadedConnections

	return &AWSPlugin{
		actions:     actionsFromDisk,
		description: *description,
	}, nil
}

func getActionParameters(request *plugin.ExecuteActionRequest) (*ActionParameters, error) {
	awsActionParameters := NewActionParameters()

	actionParameters, err := request.GetParameters()
	if err == nil {
		for key, value := range actionParameters {
			awsActionParameters.Set(key, value)
		}
	} else if err.Error() == plugin.ErrParametersAsJsonProvided {
		unmarshalledParameters, err := request.GetUnmarshalledParameters()
		if err != nil {
			return nil, fmt.Errorf("failed to get unmarshalled parameters, error %v", err)
		}

		for key, value := range unmarshalledParameters {
			awsActionParameters.Set(key, value)
		}
	}

	return awsActionParameters, nil
}

func createAWSSessionByContext(region string, context *plugin.ActionContext, timeout int32) (*session.Session, error) {
	awsCredentials, err := context.GetCredentials("aws")
	if err != nil {
		log.Errorf("Failed to get AWS awsCredentials: %v", err)
		return nil, err
	}
	return createAWSSessionByCredentials(region, awsCredentials, timeout)
}

// access keys have to be both set
// role arn can be supplied alone if it's irsa
// role arn and external id have to be supplied together for traditional assume role
func detectConnectionType(awsCredentials map[string]string) (credsType, key, value string) {
	if awsCredentials[awsAccessKeyId] == "" || awsCredentials[awsSecretAccessKey] == "" {
		if awsCredentials[roleArn] == "" {
			return "", "", ""
		}
		return "roleBased", awsCredentials[roleArn], awsCredentials[externalID]
	}
	return "userBased", awsCredentials[awsAccessKeyId], awsCredentials[awsSecretAccessKey]
}

func convertInterfaceMapToStringMap(m map[string]interface{}) map[string]string {
	mapString := make(map[string]string)
	for key, value := range m {
		var strValue string
		strKey := fmt.Sprintf("%v", key)
		if value == nil {
			strValue = ""
		} else {
			strValue = fmt.Sprintf("%v", value)
		}
		mapString[strKey] = strValue
	}
	return mapString
}

func readFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func assumeRoleWithWebIdentity(svc stsiface.STSAPI, role, sessionName string) (string, string, string, error) {
	log.Debug("assuming role with web identity")
	tokenFile, ok := os.LookupEnv("AWS_WEB_IDENTITY_TOKEN_FILE")
	if !ok {
		return "", "", "", fmt.Errorf("token file for irsa not found. make sure your pod is configured correctly and that your service account is created and annotated properly")
	}

	data, err := readFile(tokenFile)
	if err != nil {
		return "", "", "", fmt.Errorf("unable to open web identity token file with error: %w", err)
	}

	result, err := svc.AssumeRoleWithWebIdentity(&sts.AssumeRoleWithWebIdentityInput{
		DurationSeconds:  aws.Int64(3600),
		RoleArn:          aws.String(role),
		RoleSessionName:  aws.String(sessionName),
		WebIdentityToken: aws.String(string(data)),
	})
	if err != nil {
		return "", "", "", err
	}
	return *result.Credentials.AccessKeyId, *result.Credentials.SecretAccessKey, *result.Credentials.SessionToken, err
}

func assumeRoleWithTrustedIdentity(svc stsiface.STSAPI, role, externalID, sessionName string) (string, string, string, error) {
	log.Debug("assuming role with trusted entity")
	result, err := svc.AssumeRole(&sts.AssumeRoleInput{
		RoleArn:         &role,
		RoleSessionName: &sessionName,
		ExternalId:      &externalID,
	})
	if err != nil {
		return "", "", "", err
	}
	return *result.Credentials.AccessKeyId, *result.Credentials.SecretAccessKey, *result.Credentials.SessionToken, err
}

func assumeRole(svc stsiface.STSAPI, role, externalID string) (access, secret, sessionToken string, err error) {
	sessionName := strconv.Itoa(rand.Int())

	// irsa does not work with externalID, only the "traditional" assume role does
	if externalID == "" {
		return assumeRoleWithWebIdentity(svc, role, sessionName)
	}
	return assumeRoleWithTrustedIdentity(svc, role, externalID, sessionName)
}

func createAWSSessionByCredentials(region string, awsCredentials map[string]interface{}, timeout int32) (*session.Session, error) {
	var access, secret, sessionToken string

	// aws credentials are always strings. interface is annoying, let's convert to map[string]string
	m := convertInterfaceMapToStringMap(awsCredentials)
	sessionType, k, v := detectConnectionType(m)

	// checking whether the session is going to be with regular aws access keys or a role needed to be assumed
	switch sessionType {
	case "roleBased":
		sess, _ := session.NewSession(&aws.Config{
			Region: aws.String(region),
		})

		svc := sts.New(sess)
		var err error
		access, secret, sessionToken, err = assumeRole(svc, k, v)
		if err != nil {
			return nil, fmt.Errorf("unable to assume role with error: %w", err)
		}
	case "userBased":
		access, secret, sessionToken = k, v, ""
	default:
		return nil, fmt.Errorf("invalid credentials: make sure access+secret key are supplied OR role_arn and/or external_id")
	}

	// setting the credentials we just obtained
	creds := credentials.NewStaticCredentials(access, secret, sessionToken)

	// Create new session
	awsConfig := &aws.Config{
		Region:      aws.String(region),
		Credentials: creds,
	}
	if timeout > 0 {
		awsConfig.HTTPClient = &http.Client{Timeout: time.Duration(timeout) * time.Second}
	}
	sess, err := session.NewSession(awsConfig)

	if err != nil {
		return nil, errors.New("failed to create AWS session using provided credentials")
	}
	return sess, nil
}

func appendServiceToParametersByContext(packageName string, context *plugin.ActionContext, awsActionParameters map[string]interface{}, timeout int32) error {
	region, ok := awsActionParameters[awsRegionKey].(string)
	if !ok {
		return fmt.Errorf("failed to get region from action parameters")
	}

	service, err := resolvePackageServiceByContext(packageName, region, context, timeout)
	if err != nil {
		return err
	}

	awsActionParameters[serviceKey] = service
	return nil
}

func appendServiceToParametersByCredentials(packageName string, awsCredentials map[string]interface{}, awsActionParameters map[string]interface{}) error {
	region, ok := awsActionParameters[awsRegionKey].(string)
	if !ok {
		return fmt.Errorf("failed to get region from action parameters")
	}

	service, err := resolvePackageServiceByCredentials(packageName, region, awsCredentials)
	if err != nil {
		return err
	}

	awsActionParameters[serviceKey] = service
	return nil
}
