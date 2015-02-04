package rest

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"path"
	"reflect"
	"strconv"
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
		uri = strings.Replace(uri, "{"+uriParam+"}", escapePath(uriParamReplace, true), -1)
		uri = strings.Replace(uri, "{"+uriParam+"+}", escapePath(uriParamReplace, false), -1)
	}
	updatePath(r.HTTPRequest.URL, uri)

	// build query string
	query := r.HTTPRequest.URL.Query()
	for _, qsParam := range r.Operation.QueryParams {
		f, ok := v.Type().FieldByName(qsParam)
		value := reflect.Indirect(v.FieldByName(qsParam))
		if ok && value.IsValid() {
			param := fmt.Sprintf("%v", value.Interface())
			if param != "" {
				query.Set(f.Tag.Get("name"), param)
			}
		}
	}
	r.HTTPRequest.URL.RawQuery = query.Encode()
}

func updatePath(url *url.URL, urlPath string) {
	scheme, query := url.Scheme, url.RawQuery

	// clean up path
	urlPath = path.Clean(urlPath)

	// get formatted URL minus scheme so we can build this into Opaque
	url.Scheme, url.RawQuery, url.Path = "", "", ""
	s := url.String()
	url.Scheme, url.RawQuery = scheme, query

	// build opaque URI
	url.Opaque = s + urlPath
}

// Whether the byte value can be sent without escaping in AWS URLs
var noEscape [256]bool
var noEscapeInitialized = false

// initialise noEscape
func initNoEscape() {
	for i := range noEscape {
		// Amazon expects every character except these escaped
		noEscape[i] = (i >= 'A' && i <= 'Z') ||
			(i >= 'a' && i <= 'z') ||
			(i >= '0' && i <= '9') ||
			i == '-' ||
			i == '.' ||
			i == '_' ||
			i == '~'
	}
}

// escapePath escapes part of a URL path in Amazon style
func escapePath(path string, encodeSep bool) string {
	if !noEscapeInitialized {
		initNoEscape()
		noEscapeInitialized = true
	}

	var buf bytes.Buffer
	for i := 0; i < len(path); i++ {
		c := path[i]
		if noEscape[c] || (c == '/' && !encodeSep) {
			buf.WriteByte(c)
		} else {
			buf.WriteByte('%')
			buf.WriteString(strings.ToUpper(strconv.FormatUint(uint64(c), 16)))
		}
	}
	return buf.String()
}
