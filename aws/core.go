package aws

import (
	"bytes"
	"container/list"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"regexp"
	"time"
)

const SDKName = "aws-sdk-go"
const SDKVersion = "0.5.0"

func UserAgentHandler(r *Request) {
	r.HTTPRequest.Header.Set("User-Agent", SDKName+"/"+SDKVersion)
}

func SendHandler(r *Request) {
	r.HTTPResponse, r.Error = r.Service.Config.HTTPClient.Do(r.HTTPRequest)
}

type Operation struct {
	*OperationBindings
	Name       string
	HTTPMethod string
	HTTPPath   string
}

type OperationBindings struct {
	InPayload   string
	OutPayload  string
	Required    []string
	QueryParams []string
	URIParams   []string
	InHeaders   []string
	OutHeaders  []string
}

type Service struct {
	Config       *Config
	Handlers     Handlers
	ManualSend   bool
	ServiceName  string
	APIVersion   string
	Endpoint     string
	JSONVersion  string
	TargetPrefix string
}

var schemeRE = regexp.MustCompile("^([^:]+)://")

func (s *Service) Initialize() {
	if s.Config.HTTPClient == nil {
		s.Config.HTTPClient = http.DefaultClient
	}

	s.Handlers.Build.PushBack(UserAgentHandler)
	s.Handlers.Sign.PushBack(BuildContentLength)
	s.Handlers.Send.PushBack(SendHandler)
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

	built bool
}

func NewRequest(service *Service, operation *Operation, params interface{}, data interface{}) *Request {
	method := operation.HTTPMethod
	if method == "" {
		method = "POST"
	}
	httpReq, _ := http.NewRequest(method, "", nil)
	httpReq.URL, _ = url.Parse(service.Endpoint + operation.HTTPPath)

	r := &Request{
		Service:     service,
		Handlers:    service.Handlers.copy(),
		Time:        time.Now(),
		ExpireTime:  300,
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
		fmt.Printf("---[ REQUEST  ]--------------------------------------\n")
		fmt.Printf("%s\n", string(dumpedBody))
		fmt.Printf("-----------------------------------------------------\n\n")
	})
	s.Handlers.Send.PushBack(func(r *Request) {
		defer r.HTTPResponse.Body.Close()
		dumpedBody, _ := httputil.DumpResponse(r.HTTPResponse, true)

		fmt.Printf("---[ RESPONSE ]--------------------------------------\n")
		fmt.Printf("%s\n", string(dumpedBody))
		fmt.Printf("-----------------------------------------------------\n\n")
	})
}

func (r Request) ParamsFilled() bool {
	return reflect.ValueOf(r.Params).Elem().IsValid()
}

func (r Request) DataFilled() bool {
	return reflect.ValueOf(r.Data).Elem().IsValid()
}

func (r *Request) SetBufferBody(buf []byte) {
	r.SetReaderBody(bytes.NewReader(buf))
}

func (r *Request) SetReaderBody(reader io.ReadSeeker) {
	r.HTTPRequest.Body = ioutil.NopCloser(reader)
	r.Body = reader
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

	r.Handlers.Send.Run(r)
	if r.Error != nil {
		return r.Error
	}

	r.Handlers.ValidateResponse.Run(r)
	if r.Error != nil {
		return r.Error
	}

	r.Handlers.Unmarshal.Run(r)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

type Handlers struct {
	Build            HandlerList
	Sign             HandlerList
	Send             HandlerList
	ValidateResponse HandlerList
	Unmarshal        HandlerList
}

func (h *Handlers) copy() Handlers {
	return Handlers{
		Build:            h.Build.copy(),
		Sign:             h.Sign.copy(),
		Send:             h.Send.copy(),
		ValidateResponse: h.ValidateResponse.copy(),
		Unmarshal:        h.Unmarshal.copy(),
	}
}

// Clear removes callback functions for all handlers
func (h *Handlers) Clear() {
	h.Build.Init()
	h.Send.Init()
	h.Sign.Init()
	h.Unmarshal.Init()
	h.ValidateResponse.Init()
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
