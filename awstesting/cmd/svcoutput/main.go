package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice"
	"github.com/aws/aws-sdk-go/service/appstream"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/batch"
	"github.com/aws/aws-sdk-go/service/budgets"
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
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/aws/aws-sdk-go/service/codecommit"
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/aws/aws-sdk-go/service/codestar"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitosync"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/datapipeline"
	"github.com/aws/aws-sdk-go/service/dax"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/service/directconnect"
	"github.com/aws/aws-sdk-go/service/directoryservice"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/gamelift"
	"github.com/aws/aws-sdk-go/service/glacier"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/aws/aws-sdk-go/service/greengrass"
	"github.com/aws/aws-sdk-go/service/health"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/inspector"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice"
	"github.com/aws/aws-sdk-go/service/lightsail"
	"github.com/aws/aws-sdk-go/service/machinelearning"
	"github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics"
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"github.com/aws/aws-sdk-go/service/marketplacemetering"
	"github.com/aws/aws-sdk-go/service/migrationhub"
	"github.com/aws/aws-sdk-go/service/mobile"
	"github.com/aws/aws-sdk-go/service/mobileanalytics"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/aws/aws-sdk-go/service/opsworkscm"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/aws/aws-sdk-go/service/pinpoint"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/shield"
	"github.com/aws/aws-sdk-go/service/simpledb"
	"github.com/aws/aws-sdk-go/service/sms"
	"github.com/aws/aws-sdk-go/service/snowball"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/storagegateway"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/support"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
	"github.com/aws/aws-sdk-go/service/workdocs"
	"github.com/aws/aws-sdk-go/service/workspaces"
	"github.com/aws/aws-sdk-go/service/xray"
)

func main() {
	w := writer{
		buf:       bytes.NewBuffer(nil),
		indentStr: "\t",
	}
	server := setupServer(w.Indent())
	defer server.Close()

	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:               &server.URL,
		Region:                 aws.String(endpoints.UsWest2RegionID),
		Credentials:            credentials.AnonymousCredentials,
		S3ForcePathStyle:       aws.Bool(true),
		DisableParamValidation: aws.Bool(true),
		//		LogLevel:               aws.LogLevel(aws.LogDebugWithHTTPBody),
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
					DualStack: true,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		},
	}))
	sess.Handlers.Send.PushFront(func(r *request.Request) {
		w.Writef("%s.%s:", r.ClientInfo.ServiceName, r.Operation.Name)
	})

	for _, service := range createServices(sess) {
		fmt.Println("Processing:", service.name)
		if err := callService(service.value); err != nil {
			panic(fmt.Sprintf("Service, %s failed, %v", service.name, err))
		}
	}

	io.Copy(os.Stdout, w.buf)
}

const (
	allowedRecursion = 2
	sliceSize        = 3
)

func callService(svcV reflect.Value) error {
	params := []reflect.Value{
		reflect.ValueOf(aws.BackgroundContext()),
		reflect.Value{},
		reflect.ValueOf(request.Option(
			func(r *request.Request) {
				origURL, _ := url.Parse(*r.Config.Endpoint)
				r.Handlers.Build.PushBack(func(req *request.Request) {
					// Fix URL mutation caused by machinelearning predictendpoint customization
					newURL := r.HTTPRequest.URL
					r.HTTPRequest.URL = origURL
					r.HTTPRequest.URL.Path = newURL.Path
					if !strings.HasPrefix(r.HTTPRequest.URL.Path, "/") {
						r.HTTPRequest.URL.Path = "/" + r.HTTPRequest.URL.Path
					}
					r.HTTPRequest.URL.RawPath = newURL.RawPath
					r.HTTPRequest.URL.RawQuery = newURL.RawQuery

					// Correct ContentLength fields for S3 operations
					if r.ClientInfo.ServiceName == "s3" {
						switch r.Operation.Name {
						case "PutObject", "UploadPart":
							n, _ := computeBodyLength(r.GetBody())
							r.HTTPRequest.Header.Set("Content-Length", strconv.FormatInt(n, 10))
						}
					}
				})
				r.Handlers.Unmarshal.PushBack(func(req *request.Request) {
					if r.ClientInfo.ServiceName == "sqs" {
						// Ignore validation performed by the SQS handlers
						switch r.Operation.Name {
						case "SendMessage", "SendMessageBatch", "ReceiveMessage":
							aerr, ok := r.Error.(awserr.Error)
							if !ok {
								return
							}
							if aerr.Code() == "InvalidChecksum" {
								r.Error = nil
							}
						}
					}
				})
				r.Handlers.Complete.PushBack(func(req *request.Request) {
					if r.Error != nil {
						fmt.Println(r.Params)
					}
				})
			},
		)),
	}

	svcT := svcV.Type()
	n := svcT.NumMethod()

	for i := 0; i < n; i++ {
		fv := svcV.Method(i)
		fm := svcT.Method(i)
		ft := fm.Type
		fName := fm.Name
		if !strings.HasSuffix(fName, "WithContext") || strings.HasSuffix(fName, "PagesWithContext") || strings.HasPrefix(fName, "WaitUntil") {
			continue
		}
		fmt.Println("-", fm.Name)

		it := ft.In(2)
		iv := valueForType(it, visitType(it))

		params[1] = iv
		ovs := fv.Call(params)
		if v := ovs[1]; !v.IsNil() {
			return v.Interface().(error)
		}
	}

	return nil
}

func asVisited(v map[reflect.Type]int, t reflect.Type) (map[reflect.Type]int, bool) {
	if c, ok := v[t]; ok {
		if c == 0 {
			return v, false
		}
		c--
		v[t] = c
	} else {
		v[t] = allowedRecursion
	}

	return v, true
}

var ioReadSeekerType = reflect.TypeOf((*io.ReadSeeker)(nil)).Elem()
var emptyInterfaceType = reflect.TypeOf((*interface{})(nil)).Elem()

func valueForType(vt reflect.Type, visited *visitedType) reflect.Value {
	var v reflect.Value

	vtt := vt
	if vtt.Kind() == reflect.Ptr {
		vtt = vtt.Elem()
	}

	switch vtt.Kind() {
	case reflect.Map:
		v = reflect.MakeMap(vtt)
		kt := vtt.Key()
		kv := valueForType(kt, visited)
		et := vtt.Elem()
		ev := valueForType(et, visited)
		v.SetMapIndex(kv, ev)

	case reflect.Slice:
		v = reflect.MakeSlice(vtt, sliceSize, sliceSize)
		vet := vtt.Elem()
		for i := 0; i < sliceSize; i++ {
			sv := v.Index(i)
			nsv := valueForType(vet, visited)
			sv.Set(nsv)
		}

	case reflect.Interface:
		switch vtt {
		case ioReadSeekerType:
			v = reflect.New(vtt)
			v.Elem().Set(reflect.ValueOf(bytes.NewReader([]byte("byte value"))))
		case emptyInterfaceType:
			v = reflect.ValueOf("empty interface value")
		default:
			panic("value for interface, unknown type" + vtt.String())
		}

	case reflect.String:
		v = reflect.New(vtt)
		v.Elem().SetString("stringValue")

	case reflect.Bool:
		v = reflect.New(vtt)
		v.Elem().SetBool(true)

	case reflect.Uint8: // byte
		v = reflect.New(vtt)
		v.Elem().Set(reflect.ValueOf(uint8('b')))

	case reflect.Int64:
		v = reflect.New(vtt)
		v.Elem().SetInt(987654321)

	case reflect.Float64:
		v = reflect.New(vtt)
		v.Elem().SetFloat(123456789.321)

	case reflect.Struct:
		v = reflect.New(vtt)
		ve := v.Elem()
		n := ve.NumField()
		for i := 0; i < n; i++ {
			fv := ve.Field(i)
			fs := vtt.Field(i)
			ft := fv.Type()
			if len(fs.PkgPath) != 0 {
				continue
			}
			nested, keep := visited.Visit(ft)
			if !keep {
				continue
			}
			nfv := valueForType(ft, nested)
			fv.Set(nfv)
		}
	default:
		panic("unknown type, " + vtt.String())
	}

	if vt.Kind() != reflect.Ptr && v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v
}

func setupServer(out writer) *httptest.Server {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		out.Writef("Path:")
		out.Indent().Writef(r.URL.Path)
		out.Writef("Query:")

		var keys []string
		query := r.URL.Query()
		for k := range query {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			for _, v := range query[k] {
				out.Indent().Writef("%s: %s", k, v)
			}
		}
		out.Writef("Headers:")
		keys = keys[0:0]
		for k := range r.Header {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			for _, v := range r.Header[k] {
				out.Indent().Writef("%s: %s", k, v)
			}
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		r.Body.Close()
		out.Writef("Body:")
		//out.Indent().Writef(base64.StdEncoding.EncodeToString(body))
		out.Indent().Writef(string(body))
	}))

	return server
}

type writer struct {
	buf       *bytes.Buffer
	indent    int
	indentStr string
}

func (w writer) Writef(format string, args ...interface{}) error {
	indent := strings.Repeat(w.indentStr, w.indent)
	w.buf.WriteString(indent)
	w.buf.WriteString(fmt.Sprintf(format, args...))
	w.buf.WriteRune('\n')
	return nil
}
func (w writer) Indent() writer {
	newW := w
	newW.indent++
	return newW
}

type visitedType struct {
	typ  reflect.Type
	left int
	next *visitedType
}

func visitType(t reflect.Type) *visitedType {
	return &visitedType{
		typ:  t,
		left: allowedRecursion,
	}
}

func (v *visitedType) String() string {
	if v == nil {
		return "END"
	}
	return fmt.Sprintf("Type:%v,Kind:%v,Left:%v->%v", v.typ.Name(), v.typ.Kind(), v.left, v.next.String())
}

func (v *visitedType) copy() *visitedType {
	nv := &visitedType{}

	oldNext := v
	newNext := nv

	for {
		*newNext = *oldNext
		oldNext = oldNext.next

		if oldNext == nil {
			break
		}

		newNext.next = &visitedType{}
		newNext = newNext.next
	}

	return nv
}

func (v *visitedType) Visit(t reflect.Type) (*visitedType, bool) {
	if v == nil {
		return visitType(t), true
	}

	nv := v.copy()

	last := nv
	next := nv
	for next != nil {
		last = next
		if next.typ != t {
			next = next.next
			continue
		}

		next.left--
		return nv, next.left >= 0
	}

	last.next = visitType(t)

	return nv, true
}

type service struct {
	name  string
	value reflect.Value
}

func createServices(sess *session.Session) []service {
	return []service{
		{name: "acm", value: reflect.ValueOf(acm.New(sess))},
		{name: "apigateway", value: reflect.ValueOf(apigateway.New(sess))},
		{name: "applicationautoscaling", value: reflect.ValueOf(applicationautoscaling.New(sess))},
		{name: "applicationdiscoveryservice", value: reflect.ValueOf(applicationdiscoveryservice.New(sess))},
		{name: "appstream", value: reflect.ValueOf(appstream.New(sess))},
		{name: "athena", value: reflect.ValueOf(athena.New(sess))},
		{name: "autoscaling", value: reflect.ValueOf(autoscaling.New(sess))},
		{name: "batch", value: reflect.ValueOf(batch.New(sess))},
		{name: "budgets", value: reflect.ValueOf(budgets.New(sess))},
		{name: "clouddirectory", value: reflect.ValueOf(clouddirectory.New(sess))},
		{name: "cloudformation", value: reflect.ValueOf(cloudformation.New(sess))},
		{name: "cloudfront", value: reflect.ValueOf(cloudfront.New(sess))},
		{name: "cloudhsm", value: reflect.ValueOf(cloudhsm.New(sess))},
		{name: "cloudhsmv2", value: reflect.ValueOf(cloudhsmv2.New(sess))},
		{name: "cloudsearch", value: reflect.ValueOf(cloudsearch.New(sess))},
		{name: "cloudsearchdomain", value: reflect.ValueOf(cloudsearchdomain.New(sess))},
		{name: "cloudtrail", value: reflect.ValueOf(cloudtrail.New(sess))},
		{name: "cloudwatch", value: reflect.ValueOf(cloudwatch.New(sess))},
		{name: "cloudwatchevents", value: reflect.ValueOf(cloudwatchevents.New(sess))},
		{name: "cloudwatchlogs", value: reflect.ValueOf(cloudwatchlogs.New(sess))},
		{name: "codebuild", value: reflect.ValueOf(codebuild.New(sess))},
		{name: "codecommit", value: reflect.ValueOf(codecommit.New(sess))},
		{name: "codedeploy", value: reflect.ValueOf(codedeploy.New(sess))},
		{name: "codepipeline", value: reflect.ValueOf(codepipeline.New(sess))},
		{name: "codestar", value: reflect.ValueOf(codestar.New(sess))},
		{name: "cognitoidentity", value: reflect.ValueOf(cognitoidentity.New(sess))},
		{name: "cognitoidentityprovider", value: reflect.ValueOf(cognitoidentityprovider.New(sess))},
		{name: "cognitosync", value: reflect.ValueOf(cognitosync.New(sess))},
		{name: "configservice", value: reflect.ValueOf(configservice.New(sess))},
		{name: "costandusagereportservice", value: reflect.ValueOf(costandusagereportservice.New(sess))},
		{name: "databasemigrationservice", value: reflect.ValueOf(databasemigrationservice.New(sess))},
		{name: "datapipeline", value: reflect.ValueOf(datapipeline.New(sess))},
		{name: "dax", value: reflect.ValueOf(dax.New(sess))},
		{name: "devicefarm", value: reflect.ValueOf(devicefarm.New(sess))},
		{name: "directconnect", value: reflect.ValueOf(directconnect.New(sess))},
		{name: "directoryservice", value: reflect.ValueOf(directoryservice.New(sess))},
		{name: "dynamodb", value: reflect.ValueOf(dynamodb.New(sess))},
		{name: "dynamodbstreams", value: reflect.ValueOf(dynamodbstreams.New(sess))},
		{name: "ec2", value: reflect.ValueOf(ec2.New(sess))},
		{name: "ecr", value: reflect.ValueOf(ecr.New(sess))},
		{name: "ecs", value: reflect.ValueOf(ecs.New(sess))},
		{name: "efs", value: reflect.ValueOf(efs.New(sess))},
		{name: "elasticache", value: reflect.ValueOf(elasticache.New(sess))},
		{name: "elasticbeanstalk", value: reflect.ValueOf(elasticbeanstalk.New(sess))},
		{name: "elasticsearchservice", value: reflect.ValueOf(elasticsearchservice.New(sess))},
		{name: "elastictranscoder", value: reflect.ValueOf(elastictranscoder.New(sess))},
		{name: "elb", value: reflect.ValueOf(elb.New(sess))},
		{name: "elbv2", value: reflect.ValueOf(elbv2.New(sess))},
		{name: "emr", value: reflect.ValueOf(emr.New(sess))},
		{name: "firehose", value: reflect.ValueOf(firehose.New(sess))},
		{name: "gamelift", value: reflect.ValueOf(gamelift.New(sess))},
		{name: "glacier", value: reflect.ValueOf(glacier.New(sess))},
		{name: "glue", value: reflect.ValueOf(glue.New(sess))},
		{name: "greengrass", value: reflect.ValueOf(greengrass.New(sess))},
		{name: "health", value: reflect.ValueOf(health.New(sess))},
		{name: "iam", value: reflect.ValueOf(iam.New(sess))},
		{name: "inspector", value: reflect.ValueOf(inspector.New(sess))},
		{name: "iot", value: reflect.ValueOf(iot.New(sess))},
		{name: "iotdataplane", value: reflect.ValueOf(iotdataplane.New(sess))},
		{name: "kinesis", value: reflect.ValueOf(kinesis.New(sess))},
		{name: "kinesisanalytics", value: reflect.ValueOf(kinesisanalytics.New(sess))},
		{name: "kms", value: reflect.ValueOf(kms.New(sess))},
		{name: "lambda", value: reflect.ValueOf(lambda.New(sess))},
		{name: "lexmodelbuildingservice", value: reflect.ValueOf(lexmodelbuildingservice.New(sess))},
		{name: "lexruntimeservice", value: reflect.ValueOf(lexruntimeservice.New(sess))},
		{name: "lightsail", value: reflect.ValueOf(lightsail.New(sess))},
		{name: "machinelearning", value: reflect.ValueOf(machinelearning.New(sess))},
		{name: "marketplacecommerceanalytics", value: reflect.ValueOf(marketplacecommerceanalytics.New(sess))},
		{name: "marketplaceentitlementservice", value: reflect.ValueOf(marketplaceentitlementservice.New(sess))},
		{name: "marketplacemetering", value: reflect.ValueOf(marketplacemetering.New(sess))},
		{name: "migrationhub", value: reflect.ValueOf(migrationhub.New(sess))},
		{name: "mobile", value: reflect.ValueOf(mobile.New(sess))},
		{name: "mobileanalytics", value: reflect.ValueOf(mobileanalytics.New(sess))},
		{name: "mturk", value: reflect.ValueOf(mturk.New(sess))},
		{name: "opsworks", value: reflect.ValueOf(opsworks.New(sess))},
		{name: "opsworkscm", value: reflect.ValueOf(opsworkscm.New(sess))},
		{name: "organizations", value: reflect.ValueOf(organizations.New(sess))},
		{name: "pinpoint", value: reflect.ValueOf(pinpoint.New(sess))},
		{name: "polly", value: reflect.ValueOf(polly.New(sess))},
		{name: "rds", value: reflect.ValueOf(rds.New(sess))},
		{name: "redshift", value: reflect.ValueOf(redshift.New(sess))},
		{name: "rekognition", value: reflect.ValueOf(rekognition.New(sess))},
		{name: "resourcegroupstaggingapi", value: reflect.ValueOf(resourcegroupstaggingapi.New(sess))},
		{name: "route53", value: reflect.ValueOf(route53.New(sess))},
		{name: "route53domains", value: reflect.ValueOf(route53domains.New(sess))},
		{name: "s3", value: reflect.ValueOf(s3.New(sess))},
		{name: "servicecatalog", value: reflect.ValueOf(servicecatalog.New(sess))},
		{name: "ses", value: reflect.ValueOf(ses.New(sess))},
		{name: "sfn", value: reflect.ValueOf(sfn.New(sess))},
		{name: "shield", value: reflect.ValueOf(shield.New(sess))},
		{name: "simpledb", value: reflect.ValueOf(simpledb.New(sess))},
		{name: "sms", value: reflect.ValueOf(sms.New(sess))},
		{name: "snowball", value: reflect.ValueOf(snowball.New(sess))},
		{name: "sns", value: reflect.ValueOf(sns.New(sess))},
		{name: "sqs", value: reflect.ValueOf(sqs.New(sess))},
		{name: "ssm", value: reflect.ValueOf(ssm.New(sess))},
		{name: "storagegateway", value: reflect.ValueOf(storagegateway.New(sess))},
		{name: "sts", value: reflect.ValueOf(sts.New(sess))},
		{name: "support", value: reflect.ValueOf(support.New(sess))},
		{name: "swf", value: reflect.ValueOf(swf.New(sess))},
		{name: "waf", value: reflect.ValueOf(waf.New(sess))},
		{name: "wafregional", value: reflect.ValueOf(wafregional.New(sess))},
		{name: "workdocs", value: reflect.ValueOf(workdocs.New(sess))},
		{name: "workspaces", value: reflect.ValueOf(workspaces.New(sess))},
		{name: "xray", value: reflect.ValueOf(xray.New(sess))},
	}
}

func computeBodyLength(r io.ReadSeeker) (int64, error) {
	seekable := true
	// Determine if the seeker is actually seekable. ReaderSeekerCloser
	// hides the fact that a io.Readers might not actually be seekable.
	switch v := r.(type) {
	case aws.ReaderSeekerCloser:
		seekable = v.IsSeeker()
	case *aws.ReaderSeekerCloser:
		seekable = v.IsSeeker()
	}
	if !seekable {
		return -1, nil
	}

	curOffset, err := r.Seek(0, 1)
	if err != nil {
		return 0, err
	}

	endOffset, err := r.Seek(0, 2)
	if err != nil {
		return 0, err
	}

	_, err = r.Seek(curOffset, 0)
	if err != nil {
		return 0, err
	}

	return endOffset - curOffset, nil
}
