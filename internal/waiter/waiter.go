package waiter

import (
	"fmt"
	"reflect"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
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
		req := res[0].Interface().(*aws.Request)
		_ = req.Send()
		//out := res[1].Interface()

		for _, a := range w.Acceptors {
			result := false
			switch a.Matcher {
			case "pathAll":
				// TODO implement pathAll match
			case "pathAny":
				// TODO implement pathAny match
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

	return aws.APIError{
		Code:    "ResourceNotReady",
		Message: fmt.Sprintf("exceeded %d wait attempts", w.MaxAttempts),
	}
}
