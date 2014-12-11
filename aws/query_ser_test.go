package aws

import (
	"net/url"
	"reflect"
	"testing"
)

type Filter struct {
	Name   string   `xml:"Name"`
	Values []string `xml:"Value>item"`
}

type DescribeImagesRequest struct {
	DryRun          bool     `xml:"dryRun"`
	ExecutableUsers []string `xml:"ExecutableBy>ExecutableBy"`
	Filters         []Filter `xml:"Filter>Filter"`
	ImageIds        []string `xml:"ImageId>ImageId"`
	Owners          []string `xml:"Owner>Owner"`
}

func TestSerializingEC2QueryRequests(t *testing.T) {
	c := QueryClient{}
	actual := url.Values{}
	req := DescribeImagesRequest{
		Filters: []Filter{
			{Name: "owner-id", Values: []string{"yay"}},
			{Name: "is-public", Values: []string{"true"}},
		},
	}
	if err := c.loadValues(actual, req, ""); err != nil {
		t.Fatal(err)
	}

	expected := url.Values{
		"Filter.1.Name":    []string{"owner-id"},
		"Filter.1.Value.1": []string{"yay"},
		"Filter.2.Name":    []string{"is-public"},
		"Filter.2.Value.1": []string{"true"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Was %#v \n but expected \n %#v", actual, expected)
	}
}

type AutoScalingGroupNamesType struct {
	AutoScalingGroupNames []string `xml:"AutoScalingGroupNames>member"`
	MaxRecords            int      `xml:"MaxRecords"`
	NextToken             string   `xml:"NextToken"`
}

func TestSerializingNonEC2QueryRequests(t *testing.T) {
	c := QueryClient{}
	actual := url.Values{}
	req := AutoScalingGroupNamesType{
		AutoScalingGroupNames: []string{"one", "two"},
	}
	if err := c.loadValues(actual, req, ""); err != nil {
		t.Fatal(err)
	}

	expected := url.Values{
		"AutoScalingGroupNames.member.1": []string{"one"},
		"AutoScalingGroupNames.member.2": []string{"two"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Was %#v \n but expected \n %#v", actual, expected)
	}
}
