package jsonutil

import (
	"reflect"
	"sync"
)

var payloadCache struct {
	sync.RWMutex
	m map[reflect.Type]string
}

func cachedPayloadTag(t reflect.Type) string {
	payloadCache.RLock()
	payload, found := payloadCache.m[t]
	payloadCache.RUnlock()

	if found {
		return payload
	}

	if field, ok := t.FieldByName("SDKShapeTraits"); ok {
		payloadCache.Lock()
		if payloadCache.m == nil {
			payloadCache.m = map[reflect.Type]string{}
		}
		payload = field.Tag.Get("payload")
		payloadCache.m[t] = payload
		payloadCache.Unlock()
	}

	return payload
}
