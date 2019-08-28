package awstesting_test

import (

	"github.com/aws/aws-sdk-go/awstesting"
	"testing"
)
type sampleStruct struct{
	A int
	B string
	C float64
	D map[string]string
	E map[string]float64
	F map[string]interface{}
}

func TestStringEqual(t *testing.T) {
	cases := map[string]struct {
		expectString string
		actualString string
	}{
		"Test1": {
			expectString: "hello worlds",
			actualString: "hello world",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			mockT := &testing.T{}
			if !awstesting.StringEqual(mockT, c.expectString, c.actualString){
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}

func TestAssertResponseDataEquals(t *testing.T) {
	cases := map[string]struct {
		expectStruct sampleStruct
		actualStruct sampleStruct
	}{
		"Test1": {
			expectStruct: sampleStruct{
				A: 1,
				B: "hey",
				C: 3,
				D: map[string]string{"hello": "world"},
			},
			actualStruct: sampleStruct{
				A: 1,
				B: "hey",
				C: 3.000,
				D: map[string]string{"hello": "world"},
			},

		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			mockT := &testing.T{}
			if !awstesting.AssertResponseDataEquals(mockT, c.expectStruct, c.actualStruct){
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}