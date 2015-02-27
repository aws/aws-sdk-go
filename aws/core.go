package aws

import (
	"bytes"
	"container/list"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"regexp"
	"time"
)

const SDKName = "aws-sdk-go"
const SDKVersion = "0.5.0"

var currentTime = time.Now

func UserAgentHandler(r *Request) {
	r.HTTPRequest.Header.Set("User-Agent", SDKName+"/"+SDKVersion)
}

func SendHandler(r *Request) {
	r.HTTPResponse, r.Error = r.Service.Config.HTTPClient.Do(r.HTTPRequest)
}

func ValidateResponseHandler(r *Request) {
	if r.HTTPResponse.StatusCode == 0 || r.HTTPResponse.StatusCode >= 400 {
		r.Error = APIError{
			StatusCode: r.HTTPResponse.StatusCode,
			Retryable:  r.Service.ShouldRetry(r),
			RetryDelay: r.Service.RetryRules(r),
			RetryCount: r.RetryCount,
		}
	}
}

func AfterRetryHandler(r *Request) {
	delay := 0 * time.Second
	willRetry := false

	if err := Error(r.Error); err != nil {
		delay = err.RetryDelay
		if err.Retryable && r.RetryCount < r.Service.MaxRetries() {
			r.RetryCount++
			willRetry = true
		}
	}

	if willRetry {
		r.Error = nil
		time.Sleep(delay)
	}
}

type Operation struct {
	Name       string
	HTTPMethod string
	HTTPPath   string
}

type Service struct {
	Config            *Config
	Handlers          Handlers
	ManualSend        bool
	ServiceName       string
	APIVersion        string
	Endpoint          string
	JSONVersion       string
	TargetPrefix      string
	RetryRules        func(*Request) time.Duration
	ShouldRetry       func(*Request) bool
	DefaultMaxRetries uint
}

var schemeRE = regexp.MustCompile("^([^:]+)://")

func retryRules(r *Request) time.Duration {
	delay := time.Duration(math.Pow(2, float64(r.RetryCount))) * 30
	return delay * time.Millisecond
}

func shouldRetry(r *Request) bool {
	if err := Error(r.Error); err != nil {
		if err.StatusCode >= 500 {
			return true
		}

		switch err.Code {
		case "ExpiredTokenException":
		case "ProvisionedThroughputExceededException", "Throttling":
			return true
		}
	}
	return false
}

func (s *Service) Initialize() {
	if s.Config.HTTPClient == nil {
		s.Config.HTTPClient = http.DefaultClient
	}

	if s.RetryRules == nil {
		s.RetryRules = retryRules
	}

	if s.ShouldRetry == nil {
		s.ShouldRetry = shouldRetry
	}

	s.DefaultMaxRetries = 3
	s.Handlers.Build.PushBack(UserAgentHandler)
	s.Handlers.Sign.PushBack(BuildContentLength)
	s.Handlers.Send.PushBack(SendHandler)
	s.Handlers.AfterRetry.PushBack(AfterRetryHandler)
	s.Handlers.ValidateResponse.PushBack(ValidateResponseHandler)
	s.AddDebugHandlers()
	s.buildEndpoint()
}

type Request struct {
	*Service
	Handlers     Handlers
	Time         time.Time
	ExpireTime   time.Duration
	Operation    *Operation
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Body         io.ReadSeeker
	Params       interface{}
	Error        error
	Data         interface{}
	RequestID    string
	RetryCount   uint

	built bool
}

func NewRequest(service *Service, operation *Operation, params interface{}, data interface{}) *Request {
	method := operation.HTTPMethod
	if method == "" {
		method = "POST"
	}
	p := operation.HTTPPath
	if p == "" {
		p = "/"
	}

	httpReq, _ := http.NewRequest(method, "", nil)
	httpReq.URL, _ = url.Parse(service.Endpoint + p)

	r := &Request{
		Service:     service,
		Handlers:    service.Handlers.copy(),
		Time:        time.Now(),
		ExpireTime:  0,
		Operation:   operation,
		HTTPRequest: httpReq,
		Body:        nil,
		Params:      params,
		Error:       nil,
		Data:        data,
	}

	return r
}

func (s *Service) buildEndpoint() {
	if s.Config.Endpoint != "" {
		s.Endpoint = s.Config.Endpoint
	} else {
		s.Endpoint = s.endpointForRegion()
	}

	if !schemeRE.MatchString(s.Endpoint) {
		scheme := "https"
		if s.Config.DisableSSL {
			scheme = "http"
		}
		s.Endpoint = scheme + "://" + s.Endpoint
	}
}

func (s *Service) AddDebugHandlers() {
	if s.Config.LogLevel == 0 {
		return
	}

	s.Handlers.Sign.PushBack(func(r *Request) {
		dumpedBody, _ := httputil.DumpRequest(r.HTTPRequest, true)

		fmt.Printf("=> [%s] %s.%s(%+v)\n", r.Time,
			r.Service.ServiceName, r.Operation.Name, r.Params)
		fmt.Printf("---[ REQUEST PRE-SIGN ]------------------------------\n")
		fmt.Printf("%s\n", string(dumpedBody))
		fmt.Printf("-----------------------------------------------------\n")
	})
	s.Handlers.Send.PushFront(func(r *Request) {
		dumpedBody, _ := httputil.DumpRequest(r.HTTPRequest, true)

		fmt.Printf("---[ REQUEST POST-SIGN ]-----------------------------\n")
		fmt.Printf("%s\n", string(dumpedBody))
		fmt.Printf("-----------------------------------------------------\n")
	})
	s.Handlers.Send.PushBack(func(r *Request) {
		fmt.Printf("---[ RESPONSE ]--------------------------------------\n")
		if r.HTTPResponse != nil {
			if r.HTTPResponse.Body != nil {
				defer r.HTTPResponse.Body.Close()
			}
			dumpedBody, _ := httputil.DumpResponse(r.HTTPResponse, true)

			fmt.Printf("%s\n", string(dumpedBody))
		} else if r.Error != nil {
			fmt.Printf("%s\n", r.Error)
		}
		fmt.Printf("-----------------------------------------------------\n")
	})
}

func (r *Request) ParamsFilled() bool {
	return r.Params != nil && reflect.ValueOf(r.Params).Elem().IsValid()
}

func (r *Request) DataFilled() bool {
	return r.Data != nil && reflect.ValueOf(r.Data).Elem().IsValid()
}

func (r *Request) SetBufferBody(buf []byte) {
	r.SetReaderBody(bytes.NewReader(buf))
}

func (r *Request) SetReaderBody(reader io.ReadSeeker) {
	r.HTTPRequest.Body = ioutil.NopCloser(reader)
	r.Body = reader
}

func (s *Service) MaxRetries() uint {
	if s.Config.MaxRetries < 0 {
		return s.DefaultMaxRetries
	} else {
		return uint(s.Config.MaxRetries)
	}
}

func (r *Request) Presign(expireTime time.Duration) (string, error) {
	r.ExpireTime = expireTime
	r.Sign()
	if r.Error != nil {
		return "", r.Error
	} else {
		return r.HTTPRequest.URL.String(), nil
	}
}

func (r *Request) Build() error {
	if !r.built {
		r.Error = nil
		r.Handlers.Build.Run(r)
		r.built = true
	}

	return r.Error
}

func (r *Request) Sign() error {
	r.Build()
	if r.Error != nil {
		return r.Error
	}

	r.Handlers.Sign.Run(r)
	return r.Error
}

func (r *Request) Send() error {
	r.Sign()
	if r.Error != nil {
		return r.Error
	}

	for {
		r.Handlers.Send.Run(r)
		if r.Error != nil {
			return r.Error
		}

		r.Handlers.UnmarshalMeta.Run(r)
		r.Handlers.ValidateResponse.Run(r)
		if r.Error != nil {
			r.Handlers.Retry.Run(r)
			r.Handlers.AfterRetry.Run(r)
			if r.Error != nil {
				r.Handlers.UnmarshalError.Run(r)
				return r.Error
			}
			continue
		}

		r.Handlers.Unmarshal.Run(r)
		if r.Error != nil {
			return r.Error
		}

		return nil
	}
}

type Handlers struct {
	Build            HandlerList
	Sign             HandlerList
	Send             HandlerList
	ValidateResponse HandlerList
	Unmarshal        HandlerList
	UnmarshalMeta    HandlerList
	UnmarshalError   HandlerList
	Retry            HandlerList
	AfterRetry       HandlerList
}

func (h *Handlers) copy() Handlers {
	return Handlers{
		Build:            h.Build.copy(),
		Sign:             h.Sign.copy(),
		Send:             h.Send.copy(),
		ValidateResponse: h.ValidateResponse.copy(),
		Unmarshal:        h.Unmarshal.copy(),
		UnmarshalError:   h.UnmarshalError.copy(),
		UnmarshalMeta:    h.UnmarshalMeta.copy(),
		Retry:            h.Retry.copy(),
		AfterRetry:       h.AfterRetry.copy(),
	}
}

// Clear removes callback functions for all handlers
func (h *Handlers) Clear() {
	h.Build.Init()
	h.Send.Init()
	h.Sign.Init()
	h.Unmarshal.Init()
	h.UnmarshalMeta.Init()
	h.UnmarshalError.Init()
	h.ValidateResponse.Init()
	h.Retry.Init()
	h.AfterRetry.Init()
}

type HandlerList struct {
	list.List
}

func (l HandlerList) copy() HandlerList {
	var n HandlerList
	for e := l.Front(); e != nil; e = e.Next() {
		h := e.Value.(func(*Request))
		n.PushBack(h)
	}
	return n
}

func (l *HandlerList) Run(r *Request) {
	for e := l.Front(); e != nil; e = e.Next() {
		h := e.Value.(func(*Request))
		h(r)
	}
}
