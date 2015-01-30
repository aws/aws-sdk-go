package rest

import "reflect"

func PayloadMember(i interface{}, field string) interface{} {
	if i == nil || field == "" {
		return nil
	}

	v := reflect.ValueOf(i).Elem()
	payload := v.FieldByName(field)
	if payload.IsValid() && payload.Type().Kind() != reflect.Interface {
		return payload.Interface()
	} else {
		return nil
	}
}
