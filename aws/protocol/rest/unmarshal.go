package rest

import (
	"reflect"
	"strconv"

	"github.com/awslabs/aws-sdk-go/aws"
)

func Unmarshal(r *aws.Request) {
	if r.DataFilled() {
		unmarshalBody(r)
		unmarshalHeaders(r)
	}
}

func unmarshalBody(r *aws.Request) {
	v := reflect.ValueOf(r.Data).Elem()
	payload := v.FieldByName(r.Operation.OutPayload)
	if payload.IsValid() {
		payload.Set(reflect.ValueOf(r.HTTPResponse.Body))
	}
}

func unmarshalHeaders(r *aws.Request) {
	v := reflect.ValueOf(r.Data).Elem()
	for _, header := range r.Operation.OutHeaders {
		headerType, ok := v.Type().FieldByName(header)
		if ok {
			name := headerType.Tag.Get("name")
			if name == "" {
				name = header
			}
			value := r.HTTPResponse.Header.Get(name)
			if value != "" {
				headerValue := v.FieldByName(header)
				if headerType.Type.Kind() == reflect.Ptr {
					switch headerType.Type.Elem().Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						intvalue, _ := strconv.ParseInt(value, 10, 64)
						headerValue.Set(reflect.ValueOf(&intvalue))
					case reflect.String:
						headerValue.Set(reflect.ValueOf(&value))
					}
				}
			}
		}
	}
}
