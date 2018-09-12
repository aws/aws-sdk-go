package crr

import (
	"reflect"
	"testing"
)

func TestCacheAdd(t *testing.T) {
	cases := []struct {
		limit        int64
		endpoints    []Endpoint
		validKeys    map[string]Endpoint
		expectedSize int
	}{
		{
			limit: 5,
			endpoints: []Endpoint{
				{
					Key: "foo",
					Addresses: []WeightedAddress{
						{
							Address: "0",
						},
					},
				},
				{
					Key: "bar",
					Addresses: []WeightedAddress{
						{
							Address: "1",
						},
					},
				},
				{
					Key: "baz",
					Addresses: []WeightedAddress{
						{
							Address: "2",
						},
					},
				},
				{
					Key: "qux",
					Addresses: []WeightedAddress{
						{
							Address: "3",
						},
					},
				},
				{
					Key: "moo",
					Addresses: []WeightedAddress{
						{
							Address: "4",
						},
					},
				},
			},
			validKeys: map[string]Endpoint{
				"foo": Endpoint{
					Key: "foo",
					Addresses: []WeightedAddress{
						{
							Address: "0",
						},
					},
				},
				"bar": Endpoint{
					Key: "bar",
					Addresses: []WeightedAddress{
						{
							Address: "1",
						},
					},
				},
				"baz": Endpoint{
					Key: "baz",
					Addresses: []WeightedAddress{
						{
							Address: "2",
						},
					},
				},
				"qux": Endpoint{
					Key: "qux",
					Addresses: []WeightedAddress{
						{
							Address: "3",
						},
					},
				},
				"moo": Endpoint{
					Key: "moo",
					Addresses: []WeightedAddress{
						{
							Address: "4",
						},
					},
				},
			},
			expectedSize: 5,
		},
		{
			limit: 2,
			endpoints: []Endpoint{
				{
					Key: "bar",
					Addresses: []WeightedAddress{
						{
							Address: "1",
						},
					},
				},
				{
					Key: "foo",
					Addresses: []WeightedAddress{
						{
							Address: "0",
						},
					},
				},
				{
					Key: "baz",
					Addresses: []WeightedAddress{
						{
							Address: "2",
						},
					},
				},
				{
					Key: "qux",
					Addresses: []WeightedAddress{
						{
							Address: "3",
						},
					},
				},
				{
					Key: "moo",
					Addresses: []WeightedAddress{
						{
							Address: "4",
						},
					},
				},
			},
			validKeys: map[string]Endpoint{
				"foo": Endpoint{
					Key: "foo",
					Addresses: []WeightedAddress{
						{
							Address: "0",
						},
					},
				},
				"bar": Endpoint{
					Key: "bar",
					Addresses: []WeightedAddress{
						{
							Address: "1",
						},
					},
				},
				"baz": Endpoint{
					Key: "baz",
					Addresses: []WeightedAddress{
						{
							Address: "2",
						},
					},
				},
				"qux": Endpoint{
					Key: "qux",
					Addresses: []WeightedAddress{
						{
							Address: "3",
						},
					},
				},
				"moo": Endpoint{
					Key: "moo",
					Addresses: []WeightedAddress{
						{
							Address: "4",
						},
					},
				},
			},
			expectedSize: 2,
		},
	}

	for _, c := range cases {
		cache := NewEndpointCache(c.limit)

		for _, endpoint := range c.endpoints {
			cache.Add(endpoint)
		}

		count := 0
		endpoints := map[string]Endpoint{}
		cache.endpoints.Range(func(key, value interface{}) bool {
			count++

			endpoints[key.(string)] = value.(Endpoint)
			return true
		})

		if e, a := c.expectedSize, cache.size; int64(e) != a {
			t.Errorf("expected %v, but received %v", e, a)
		}

		if e, a := c.expectedSize, count; e != a {
			t.Errorf("expected %v, but received %v", e, a)
		}

		for k, ep := range endpoints {
			endpoint, ok := c.validKeys[k]
			if !ok {
				t.Errorf("unrecognized key %q in cache", k)
			}
			if e, a := endpoint, ep; !reflect.DeepEqual(e, a) {
				t.Errorf("expected %v, but received %v", e, a)
			}
		}
	}
}

func TestCacheGet(t *testing.T) {
	cases := []struct {
		addEndpoints []Endpoint
		validKeys    map[string]Endpoint
		limit        int64
	}{
		{
			limit: 5,
			addEndpoints: []Endpoint{
				{
					Key: "foo",
					Addresses: []WeightedAddress{
						{
							Address: "0",
						},
					},
				},
				{
					Key: "bar",
					Addresses: []WeightedAddress{
						{
							Address: "1",
						},
					},
				},
				{
					Key: "baz",
					Addresses: []WeightedAddress{
						{
							Address: "2",
						},
					},
				},
				{
					Key: "qux",
					Addresses: []WeightedAddress{
						{
							Address: "3",
						},
					},
				},
				{
					Key: "moo",
					Addresses: []WeightedAddress{
						{
							Address: "4",
						},
					},
				},
			},
			validKeys: map[string]Endpoint{
				"foo": Endpoint{
					Key: "foo",
					Addresses: []WeightedAddress{
						{
							Address: "0",
						},
					},
				},
				"bar": Endpoint{
					Key: "bar",
					Addresses: []WeightedAddress{
						{
							Address: "1",
						},
					},
				},
				"baz": Endpoint{
					Key: "baz",
					Addresses: []WeightedAddress{
						{
							Address: "2",
						},
					},
				},
				"qux": Endpoint{
					Key: "qux",
					Addresses: []WeightedAddress{
						{
							Address: "3",
						},
					},
				},
				"moo": Endpoint{
					Key: "moo",
					Addresses: []WeightedAddress{
						{
							Address: "4",
						},
					},
				},
			},
		},
		{
			limit: 2,
			addEndpoints: []Endpoint{
				{
					Key: "bar",
					Addresses: []WeightedAddress{
						{
							Address: "1",
						},
					},
				},
				{
					Key: "foo",
					Addresses: []WeightedAddress{
						{
							Address: "0",
						},
					},
				},
				{
					Key: "baz",
					Addresses: []WeightedAddress{
						{
							Address: "2",
						},
					},
				},
				{
					Key: "qux",
					Addresses: []WeightedAddress{
						{
							Address: "3",
						},
					},
				},
				{
					Key: "moo",
					Addresses: []WeightedAddress{
						{
							Address: "4",
						},
					},
				},
			},
			validKeys: map[string]Endpoint{
				"foo": Endpoint{
					Key: "foo",
					Addresses: []WeightedAddress{
						{
							Address: "0",
						},
					},
				},
				"bar": Endpoint{
					Key: "bar",
					Addresses: []WeightedAddress{
						{
							Address: "1",
						},
					},
				},
				"baz": Endpoint{
					Key: "baz",
					Addresses: []WeightedAddress{
						{
							Address: "2",
						},
					},
				},
				"qux": Endpoint{
					Key: "qux",
					Addresses: []WeightedAddress{
						{
							Address: "3",
						},
					},
				},
				"moo": Endpoint{
					Key: "moo",
					Addresses: []WeightedAddress{
						{
							Address: "4",
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		cache := NewEndpointCache(c.limit)

		for _, endpoint := range c.addEndpoints {
			cache.Add(endpoint)
		}

		keys := []string{}
		cache.endpoints.Range(func(key, value interface{}) bool {
			a := value.(Endpoint)
			e, ok := c.validKeys[key.(string)]
			if !ok {
				t.Errorf("unrecognized key %q in cache", key.(string))
			}

			if !reflect.DeepEqual(e, a) {
				t.Errorf("expected %v, but received %v", e, a)
			}

			keys = append(keys, key.(string))
			return true
		})

		for _, key := range keys {
			a, ok := cache.get(key)
			if !ok {
				t.Errorf("expected key to be present: %q", key)
			}

			e := c.validKeys[key]
			if !reflect.DeepEqual(e, a) {
				t.Errorf("expected %v, but received %v", e, a)
			}
		}
	}
}
