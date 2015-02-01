package rest

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/awslabs/aws-sdk-go/aws"
)

func Build(r *aws.Request) {
	if r.ParamsFilled() {
		v := reflect.ValueOf(r.Params).Elem()
		buildURI(r, v)
		buildHeaders(r, v)
		buildBody(r, v)
	}
}

func buildHeaders(r *aws.Request, v reflect.Value) {
	for _, header := range r.Operation.InHeaders {
		headerType, _ := v.Type().FieldByName(header)
		headerValue := v.FieldByName(header)
		if headerValue.Kind() == reflect.Ptr {
			if headerValue.IsValid() && headerValue.Elem().IsValid() {
				value := fmt.Sprintf("%v", headerValue.Elem().Interface())
				name := headerType.Tag.Get("name")
				if name == "" {
					name = header
				}
				r.HTTPRequest.Header.Add(name, value)
			}
		}
	}
}

func buildBody(r *aws.Request, v reflect.Value) {
	payload := v.FieldByName(r.Operation.InPayload)
	if payload.IsValid() && payload.Type().Kind() == reflect.Interface {
		reader := payload.Interface().(io.ReadSeeker)
		r.SetReaderBody(reader)
	}
}

func buildURI(r *aws.Request, v reflect.Value) {
	uri := r.HTTPRequest.URL.Path

	// build URI part
	for _, uriParam := range r.Operation.URIParams {
		v := reflect.Indirect(v.FieldByName(uriParam))
		uriParamReplace := ""
		if v.IsValid() {
			uriParamReplace = fmt.Sprintf("%v", v.Interface())
		}
		uri = strings.Replace(uri, "{"+uriParam+"}", aws.EscapePath(uriParamReplace), -1)
		uri = strings.Replace(uri, "{"+uriParam+"+}", uriParamReplace, -1)
	}
	r.HTTPRequest.URL.Path = uri

	// build query string
	for _, qsParam := range r.Operation.QueryParams {
		f, ok := v.Type().FieldByName(qsParam)
		value := reflect.Indirect(v.FieldByName(qsParam))
		if ok && value.IsValid() {
			param := fmt.Sprintf("%v", value.Interface())
			if param != "" {
				r.HTTPRequest.URL.Query().Set(f.Tag.Get("name"), param)
			}
		}
	}
}
