package rest

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

func Unmarshal(r *aws.Request) {
	if r.DataFilled() {
		v := reflect.Indirect(reflect.ValueOf(r.Data))
		unmarshalBody(r, v)
		unmarshalLocationElements(r, v)
	}
}

func unmarshalBody(r *aws.Request, v reflect.Value) {
	if field, ok := v.Type().FieldByName("SDKShapeTraits"); ok {
		if payloadName := field.Tag.Get("payload"); payloadName != "" {
			pfield, _ := v.Type().FieldByName(payloadName)
			if ptag := pfield.Tag.Get("type"); ptag != "" && ptag != "structure" {
				payload := reflect.Indirect(v.FieldByName(payloadName))
				if payload.IsValid() {
					switch payload.Interface().(type) {
					case io.ReadCloser:
						payload.Set(reflect.ValueOf(r.HTTPResponse.Body))
					case []byte:
						b, err := ioutil.ReadAll(r.HTTPResponse.Body)
						if err != nil {
							r.Error = err
						} else {
							payload.Set(reflect.ValueOf(b))
						}
					case string:
						b, err := ioutil.ReadAll(r.HTTPResponse.Body)
						if err != nil {
							r.Error = err
						} else {
							payload.Set(reflect.ValueOf(string(b)))
						}
					default:
						r.Error = fmt.Errorf("unknown payload type %s", payload.Type())
					}
				}
			}
		}
	}
}

func unmarshalLocationElements(r *aws.Request, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		m, field := v.Field(i), v.Type().Field(i)
		if n := field.Name; n[0:1] == strings.ToLower(n[0:1]) {
			continue
		}

		if m.IsValid() {
			name := field.Tag.Get("locationName")
			if name == "" {
				name = field.Name
			}

			switch field.Tag.Get("location") {
			case "header": // we should only ever be pulling out headers
				unmarshalHeader(m, r.HTTPResponse.Header.Get(name))
			}
		}
		if r.Error != nil {
			return
		}
	}
}

func unmarshalHeader(v reflect.Value, header string) error {
	if !v.IsValid() {
		return nil
	}

	switch v.Interface().(type) {
	case *string:
		v.Set(reflect.ValueOf(&header))
	case []byte:
		b, err := base64.StdEncoding.DecodeString(header)
		if err != nil {
			return err
		} else {
			v.Set(reflect.ValueOf(&b))
		}
	case *bool:
		b, err := strconv.ParseBool(header)
		if err != nil {
			return err
		} else {
			v.Set(reflect.ValueOf(&b))
		}
	case *int64:
		i, err := strconv.ParseInt(header, 10, 64)
		if err != nil {
			return err
		} else {
			v.Set(reflect.ValueOf(&i))
		}
	case *int:
		i, err := strconv.ParseInt(header, 10, 32)
		if err != nil {
			return err
		} else {
			v.Set(reflect.ValueOf(&i))
		}
	case *float64:
		f, err := strconv.ParseFloat(header, 64)
		if err != nil {
			return err
		} else {
			v.Set(reflect.ValueOf(&f))
		}
	case *float32:
		f, err := strconv.ParseFloat(header, 32)
		if err != nil {
			return err
		} else {
			v.Set(reflect.ValueOf(&f))
		}
	case *time.Time:
		t, err := time.Parse(time.RFC822Z, header)
		if err != nil {
			return err
		} else {
			v.Set(reflect.ValueOf(&t))
		}
	default:
		err := fmt.Errorf("Unsupported value for param %v (%s)", v.Interface(), v.Type())
		return err
	}
	return nil
}
