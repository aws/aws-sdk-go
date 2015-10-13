package waiter

import (
	"fmt"
	"reflect"
	"time"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

type Config struct {
	Name        string
	Delay       int
	MaxAttempts int
	Operation   string
	Acceptors   []WaitAcceptor
}

type WaitAcceptor struct {
	Expected interface{}
	Matcher  string
	State    string
	Argument string
}

type Waiter struct {
	*Config
	Client interface{}
	Input  interface{}
}

func (w *Waiter) Wait() error {
	client := reflect.ValueOf(w.Client)
	in := reflect.ValueOf(w.Input)
	method := client.MethodByName(w.Config.Operation + "Request")

	for i := 0; i < w.MaxAttempts; i++ {
		res := method.Call([]reflect.Value{in})
		req := res[0].Interface().(*request.Request)
		_ = req.Send()

		for _, a := range w.Acceptors {
			result := false
			switch a.Matcher {
			case "pathAll":
				if vals := awsutil.ValuesAtPath(req.Data, a.Argument); req.Error == nil && vals != nil {
					result = true
					fmt.Println(awsutil.StringValue(req.Data))
					fmt.Println(a.Argument, vals)
					for _, val := range vals {
						if !reflect.DeepEqual(val, a.Expected) {
							result = false
							break
						}
					}
				}
			case "pathAny":
				if vals := awsutil.ValuesAtPath(req.Data, a.Argument); req.Error == nil && vals != nil {
					for _, val := range vals {
						if reflect.DeepEqual(val, a.Expected) {
							result = true
							break
						}
					}
				}
			case "status":
				s := a.Expected.(int)
				result = s == req.HTTPResponse.StatusCode
			}

			if result {
				switch a.State {
				case "success":
					return nil // waiter completed
				case "failure":
					return req.Error // waiter failed
				case "retry":
					// do nothing, just retry
				}
				break
			}
		}

		time.Sleep(time.Second * time.Duration(w.Delay))
	}

	return awserr.New("ResourceNotReady",
		fmt.Sprintf("exceeded %d wait attempts", w.MaxAttempts), nil)
}
