package aws

import (
	"bytes"
	"container/list"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/stripe/aws-go/gen/endpoints"
)

func SendHandler(req *Request) {
	req.HTTPResponse, req.Error = req.Service.HTTPClient.Do(req.HTTPRequest)
}

func V4Signer(req *Request) {
	req.Error = req.Service.Context.sign(req.HTTPRequest)
}

type Operation struct {
	Name        string
	HTTPMethod  string
	HTTPPath    string
	InPayload   string
	OutPayload  string
	Required    []string
	QueryParams []string
	UriParams   []string
	InHeaders   []string
	OutHeaders  []string
}

type Service struct {
	Context
	Handlers     Handlers
	HTTPClient   *http.Client
	ManualSend   bool
	Debug        uint
	Endpoint     string
	JSONVersion  string
	TargetPrefix string
}

func (s *Service) Initialize() {
	if s.HTTPClient == nil {
		s.HTTPClient = http.DefaultClient
	}

	s.Handlers.Send.PushBack(SendHandler)

	endpoint, _, _ := endpoints.Lookup(s.Context.Service, s.Context.Region)
	s.Endpoint = endpoint
}

type Request struct {
	*Service
	Time         time.Time
	Operation    *Operation
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Body         io.ReadSeeker
	Params       interface{}
	Error        error
	Data         interface{}
	RequestID    string
	Debug        uint
}

func NewRequest(service *Service, operation *Operation, params interface{}, data interface{}) *Request {
	method := operation.HTTPMethod
	if method == "" {
		method = "POST"
	}
	path := service.Endpoint + operation.HTTPPath

	httpReq, _ := http.NewRequest(method, path, nil)
	if httpReq.URL.Path == "" {
		httpReq.URL.Path = "/"
	}

	r := &Request{
		Service:     service,
		Time:        time.Now(),
		Operation:   operation,
		HTTPRequest: httpReq,
		Body:        nil,
		Params:      params,
		Error:       nil,
		Data:        data,
		Debug:       service.Debug,
	}

	r.AddDebugHandlers()

	return r
}

func (req *Request) AddDebugHandlers() {
	if req.Debug == 0 {
		return
	}

	req.Handlers.Sign.PushBack(func(r *Request) {
		fmt.Printf("=> [%s] %s.%s(%+v)\n", r.Time,
			r.Service.Context.Service, r.Operation.Name, r.Params)
		println("---[ REQUEST  ]--------------------------------------")
		dumpedBody, _ := httputil.DumpRequest(r.HTTPRequest, true)
		println(string(dumpedBody))
		println("-----------------------------------------------------\n")
	})
	req.Handlers.Send.PushBack(func(r *Request) {
		println("---[ RESPONSE ]--------------------------------------")
		defer r.HTTPResponse.Body.Close()
		dumpedBody, _ := httputil.DumpResponse(r.HTTPResponse, true)
		println(string(dumpedBody) + "\n")
		println("-----------------------------------------------------\n")
	})
}

func (r *Request) SetBufferBody(buf []byte) {
	r.SetReaderBody(bytes.NewReader(buf))
}

func (r *Request) SetReaderBody(reader io.ReadSeeker) {
	r.HTTPRequest.Body = ioutil.NopCloser(reader)
	r.Body = reader
}

func (r *Request) Send() error {
	r.Error = nil

	r.Handlers.Build.Run(r)
	if r.Error != nil {
		return r.Error
	}

	r.Handlers.Sign.Run(r)
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

type HandlerList struct {
	list.List
}

func (l *HandlerList) Run(r *Request) {
	for e := l.Front(); e != nil; e = e.Next() {
		h := e.Value.(func(*Request))
		h(r)
	}
}
